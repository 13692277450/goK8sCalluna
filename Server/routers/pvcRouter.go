package routers

import (
	"gok8s/controllers/admin/pvc"

	"github.com/gin-gonic/gin"
)

func PVCRoutersInit(r *gin.Engine) {
	pvcRouters := r.Group("/api")
	{
		// defaultRouters.GET("/", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "Hello default ",
		// 	})
		// })
		//loginRouters.GET("/login", (&loginControllers.LoginController{}).Login)
		pvcRouters.GET("/pvcinfo", pvc.PVCControllers{}.PVCController)

	}

}
