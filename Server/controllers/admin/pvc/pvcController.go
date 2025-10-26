package pvc

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type PVCControllers struct {
	//clientSet *kubernetes.Clientset
}

func (con PVCControllers) PVCController(c *gin.Context) {
	namespace := "default" //namespace

	// Get PVC list
	pvcList, err := kubernetsServ.GetPVCList(namespace)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to get PVC list",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, pvcList)
}
