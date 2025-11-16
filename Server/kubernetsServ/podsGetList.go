package kubernetsServ

import (
	"context"
	"fmt"
	"log"

	"gok8s/config"
	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetK8sPods() []models.PodInfo {

	pods, err := Clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("List Pods error: ", err)
		log.Fatal(err)
	}

	fmt.Println("Pod List: ")
	var PodInfos []models.PodInfo
	for _, pod := range pods.Items {
		podInfo := models.PodInfo{
			Name:     pod.Name,
			Status:   string(pod.Status.Phase),
			NodeName: pod.Spec.NodeName,
			HostIP:   pod.Status.HostIP,
			PodIP:    pod.Status.PodIP,
			StartTime: func() string {
				if pod.Status.StartTime == nil || pod.Status.StartTime.IsZero() {
					return "1970-01-01 08:00:00 +0800 CST"
				}
				return pod.Status.StartTime.String()
			}(),
			Namespace: pod.Namespace,
		}
		PodInfos = append(PodInfos, podInfo)
		// fmt.Printf("名称：%s，状态：%s, NodeName: %s, HostIP: %s, PodIP: %s, StartTime: %s, NameSpace: %s \n",
		// 	pod.Name, pod.Status.Phase, pod.Spec.NodeName, pod.Status.HostIP, pod.Status.PodIP, pod.Status.StartTime, pod.Namespace)
	}
	models.SetPods(PodInfos)
	config.Lg.Infof("Get Pods succeeded")
	return PodInfos

}

// GetPodResources 获取指定命名空间中Pod的资源使用情况
func GetPodResources1(namespace, podName string) {
	// 从 Kubernetes API 检索 Pod 信息
	pod, err := Clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting pod: %v\n", err)
		return
	}

	// scan all containers
	for _, container := range pod.Spec.Containers {
		fmt.Printf("Container: %s\n", container.Name)

		// cpu  limit
		if container.Resources.Limits != nil {
			if cpuLimit, exists := container.Resources.Limits["cpu"]; exists {
				fmt.Printf("  CPU Limit: %s\n", cpuLimit.String())
			}
		}

		// cpu request
		if container.Resources.Requests != nil {
			if cpuRequest, exists := container.Resources.Requests["cpu"]; exists {
				fmt.Printf("  CPU Request: %s\n", cpuRequest.String())
			}
		}

		// memory limit
		if container.Resources.Limits != nil {
			if memLimit, exists := container.Resources.Limits["memory"]; exists {
				fmt.Printf("  Memory Limit: %s\n", memLimit.String())
			}
		}

		// memory request
		if container.Resources.Requests != nil {
			if memRequest, exists := container.Resources.Requests["memory"]; exists {
				fmt.Printf("  Memory Request: %s\n", memRequest.String())
			}
		}
	}
}

//fmt.Println("Pvcs: ", pvcs)

//======================
// deployment := &appsv1.Deployment{
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "example-deployment",
// 	},
// 	Spec: appsv1.DeploymentSpec{
// 		Replicas: int32Ptr(3),
// 		Selector: &metav1.LabelSelector{
// 			MatchLabels: map[string]string{"app": "example-app"},
// 		},
// 		Template: corev1.PodTemplateSpec{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Labels: map[string]string{"app": "example-app"},
// 			},
// 			Spec: corev1.PodSpec{
// 				Containers: []corev1.Container{
// 					{
// 						Name:  "example-container",
// 						Image: "nginx:latest",
// 					},
// 				},
// 			},
// 		},
// 	},
// }
// result, err := Clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
// if err != nil {
// 	fmt.Println("Create deployment error: ", err)
// 	log.Fatal(err)
// }
// fmt.Printf("Deployment %q 已创建\n", result.GetObjectMeta().GetName())
//}

//func int32Ptr(i int32) *int32 { return &i }
