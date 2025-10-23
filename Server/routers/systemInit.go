package routers

import (
	"gok8s/controllers/system"

	"github.com/gin-gonic/gin"
)

func SystemControllerInit(r *gin.Engine) {
	SystemRouters := r.Group("/")
	{
		SystemRouters.GET("/systeminfo", system.SystemController{}.SystemControllers)
	}
}
