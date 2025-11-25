package metrics

import (
	"gok8s/serverServices"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetricsNodesControllers struct {
}
type MetricsPodsControllers struct{}

func (mn *MetricsNodesControllers) MetricsNodesController(c *gin.Context) {
	c.JSON(http.StatusOK, serverServices.K8sNodesPerformance())
}

func (mp *MetricsPodsControllers) MetricsPodsController(c *gin.Context) {
	c.JSON(http.StatusOK, serverServices.K8sPodsPerformance())
}
