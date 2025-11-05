package routers

import (
	"gok8s/controllers/logs"

	"github.com/gin-gonic/gin"
)

func PodsLogRoutersInit(r *gin.Engine) {

	podsLogRouters := r.Group("/api")
	{
		podsLogRouters.GET("/pods/logs", logs.PodLogsController{}.PodsLogController)
	}

}
