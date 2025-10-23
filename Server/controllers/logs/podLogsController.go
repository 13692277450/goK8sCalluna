package logs

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type PodLogsController struct {
}

func (pl PodLogsController) PodsLogController(c *gin.Context) {
	var podsLogInfo = kubernetsServ.GetLogsFromMultiPods("default", "")
	c.JSON(200, gin.H{
		"data":   podsLogInfo,
		"status": "success",
	})
}
