package routers

import (
	"gok8s/controllers/admin"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", admin.ApiResponse{}.Default)
		apiRouters.GET("/apiDetails", admin.ApiResponse{}.Details)
		// Add DELETE /api/deletepod route
		//apiRouters.DELETE("/deletepod", pods.DecodeDeletePodController)
	}

}
