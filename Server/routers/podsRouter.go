package routers

import (
	"gok8s/controllers/pods"

	"github.com/gin-gonic/gin"
)

func K8sManageCenterInit(r *gin.Engine) {
	PodsRouters := r.Group("/")
	{
		PodsRouters.GET("/k8spodlist.html", pods.PodsController{}.ResponsePodsListController)
	}
}
