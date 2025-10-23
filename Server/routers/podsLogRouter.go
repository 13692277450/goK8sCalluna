package routers

import (
	"gok8s/controllers/logs"

	"github.com/gin-gonic/gin"
)

func PodsLogRouterInit(r *gin.Engine) {

	podsLogRouters := r.Group("/")
	{
		podsLogRouters.GET("/api/pods/logs", logs.PodLogsController{}.PodsLogController)
		podsLogRouters.GET("/podsLogs.html", logs.PodLogsController{}.PodsLogController)
	}

}
