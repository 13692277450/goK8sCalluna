package admin

import (
	"github.com/gin-gonic/gin"
)

type PVCController struct {
	//clientSet *kubernetes.Clientset
}

func (con PVCController) Home(c *gin.Context) {
	//pvc := models.GetPVC()
	// kubernetsServ.PVCList
	c.HTML(200, "/index.html", gin.H{
		"PvcList": []string{"test1", "test2", "test3"}, // 测试数据
		"title":   "PVC",
	})
}
