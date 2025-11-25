package kubernetsServ

import (
	"context"
	"fmt"
	"log"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetEtcdInfo 获取etcd相关信息
func GetEtcdInfo() {
	// 情况1：如果etcd作为独立Pod运行（常见于外部etcd集群）
	etcdPods, err := Clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "component=etcd",
	})
	if err != nil {
		log.Printf("Failed to list etcd pods: %v", err)
		return
	}

	if len(etcdPods.Items) > 0 {
		fmt.Println("Found etcd pods:")
		for _, pod := range etcdPods.Items {
			fmt.Printf("- Name: %s\n- Status: %s\n- Node: %s\n- IP: %s\n",
				pod.Name,
				pod.Status.Phase,
				pod.Spec.NodeName,
				pod.Status.PodIP,
			)
			// 检查etcd容器日志（如果需要）
			// GetEtcdLogs(pod.Name)
		}
		return
	}

	// 情况2：如果etcd由Kubernetes API Server管理（常见于内置etcd）
	fmt.Println("No standalone etcd pods found. Trying to get etcd info via API server...")
	
	// 获取节点信息（etcd通常运行在控制平面节点上）
	nodes, err := Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: "node-role.kubernetes.io/control-plane=",
	})
	if err != nil {
		log.Printf("Failed to list control plane nodes: %v", err)
		return
	}

	if len(nodes.Items) > 0 {
		fmt.Println("Control plane nodes (where etcd might be running):")
		for _, node := range nodes.Items {
			fmt.Printf("- Name: %s\n- Status: %s\n- Addresses: %v\n",
				node.Name,
				node.Status.Conditions[len(node.Status.Conditions)-1].Type, // 最后一个状态
				node.Status.Addresses,
			)
		}
	} else {
		fmt.Println("No control plane nodes found. This may indicate a non-standard Kubernetes setup.")
	}
}

// GetEtcdLogs 获取特定etcd容器的日志（可选）
func GetEtcdLogs(podName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	logReader, err := Clientset.CoreV1().Pods("kube-system").GetLogs(podName, &corev1.PodLogOptions{}).Stream(ctx)
	if err != nil {
		log.Printf("Failed to get logs for pod %s: %v", podName, err)
		return
	}
	defer logReader.Close()

	buf := make([]byte, 1024)
	n, _ := logReader.Read(buf)
	fmt.Printf("Logs for %s:\n%s\n", podName, string(buf[:n]))
}