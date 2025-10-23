package routers

import (
	"gok8s/controllers/resources"

	"github.com/gin-gonic/gin"
)

func K8sResourcesInit(r *gin.Engine) {
	resourcesRouters := r.Group("/")
	{
		resourcesRouters.GET("/resourcesInfo.html", resources.ResourcesController{}.GetResources)
	}
}
