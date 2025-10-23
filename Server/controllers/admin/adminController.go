package admin

import (
	"fmt"

	"gok8s/models"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (con AdminController) Add(c *gin.Context) {
	Username, _ := c.Get("Username")
	fmt.Println(Username)
	c.String(200, "admin add {{.Username}}"+Username.(string))

}

func (con AdminController) Account(c *gin.Context) {
	c.String(200, "admin account")
}

func (con AdminController) Home(c *gin.Context) {
	//pvc := models.GetPVC()

	pods := models.GetPods()
	pvcs := models.GetPVC()
	resources := models.GetResources()

	// fmt.Println("Resources: &&&&&&&&&&&&&&", resources)
	//strpvc := fmt.Sprintf("%v", pvcs)
	//var strpvc = "test only"
	c.HTML(200, "/index.html", gin.H{

		"Podlist":       pods,
		"PodName":       pods[0].Name,
		"PodPhase":      pods[0].Status,
		"PodIP":         pods[0].PodIP,
		"NodeName":      pods[0].NodeName,
		"HostIP":        pods[0].HostIP,
		"StartTime":     pods[0].StartTime,
		"Namespace":     pods[0].Namespace,
		"ResourcesList": resources,
		"ResourcesName": resources[1].Name,
		"PvcList":       pvcs, // 传递完整列表
		// "strvpc":    "test only",
	})

}
