package routers

import (
	"gok8s/controllers/metrics"

	"github.com/gin-gonic/gin"
)

func MetricsRoutersInit(r *gin.Engine) {
	metricsRouters := r.Group("/api/metrics")
	{
		// 创建控制器实例
		nodesController := &metrics.MetricsNodesControllers{}
		podsController := &metrics.MetricsPodsControllers{}

		// 注册路由
		metricsRouters.GET("/nodes", nodesController.MetricsNodesController)
		metricsRouters.GET("/pods", podsController.MetricsPodsController)
	}
}
