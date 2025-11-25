package serverServices

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type NetworkWatcher struct {
	clientset *kubernetes.Clientset
}

func NewNetworkWatcher() (*NetworkWatcher, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &NetworkWatcher{clientset: clientset}, nil
}

// WatchNetworkChanges 监视网络变化
func (w *NetworkWatcher) WatchNetworkChanges(ctx context.Context) {
	fmt.Println("开始监视 Kubernetes 网络变化...")

	// 监视 Pod 变化
	go w.watchPods(ctx)

	// 监视 Service 变化
	go w.watchServices(ctx)

	// 监视 Endpoints 变化
	go w.watchEndpoints(ctx)
}

func (w *NetworkWatcher) watchPods(ctx context.Context) {
	watcher, err := w.clientset.CoreV1().Pods("").Watch(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("监视 Pod 失败: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-watcher.ResultChan():
			if !ok {
				log.Println("Pod 监视通道关闭")
				return
			}

			pod, ok := event.Object.(*corev1.Pod)
			if !ok {
				continue
			}

			switch event.Type {
			case watch.Added:
				fmt.Printf("[POD CREATED] %s/%s IP: %s Node: %s\n",
					pod.Namespace, pod.Name, pod.Status.PodIP, pod.Spec.NodeName)
			case watch.Modified:
				if pod.Status.PodIP != "" {
					fmt.Printf("[POD UPDATED] %s/%s IP: %s Status: %s\n",
						pod.Namespace, pod.Name, pod.Status.PodIP, pod.Status.Phase)
				}
			case watch.Deleted:
				fmt.Printf("[POD DELETED] %s/%s\n", pod.Namespace, pod.Name)
			}
		}
	}
}

func (w *NetworkWatcher) watchServices(ctx context.Context) {
	watcher, err := w.clientset.CoreV1().Services("").Watch(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("监视 Service 失败: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-watcher.ResultChan():
			if !ok {
				log.Println("Service 监视通道关闭")
				return
			}

			service, ok := event.Object.(*corev1.Service)
			if !ok {
				continue
			}

			switch event.Type {
			case watch.Added:
				fmt.Printf("[SERVICE CREATED] %s/%s Type: %s ClusterIP: %s\n",
					service.Namespace, service.Name, service.Spec.Type, service.Spec.ClusterIP)
			case watch.Modified:
				fmt.Printf("[SERVICE UPDATED] %s/%s ClusterIP: %s\n",
					service.Namespace, service.Name, service.Spec.ClusterIP)
			case watch.Deleted:
				fmt.Printf("[SERVICE DELETED] %s/%s\n", service.Namespace, service.Name)
			}
		}
	}
}

func (w *NetworkWatcher) watchEndpoints(ctx context.Context) {
	watcher, err := w.clientset.CoreV1().Endpoints("").Watch(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("监视 Endpoints 失败: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-watcher.ResultChan():
			if !ok {
				log.Println("Endpoints 监视通道关闭")
				return
			}

			endpoints, ok := event.Object.(*corev1.Endpoints)
			if !ok {
				continue
			}

			if event.Type == watch.Modified {
				fmt.Printf("[ENDPOINTS UPDATED] %s/%s\n", endpoints.Namespace, endpoints.Name)
				for _, subset := range endpoints.Subsets {
					fmt.Printf("  Addresses: %d, NotReady: %d\n",
						len(subset.Addresses), len(subset.NotReadyAddresses))
				}
			}
		}
	}
}

func main() {
	watcher, err := NewNetworkWatcher()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	watcher.WatchNetworkChanges(ctx)

	// 保持程序运行
	select {}
}
