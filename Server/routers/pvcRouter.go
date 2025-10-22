package routers

import (
	"gok8s/controllers/admin"

	"github.com/gin-gonic/gin"
)

func PVCControllerInit(r *gin.Engine) {
	pvcRouters := r.Group("/default")
	{
		// defaultRouters.GET("/", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "Hello default ",
		// 	})
		// })
		//loginRouters.GET("/login", (&loginControllers.LoginController{}).Login)
		pvcRouters.GET("/index", admin.PVCController{}.Home)

	}

}
