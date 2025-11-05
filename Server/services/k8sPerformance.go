package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gok8s/kubernetsServ"
	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

var (
	maxRetries    = 3
	retryDelay    = 2 * time.Second
	NodesResult   = []models.MetricsNodes{}
	PodsResult    = []models.MetricsPods{}
	ctx           = context.Background()
	metricsClient *versioned.Clientset // 包级别初始化metricsClient
)

// initMetricsClient 初始化metrics客户端
func initMetricsClient() error {
	if metricsClient != nil {
		return nil
	}

	// 检查是否在Kubernetes集群中运行
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" || os.Getenv("KUBERNETES_SERVICE_PORT") == "" {
		log.Println("Warning: Not running inside Kubernetes cluster, metrics-server may not be accessible")
	}

	// 初始化metrics客户端
	client, err := versioned.NewForConfig(kubernetsServ.Config)
	if err != nil {
		log.Printf("Failed to create metrics client: %v", err)
		if kubernetsServ.Config == nil {
			log.Println("Error: Config object is nil")
		} else {
			log.Printf("Config Host: %v", kubernetsServ.Config.Host)
		}
		return err
	}
	metricsClient = client
	log.Println("Metrics client created successfully")
	return nil
}

// K8sNodesPerformance collects and displays Kubernetes node performance metrics
func K8sNodesPerformance() []models.MetricsNodes {
	NodesResult = []models.MetricsNodes{}
	// 初始化metrics客户端
	if err := initMetricsClient(); err != nil {
		return nil
	}

	// 获取所有节点并显示CPU和内存使用情况
	nodes, err := kubernetsServ.Clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Failed to get node list: %v", err)
		return nil
	}
	log.Printf("Got %d nodes", len(nodes.Items))

	for _, node := range nodes.Items {
		nodeName := node.GetName()
		var nodeMetrics *metricsv1beta1.NodeMetrics
		var metricsErr error

		// 带重试的获取节点指标
		for i := 0; i < maxRetries; i++ {
			log.Printf("Attempting to get metrics for node %s (attempt %d/%d)...", nodeName, i+1, maxRetries)

			nodeMetrics, metricsErr = metricsClient.MetricsV1beta1().NodeMetricses().Get(ctx, nodeName, metav1.GetOptions{})
			if metricsErr == nil {
				break // 成功，退出重试循环
			}

			log.Printf("Failed to get metrics for node %s (attempt %d/%d): %v", nodeName, i+1, maxRetries, metricsErr)
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
			}
		}

		if metricsErr != nil {
			// 提供更详细的错误诊断
			if metricsErr.Error() == "the server is currently unable to handle the request" {
				log.Printf("Error: metrics-server is currently unable to handle the request for node %s. Possible reasons:\n"+
					"1. metrics-server is not fully started\n"+
					"2. metrics-server lacks necessary RBAC permissions\n"+
					"3. metrics-server is misconfigured\n"+
					"4. Network connectivity issues", nodeName)
			}
			log.Printf("Unable to get metrics for node %s: %v", nodeName, metricsErr)
			continue
		}

		cpuUsage := nodeMetrics.Usage.Cpu()
		memoryUsage := nodeMetrics.Usage.Memory()
		nodesData := models.MetricsNodes{
			NodeName: nodeName,
			Cpu:      cpuUsage.AsApproximateFloat64(),
			Memory:   memoryUsage.AsApproximateFloat64(),
		}
		NodesResult = append(NodesResult, nodesData)
		models.SetMetricsNods(NodesResult)

		fmt.Printf("Node %s: CPU usage=%s, memory usage=%s\n", nodeName, cpuUsage.String(), memoryUsage.String())
	}
	return NodesResult
}

// K8sPodsPerformance collects and displays Kubernetes pod performance metrics
func K8sPodsPerformance() []models.MetricsPods {
	PodsResult = []models.MetricsPods{}
	// 初始化metrics客户端
	if err := initMetricsClient(); err != nil {
		return nil
	}

	// 获取所有pod并显示CPU和内存使用情况
	pods, err := kubernetsServ.Clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Failed to get pod list: %v", err)
		return nil
	}
	log.Printf("Got %d pods", len(pods.Items))

	for _, pod := range pods.Items {
		podName := pod.GetName()
		namespaceName := pod.GetNamespace()

		// 带重试的获取pod指标
		var podMetrics *metricsv1beta1.PodMetrics
		var podMetricsErr error

		for i := 0; i < maxRetries; i++ {
			log.Printf("Attempting to get metrics for pod %s/%s (attempt %d/%d)...", namespaceName, podName, i+1, maxRetries)
			podMetrics, podMetricsErr = metricsClient.MetricsV1beta1().PodMetricses(namespaceName).Get(ctx, podName, metav1.GetOptions{})
			if podMetricsErr == nil {
				break // 成功，退出重试循环
			}

			log.Printf("Failed to get metrics for pod %s/%s (attempt %d/%d): %v", namespaceName, podName, i+1, maxRetries, podMetricsErr)
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
			}
		}

		// 处理所有重试后仍然失败的情况
		if podMetricsErr != nil {
			// 提供更详细的错误诊断
			if podMetricsErr.Error() == "the server is currently unable to handle the request" {
				log.Printf("Error: metrics-server is currently unable to handle the request for pod %s/%s. Possible reasons:\n"+
					"1. metrics-server is not fully started\n"+
					"2. metrics-server lacks necessary RBAC permissions\n"+
					"3. metrics-server is misconfigured\n"+
					"4. Network connectivity issues", namespaceName, podName)
			}
			log.Printf("Unable to get metrics for pod %s/%s after %d retries: %v", namespaceName, podName, maxRetries, podMetricsErr)
			continue
		}

		// 确保podMetrics不为nil
		if podMetrics == nil {
			log.Printf("Critical error: podMetrics is nil for pod %s/%s after successful API call", namespaceName, podName)
			continue
		}

		// 成功获取指标后的处理
		containers := podMetrics.Containers
		for _, container := range containers {
			cpuUsage := container.Usage.Cpu()
			memoryUsage := container.Usage.Memory()
			podData := models.MetricsPods{
				PodName:   pod.Name,
				Namespace: pod.Namespace,
				Container: container.Name,
				Cpu:       cpuUsage.AsApproximateFloat64(),
				Memory:    memoryUsage.AsApproximateFloat64(),
				Size:      strconv.Itoa(pod.Size()),
			}
			PodsResult = append(PodsResult, podData)
			models.SetMetricsPods(PodsResult)
			fmt.Printf("Pod %s in Namespace %s: Container %s - CPU usage=%s, memory usage=%s\n",
				podName, namespaceName, container.Name, cpuUsage.String(), memoryUsage.String())
		}
	}
	return PodsResult
}
