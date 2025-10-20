package pods

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type PodsController struct{}

// var podsJsonInfo models.PodInfo

func (p PodsController) ResponsePodsListController(c *gin.Context) {
	//podsJsonInfo, err := json.marshal(kubernetsServ.GetK8sPods())
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to marshal pod info"})
	// 	return
	// }
	c.JSON(200, kubernetsServ.GetK8sPods())
	//c.IndentedJSON(200, podsJsonInfo)
}

// func GetPodInfo() []corev1.Pod {
// 	// 使用默认的 kubeconfig 路径
// 	kubeconfig := clientcmd.RecommendedHomeFile
// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return podList.Items
// }
