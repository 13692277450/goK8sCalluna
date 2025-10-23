package routers

import (
	"gok8s/controllers/system"

	"github.com/gin-gonic/gin"
)

func NamespaceControllerInit(r *gin.Engine) {
	SystemRouters := r.Group("/")
	{
		SystemRouters.GET("/namespaceinfo", system.NamspaceController{}.NamespaceController)
	}
}
