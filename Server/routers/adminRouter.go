package routers

import (
	"gok8s/controllers/admin"
	"gok8s/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.MiddleWaresInit)
	{
		adminRouters.GET("/", admin.AdminController{}.Home)
		adminRouters.GET("/account", admin.AdminController{}.Account)
		adminRouters.GET("/add", admin.AdminController{}.Add)
		adminRouters.GET("/home", admin.AdminController{}.Home)

	}
}
