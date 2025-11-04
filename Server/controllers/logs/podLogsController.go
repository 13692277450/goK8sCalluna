package logs

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type PodLogsController struct {
}

// 在kubernetsServ包中添加

func (pl PodLogsController) PodsLogController(c *gin.Context) {

	var podsLogInfo = kubernetsServ.GetLogsFromMultiPods(kubernetsServ.NameSpacesTotal, "")
	c.JSON(200, gin.H{
		"data":   podsLogInfo,
		"status": "success",
	})
}
