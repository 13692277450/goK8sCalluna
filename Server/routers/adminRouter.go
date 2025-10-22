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
		// pvcController, err := admin.PVCController()
		// if err != nil {
		// 	panic(err)
		// }
		// adminRouters.GET("/pvcs", func(c *gin.Context) {
		// 	pvcs, err := PVCController.GetPVCList(c.Request.Context(), "default")
		// 	if err != nil {
		// 		c.JSON(500, gin.H{"error": err.Error()})
		// 		return
		// 	}
		// 	c.JSON(200, pvcs)
		// })
	}
}
