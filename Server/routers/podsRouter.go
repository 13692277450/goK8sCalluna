package routers

import (
	"gok8s/controllers/pods"

	"github.com/gin-gonic/gin"
)

func K8sManageCenterRoutersInit(r *gin.Engine) {
	PodsRouters := r.Group("/api")
	{
		PodsRouters.GET("/k8spodsinfo", pods.PodsControllers{}.ResponsePodsListController)
		PodsRouters.POST("/deletepod", pods.PodsControllers{}.DeletePodController)
	}
}
