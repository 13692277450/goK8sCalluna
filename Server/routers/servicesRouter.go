package routers

import (
	"gok8s/controllers/services"

	"github.com/gin-gonic/gin"
)

func ServicesRoutersInit(r *gin.Engine) {
	ServicesRouters := r.Group("/api")
	{
		ServicesRouters.GET("/servicesinfo", services.ServicesController)

	}
}
