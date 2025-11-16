package routers

import (
	"gok8s/controllers/system"
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

func SystemCenterRoutersInit(r *gin.Engine) {
	SystemRouters := r.Group("/api")
	{
		SystemRouters.GET("/systeminfo", system.SystemController{}.SystemControllers)
		kubernetsServ.GetNameSpaceList()

		// Register K8sManageCenterRoutersInit directly
		//K8sManageCenterRoutersInit(r)
	}
}
