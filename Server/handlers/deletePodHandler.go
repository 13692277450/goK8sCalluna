package handlers

import (
	"context"
	"fmt"

	"gok8s/kubernetsServ"
	"gok8s/models"
	"log"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeletePodHandler 处理 Pod 删除请求
// 参数：gin.Context, podName, namespace
func DeletePodHandler(c *gin.Context, podName string, namespace string) error {

	// 2. 删除 Pod
	err := kubernetsServ.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Failed to delete pod %s in namespace %s: %v", podName, namespace, err)
		c.JSON(500, gin.H{
			"error":   "Failed to delete pod",
			"details": err.Error(),
		})
		return err
	}

	// 3. 成功响应
	log.Printf("Successfully deleted pod %s in namespace %s", podName, namespace)
	c.JSON(200, gin.H{
		"message":   "Pod deleted successfully",
		"podName":   podName,
		"namespace": namespace,
	})

	return nil
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
			fmt.Printf("Pod:===========================%v --- %v deployment was removed.\n", models.Request.PodName, deploymentName)
			return
		case "StatefulSet":
			statefulSetName := owner.Name
			err := kubernetsServ.Clientset.AppsV1().StatefulSets(namespace).Delete(context.TODO(), statefulSetName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting StatefulSet:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v statefulSetName was removed.\n", models.Request.PodName, statefulSetName)

			return
		case "DaemonSet":
			daemonSetName := owner.Name
			err := kubernetsServ.Clientset.AppsV1().DaemonSets(namespace).Delete(context.TODO(), daemonSetName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting DaemonSet:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v daemonSetName was removed.\n", models.Request.PodName, daemonSetName)
			return
		case "Job":
			jobName := owner.Name
			err := kubernetsServ.Clientset.BatchV1().Jobs(namespace).Delete(context.TODO(), jobName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting Job by CronJob:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v jobName was removed.\n", models.Request.PodName, jobName)
			return
		case "CronJob":
			cronJobName := owner.Name
			err := kubernetsServ.Clientset.BatchV1().CronJobs(namespace).Delete(context.TODO(), cronJobName, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting CronJob:", err)
				return
			}
			fmt.Printf("Pod:===========================%v --- %v cronJobName was removed.\n", models.Request.PodName, cronJobName)
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
	fmt.Printf("Pod:================================== %v yaml file was removed.\n", models.Request.PodName)

}
