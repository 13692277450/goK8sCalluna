package routers

import (
	"gok8s/controllers/system"

	"github.com/gin-gonic/gin"
)

func NamespaceRoutersInit(r *gin.Engine) {
	SystemRouters := r.Group("/api")
	{
		SystemRouters.GET("/namespaceinfo", system.NamspaceController{}.NamespaceController)
	}
}
