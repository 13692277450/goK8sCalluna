package routers

import (
	"gok8s/controllers/resources"

	"github.com/gin-gonic/gin"
)

func K8sResourcesRoutersInit(r *gin.Engine) {
	resourcesRouters := r.Group("/api")
	{
		resourcesRouters.GET("/resourcesinfo", resources.ResourcesController{}.GetResources)
	}
}
