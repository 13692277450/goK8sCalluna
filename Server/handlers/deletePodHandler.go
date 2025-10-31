package handlers

import (
	"context"
	"gok8s/kubernetsServ"
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
