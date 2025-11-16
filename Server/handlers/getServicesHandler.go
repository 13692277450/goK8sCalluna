package handlers

import (
	"context"
	"fmt"
	"gok8s/kubernetsServ"
	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetServicesHandler() []models.ServiceInfo {
	// 获取所有命名空间中的Service
	services, err := kubernetsServ.Clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("GetServicesHandler err: %v\n", err)
		return nil
	}

	var serviceInfos []models.ServiceInfo
	for _, service := range services.Items {
		// 转换端口信息
		var ports []models.ServicePort
		for _, port := range service.Spec.Ports {
			servicePort := models.ServicePort{
				Name:       port.Name,
				Port:       port.Port,
				TargetPort: port.TargetPort.IntVal,
				Protocol:   string(port.Protocol),
				NodePort:   port.NodePort,
			}
			ports = append(ports, servicePort)
		}

		// 创建Service信息对象
		serviceInfo := models.ServiceInfo{
			Name:        service.Name,
			Namespace:   service.Namespace,
			Type:        string(service.Spec.Type),
			ClusterIP:   service.Spec.ClusterIP,
			ExternalIPs: service.Spec.ExternalIPs,
			Ports:       ports,
			Selector:    service.Spec.Selector,
			CreateTime:  service.CreationTimestamp,
		}
		serviceInfos = append(serviceInfos, serviceInfo)
	}

	// 设置全局Service信息
	models.SetServices(serviceInfos)
	// 返回响应
	// 如果需要返回详细JSON数据，可以使用：
	// json.NewEncoder(w).Encode(serviceInfos)
	return serviceInfos
}
