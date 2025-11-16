package models

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceInfo 存储Service的基本信息
type ServiceInfo struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Type        string            `json:"type"`
	ClusterIP   string            `json:"clusterIP"`
	ExternalIPs []string          `json:"externalIPs"`
	Ports       []ServicePort     `json:"ports"`
	Selector    map[string]string `json:"selector"`
	CreateTime  v1.Time           `json:"createTime"`
}

// ServicePort 存储Service端口信息
type ServicePort struct {
	Name       string `json:"name"`
	Port       int32  `json:"port"`
	TargetPort int32  `json:"targetPort"`
	Protocol   string `json:"protocol"`
	NodePort   int32  `json:"nodePort,omitempty"`
}

// SetServices 设置全局Service信息
var Services []ServiceInfo

func SetServices(services []ServiceInfo) {
	Services = services
}

func GetServices() []ServiceInfo {
	return Services
}
