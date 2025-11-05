package routers

import (
	"gok8s/controllers/admin"

	"github.com/gin-gonic/gin"
)

func LoginRoutersInit(r *gin.Engine) {
	loginRouters := r.Group("/")
	{
		// defaultRouters.GET("/", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "Hello default ",
		// 	})
		// })
		//loginRouters.GET("/login", (&loginControllers.LoginController{}).Login)
		loginRouters.GET("/index", admin.AdminController{}.Home)

	}

}
