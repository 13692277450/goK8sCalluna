package pods

import (
	"context"
	"encoding/json"
	"fmt"
	"gok8s/handlers"
	"gok8s/kubernetsServ"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var request struct {
	PodName   string `json:"podname"`
	NameSpace string `json:"namespace"`
}

type PodsControllers struct{}

func (p PodsControllers) ResponsePodsListController(c *gin.Context) {
	c.JSON(200, kubernetsServ.GetK8sPods())
}

func (p PodsControllers) DeletePodController(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read request body in delete pod handler",
			"details": err.Error(),
		})
		return
	}
	if len(bodyBytes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Empty pod content",
			"details": "The pod content cannot be empty",
		})
		return
	}
	err = json.Unmarshal(bodyBytes, &request)
	if request.PodName == "" {
		c.JSON(400, gin.H{"error": "podName is required"})
		return
	}
	if request.NameSpace == "" {
		c.JSON(400, gin.H{"error": "namespace is required"})
		return
	}
	DeletePodYamlHandler(request.NameSpace, request.PodName)

	handlers.DeletePodHandler(c, request.PodName, request.NameSpace)

}

func DeletePodYamlHandler(namespace, podName string) {
	// 获取Pod
	pod, err := kubernetsServ.Clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return
	}

	// 检测Pod的创建方式
	if len(pod.OwnerReferences) > 0 {
		owner := pod.OwnerReferences[0]
		switch owner.Kind {
		case "Deployment":
			deploymentName := owner.Name
			err := kubernetsServ.Clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting Deployment:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v deployment was removed.\n", request.PodName, deploymentName)
			return
		case "StatefulSet":
			statefulSetName := owner.Name
			err := kubernetsServ.Clientset.AppsV1().StatefulSets(namespace).Delete(context.TODO(), statefulSetName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting StatefulSet:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v statefulSetName was removed.\n", request.PodName, statefulSetName)

			return
		case "DaemonSet":
			daemonSetName := owner.Name
			err := kubernetsServ.Clientset.AppsV1().DaemonSets(namespace).Delete(context.TODO(), daemonSetName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting DaemonSet:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v daemonSetName was removed.\n", request.PodName, daemonSetName)
			return
		case "Job":
			jobName := owner.Name
			err := kubernetsServ.Clientset.BatchV1().Jobs(namespace).Delete(context.TODO(), jobName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting Job by CronJob:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v jobName was removed.\n", request.PodName, jobName)
			return
		case "CronJob":
			cronJobName := owner.Name
			err := kubernetsServ.Clientset.BatchV1().CronJobs(namespace).Delete(context.TODO(), cronJobName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting CronJob:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v cronJobName was removed.\n", request.PodName, cronJobName)
			return
		default:

		}
	} else {
		// There is no ownerReferences，then this Pod was created directly by a user.
		err := kubernetsServ.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
		if err != nil {
			fmt.Println("Error deleting Pod:", err)
			return
		}
		return
	}
	fmt.Printf("Pod:================================== %v yaml file was removed.\n", request.PodName)

}
