package serverServices

import (
	"context"
	"encoding/json"
	"fmt"
	"gok8s/kubernetsServ"
	"log"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type NetworkInfo struct {
	Pods     []PodInfo     `json:"pods"`
	Services []ServiceInfo `json:"services"`
	Nodes    []NodeInfo    `json:"nodes"`
	CNIInfo  CNIInfo       `json:"cniInfo"`
}

type PodInfo struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	PodIP             string            `json:"podIP"`
	HostIP            string            `json:"hostIP"`
	Status            string            `json:"status"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Containers        []ContainerInfo   `json:"containers"`
	NodeName          string            `json:"nodeName"`
	CreationTimestamp metav1.Time       `json:"creationTimestamp"`
}

type ContainerInfo struct {
	Name  string                 `json:"name"`
	Image string                 `json:"image"`
	Ports []corev1.ContainerPort `json:"ports"`
	Ready bool                   `json:"ready"`
}

type ServiceInfo struct {
	Name           string             `json:"name"`
	Namespace      string             `json:"namespace"`
	Type           corev1.ServiceType `json:"type"`
	ClusterIP      string             `json:"clusterIP"`
	ExternalIPs    []string           `json:"externalIPs"`
	LoadBalancerIP string             `json:"loadBalancerIP,omitempty"`
	Ports          []ServicePort      `json:"ports"`
	Selector       map[string]string  `json:"selector"`
	Endpoints      []EndpointInfo     `json:"endpoints"`
	Labels         map[string]string  `json:"labels"`
	Annotations    map[string]string  `json:"annotations"`
}

type ServicePort struct {
	Name       string `json:"name"`
	Port       int32  `json:"port"`
	TargetPort string `json:"targetPort"`
	NodePort   int32  `json:"nodePort,omitempty"`
	Protocol   string `json:"protocol"`
}

type EndpointInfo struct {
	PodName string   `json:"podName"`
	PodIP   string   `json:"podIP"`
	Ports   []string `json:"ports"`
	Ready   bool     `json:"ready"`
}

type NodeInfo struct {
	Name             string            `json:"name"`
	InternalIP       string            `json:"internalIP"`
	ExternalIP       string            `json:"externalIP"`
	Hostname         string            `json:"hostname"`
	OS               string            `json:"os"`
	KernelVersion    string            `json:"kernelVersion"`
	ContainerRuntime string            `json:"containerRuntime"`
	KubeletVersion   string            `json:"kubeletVersion"`
	PodCIDR          string            `json:"podCIDR"`
	Conditions       []NodeCondition   `json:"conditions"`
	Labels           map[string]string `json:"labels"`
	Annotations      map[string]string `json:"annotations"`
}

type NodeCondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type CNIInfo struct {
	CNIConfig       string          `json:"cniConfig"`
	NetworkPolicies []NetworkPolicy `json:"networkPolicies,omitempty"`
	NetworkPlugins  []string        `json:"networkPlugins"`
}

type NetworkPolicy struct {
	Name        string   `json:"name"`
	Namespace   string   `json:"namespace"`
	PolicyTypes []string `json:"policyTypes"`
}

type K8sNetworkCollector struct {
	clientset *kubernetes.Clientset
}

var (
	KubeconfigNetworkInfo *string
	ConfigNetworkInfo     *rest.Config // 主配置
	ClientsetN            *kubernetes.Clientset
)

// NewK8sNetworkCollector 创建 Kubernetes 网络收集器
func NewK8sNetworkCollector() (*K8sNetworkCollector, error) {
	// var config *rest.Config
	// var err error

	// // 优先使用 in-cluster 配置，如果失败则使用 kubeconfig
	// config, err = rest.InClusterConfig()
	// if err != nil {
	// 	// 使用 kubeconfig 文件
	// 	kubeconfigPath := os.Getenv("KUBECONFIG")
	// 	if kubeconfigPath == "" {
	// 		if home := homedir.HomeDir(); home != "" {
	// 			kubeconfigPath = filepath.Join(home, ".kube", "config")
	// 		} else {
	// 			return nil, fmt.Errorf("无法获取用户主目录")
	// 		}
	// 	}

	// 	// 检查 kubeconfig 文件是否存在
	// 	if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
	// 		return nil, fmt.Errorf("kubeconfig 文件不存在: %s", kubeconfigPath)
	// 	}

	// 	// 从 kubeconfig 文件加载配置
	// 	config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("无法加载 Kubernetes 配置: %v", err)
	// 	}
	// }

	// 创建 clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	return nil, fmt.Errorf("创建 Kubernetes 客户端失败: %v", err)
	// }
	ClientsetN = kubernetsServ.Clientset

	return &K8sNetworkCollector{clientset: ClientsetN}, nil
}

// CollectAllNetworkInfo 收集所有网络信息
func (k *K8sNetworkCollector) CollectAllNetworkInfo(ctx context.Context) (*NetworkInfo, error) {
	networkInfo := &NetworkInfo{}

	// 收集 Pod 信息
	pods, err := k.collectPods(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failure to gather Pod information: %v", err)
	}
	networkInfo.Pods = pods

	// 收集 Service 信息
	services, err := k.collectServices(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failure to gath er Services information: %v", err)
	}
	networkInfo.Services = services

	// 收集 Node 信息
	nodes, err := k.collectNodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failure to gather Node information: %v", err)
	}
	networkInfo.Nodes = nodes

	// 收集 CNI 信息
	cniInfo, err := k.collectCNIInfo(ctx)
	if err != nil {
		log.Printf("Failure to gather CNI information: %v", err)
		// 不返回错误，因为 CNI 信息可能无法获取
	}
	networkInfo.CNIInfo = cniInfo
	fmt.Println("NetworkInfo..............\n", networkInfo)
	return networkInfo, nil
}

// collectPods 收集所有 Pod 信息
func (k *K8sNetworkCollector) collectPods(ctx context.Context) ([]PodInfo, error) {
	pods, err := k.clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podInfos []PodInfo
	for _, pod := range pods.Items {
		podInfo := PodInfo{
			Name:              pod.Name,
			Namespace:         pod.Namespace,
			PodIP:             pod.Status.PodIP,
			HostIP:            pod.Status.HostIP,
			Status:            string(pod.Status.Phase),
			Labels:            pod.Labels,
			Annotations:       pod.Annotations,
			NodeName:          pod.Spec.NodeName,
			CreationTimestamp: pod.CreationTimestamp,
		}

		// 收集容器信息
		for _, container := range pod.Spec.Containers {
			containerInfo := ContainerInfo{
				Name:  container.Name,
				Image: container.Image,
				Ports: container.Ports,
			}

			// 检查容器就绪状态
			for _, containerStatus := range pod.Status.ContainerStatuses {
				if containerStatus.Name == container.Name {
					containerInfo.Ready = containerStatus.Ready
					break
				}
			}

			podInfo.Containers = append(podInfo.Containers, containerInfo)
		}

		podInfos = append(podInfos, podInfo)
	}

	return podInfos, nil
}

// collectServices 收集所有 Service 信息
func (k *K8sNetworkCollector) collectServices(ctx context.Context) ([]ServiceInfo, error) {
	services, err := k.clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var serviceInfos []ServiceInfo
	for _, service := range services.Items {
		serviceInfo := ServiceInfo{
			Name:        service.Name,
			Namespace:   service.Namespace,
			Type:        service.Spec.Type,
			ClusterIP:   service.Spec.ClusterIP,
			ExternalIPs: service.Spec.ExternalIPs,
			Selector:    service.Spec.Selector,
			Labels:      service.Labels,
			Annotations: service.Annotations,
		}

		if service.Spec.LoadBalancerIP != "" {
			serviceInfo.LoadBalancerIP = service.Spec.LoadBalancerIP
		}

		// 收集 Service 端口信息
		for _, port := range service.Spec.Ports {
			servicePort := ServicePort{
				Name:     port.Name,
				Port:     port.Port,
				NodePort: port.NodePort,
				Protocol: string(port.Protocol),
			}

			if port.TargetPort.StrVal != "" {
				servicePort.TargetPort = port.TargetPort.StrVal
			} else {
				servicePort.TargetPort = fmt.Sprintf("%d", port.TargetPort.IntVal)
			}

			serviceInfo.Ports = append(serviceInfo.Ports, servicePort)
		}

		// 收集 Endpoints 信息
		endpoints, err := k.collectEndpointsForService(ctx, service.Namespace, service.Name)
		if err != nil {
			log.Printf("Cannot get Service %s/%s endpoints: %v", service.Namespace, service.Name, err)
		} else {
			serviceInfo.Endpoints = endpoints
		}

		serviceInfos = append(serviceInfos, serviceInfo)
	}

	return serviceInfos, nil
}

// collectEndpointsForService 收集 Service 的 Endpoints
func (k *K8sNetworkCollector) collectEndpointsForService(ctx context.Context, namespace, serviceName string) ([]EndpointInfo, error) {
	endpoints, err := k.clientset.CoreV1().Endpoints(namespace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var endpointInfos []EndpointInfo
	for _, subset := range endpoints.Subsets {
		// 获取端口信息
		var ports []string
		for _, port := range subset.Ports {
			ports = append(ports, fmt.Sprintf("%s:%d", port.Protocol, port.Port))
		}

		// 获取地址信息
		for _, address := range subset.Addresses {
			endpointInfo := EndpointInfo{
				PodIP: address.IP,
				Ports: ports,
				Ready: true,
			}

			if address.TargetRef != nil && address.TargetRef.Kind == "Pod" {
				endpointInfo.PodName = address.TargetRef.Name
			}

			endpointInfos = append(endpointInfos, endpointInfo)
		}

		// 处理未就绪的 endpoints
		for _, address := range subset.NotReadyAddresses {
			endpointInfo := EndpointInfo{
				PodIP: address.IP,
				Ports: ports,
				Ready: false,
			}

			if address.TargetRef != nil && address.TargetRef.Kind == "Pod" {
				endpointInfo.PodName = address.TargetRef.Name
			}

			endpointInfos = append(endpointInfos, endpointInfo)
		}
	}

	return endpointInfos, nil
}

// collectNodes 收集所有 Node 信息
func (k *K8sNetworkCollector) collectNodes(ctx context.Context) ([]NodeInfo, error) {
	nodes, err := k.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nodeInfos []NodeInfo
	for _, node := range nodes.Items {
		nodeInfo := NodeInfo{
			Name:             node.Name,
			Hostname:         node.Status.NodeInfo.MachineID,
			OS:               fmt.Sprintf("%s %s", node.Status.NodeInfo.OperatingSystem, node.Status.NodeInfo.OSImage),
			KernelVersion:    node.Status.NodeInfo.KernelVersion,
			ContainerRuntime: node.Status.NodeInfo.ContainerRuntimeVersion,
			KubeletVersion:   node.Status.NodeInfo.KubeletVersion,
			PodCIDR:          node.Spec.PodCIDR,
			Labels:           node.Labels,
			Annotations:      node.Annotations,
		}

		// 获取节点 IP 地址
		for _, address := range node.Status.Addresses {
			switch address.Type {
			case corev1.NodeInternalIP:
				nodeInfo.InternalIP = address.Address
			case corev1.NodeExternalIP:
				nodeInfo.ExternalIP = address.Address
			case corev1.NodeHostName:
				nodeInfo.Hostname = address.Address
			}
		}

		// 收集节点状态
		for _, condition := range node.Status.Conditions {
			nodeCondition := NodeCondition{
				Type:    string(condition.Type),
				Status:  string(condition.Status),
				Message: condition.Message,
			}
			nodeInfo.Conditions = append(nodeInfo.Conditions, nodeCondition)
		}

		nodeInfos = append(nodeInfos, nodeInfo)
	}

	return nodeInfos, nil
}

// collectCNIInfo 收集 CNI 相关信息
func (k *K8sNetworkCollector) collectCNIInfo(ctx context.Context) (CNIInfo, error) {
	cniInfo := CNIInfo{
		NetworkPlugins: []string{},
	}

	// 尝试检测 CNI 配置
	daemonsets, err := k.clientset.AppsV1().DaemonSets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		for _, ds := range daemonsets.Items {
			if strings.Contains(ds.Name, "cni") || strings.Contains(ds.Name, "flannel") ||
				strings.Contains(ds.Name, "calico") || strings.Contains(ds.Name, "weave") {
				cniInfo.NetworkPlugins = append(cniInfo.NetworkPlugins, ds.Name)
			}
		}
	}

	// 尝试获取 NetworkPolicies
	networkPolicies, err := k.clientset.NetworkingV1().NetworkPolicies("").List(ctx, metav1.ListOptions{})
	if err == nil {
		for _, np := range networkPolicies.Items {
			policyTypes := make([]string, len(np.Spec.PolicyTypes))
			for i, policyType := range np.Spec.PolicyTypes {
				policyTypes[i] = string(policyType)
			}

			networkPolicy := NetworkPolicy{
				Name:        np.Name,
				Namespace:   np.Namespace,
				PolicyTypes: policyTypes,
			}
			cniInfo.NetworkPolicies = append(cniInfo.NetworkPolicies, networkPolicy)
		}
	}

	// 通过 kube-system 命名空间的 Pod 推断 CNI
	pods, err := k.clientset.CoreV1().Pods("kube-system").List(ctx, metav1.ListOptions{})
	if err == nil {
		for _, pod := range pods.Items {
			if strings.Contains(pod.Name, "cni") || strings.Contains(pod.Name, "flannel") ||
				strings.Contains(pod.Name, "calico") || strings.Contains(pod.Name, "weave") {
				cniInfo.CNIConfig = fmt.Sprintf("Collected CNI Pod: %s", pod.Name)
				break
			}
		}
	}

	return cniInfo, nil
}

// PrintNetworkInfo 打印网络信息
func (k *K8sNetworkCollector) PrintNetworkInfo(info *NetworkInfo) {
	fmt.Println("=== Kubernetes Network Structure ===")

	fmt.Printf("\n=== Node Info (%d nodes) ===\n", len(info.Nodes))
	for _, node := range info.Nodes {
		fmt.Printf("Node: %s\n", node.Name)
		fmt.Printf("  Internal IP: %s\n", node.InternalIP)
		fmt.Printf("  External IP: %s\n", node.ExternalIP)
		fmt.Printf("  Pod CIDR: %s\n", node.PodCIDR)
		fmt.Printf("  OS: %s\n", node.OS)
		fmt.Printf("  Kubelet: %s\n", node.KubeletVersion)
		fmt.Println()
	}

	fmt.Printf("\n=== Pod Info (%d Pods) ===\n", len(info.Pods))
	for _, pod := range info.Pods {
		fmt.Printf("Pod: %s/%s\n", pod.Namespace, pod.Name)
		fmt.Printf("  IP: %s (Host: %s)\n", pod.PodIP, pod.HostIP)
		fmt.Printf("  Status: %s, Node: %s\n", pod.Status, pod.NodeName)
		if len(pod.Containers) > 0 {
			fmt.Printf("  Capacitor:\n")
			for _, container := range pod.Containers {
				fmt.Printf("    - %s (%s) Ready: %t\n", container.Name, container.Image, container.Ready)
				if len(container.Ports) > 0 {
					for _, port := range container.Ports {
						fmt.Printf("      Port: %d/%s\n", port.ContainerPort, port.Protocol)
					}
				}
			}
		}
		fmt.Println()
	}

	fmt.Printf("\n=== Service Information (%d 个 Service) ===\n", len(info.Services))
	for _, service := range info.Services {
		fmt.Printf("Service: %s/%s\n", service.Namespace, service.Name)
		fmt.Printf("  Type: %s, ClusterIP: %s\n", service.Type, service.ClusterIP)

		if len(service.ExternalIPs) > 0 {
			fmt.Printf("  ExternalIPs: %v\n", service.ExternalIPs)
		}
		if service.LoadBalancerIP != "" {
			fmt.Printf("  LoadBalancerIP: %s\n", service.LoadBalancerIP)
		}

		if len(service.Ports) > 0 {
			fmt.Printf("  Port:\n")
			for _, port := range service.Ports {
				fmt.Printf("    - %s: %d -> %s", port.Name, port.Port, port.TargetPort)
				if port.NodePort > 0 {
					fmt.Printf(" (NodePort: %d)", port.NodePort)
				}
				fmt.Printf(" %s\n", port.Protocol)
			}
		}

		if len(service.Endpoints) > 0 {
			fmt.Printf("  Endpoints:\n")
			for _, endpoint := range service.Endpoints {
				status := "Ready"
				if !endpoint.Ready {
					status = "NotReady"
				}
				fmt.Printf("    - %s (%s) %s 端口: %v\n", endpoint.PodName, endpoint.PodIP, status, endpoint.Ports)
			}
		}
		fmt.Println()
	}

	fmt.Printf("\n=== CNI Information ===\n")
	fmt.Printf("CNI config: %s\n", info.CNIInfo.CNIConfig)
	fmt.Printf("Network Plugin: %v\n", info.CNIInfo.NetworkPlugins)
	if len(info.CNIInfo.NetworkPolicies) > 0 {
		fmt.Printf("Network strategy (%d pcs):\n", len(info.CNIInfo.NetworkPolicies))
		for _, policy := range info.CNIInfo.NetworkPolicies {
			fmt.Printf("  - %s/%s Type: %v\n", policy.Namespace, policy.Name, policy.PolicyTypes)
		}
	}
}

// SaveToFile 将网络信息保存到 JSON 文件
func (k *K8sNetworkCollector) SaveToFile(info *NetworkInfo, filename string) error {
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func K8sNetworkInfo() {
	fmt.Println("Starting collect Kubernetes network information...")

	// 创建网络收集器
	collector, err := NewK8sNetworkCollector()
	if err != nil {
		log.Fatalf("Failure to create Kubernetes network information collector: %v", err)
	}

	// 收集网络信息
	ctx := context.Background()
	networkInfo, err := collector.CollectAllNetworkInfo(ctx)
	if err != nil {
		log.Fatalf("Failure to gather network information: %v", err)
	}

	// 打印信息
	collector.PrintNetworkInfo(networkInfo)

	// 保存到文件
	err = collector.SaveToFile(networkInfo, "k8s-network-info.json")
	if err != nil {
		log.Printf("Failure to save file: %v", err)
	} else {
		fmt.Println("Network infromation was saved to k8s-network-info.json")
	}

	fmt.Println("Gathering completed!")
}
