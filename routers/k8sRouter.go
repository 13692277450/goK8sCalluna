package routers

import (
	"gok8s/models"

	"github.com/gin-gonic/gin"
)

type K8sController struct{}

func K8sManageCenterInit(r *gin.Engine) {
	k8sController := K8sController{}
	k8sRouters := r.Group("/")
	{
		k8sRouters.GET("/", k8sController.K8sManageCenter)
	}
}

func (con K8sController) K8sManageCenter(c *gin.Context) {
	pods := models.GetPods()
	//pvcs := models.GetPVC()

	data := gin.H{
		"Podlist": pods,
		"PvcList": "这里是PVC列表信息...",
	}

	// 只有pods不为空时才添加单个pod的详细信息
	if len(pods) > 0 {
		data["PodName"] = pods[0].Name
		data["PodPhase"] = pods[0].Status
		data["PodIP"] = pods[0].PodIP
		data["NodeName"] = pods[0].NodeName
		data["HostIP"] = pods[0].HostIP
		data["StartTime"] = pods[0].StartTime
		data["Namespace"] = pods[0].Namespace
	}

	c.HTML(200, "/index.html", data)
}
