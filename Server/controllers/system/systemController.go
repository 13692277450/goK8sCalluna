package system

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
}

func (sc SystemController) SystemControllers(c *gin.Context) {
	var getNodesInfor = kubernetsServ.GetNodesInfo()
	c.JSON(200, getNodesInfor)
}

func (sc SystemController) K8sSystemLogsController(c *gin.Context) {
	var getNodesInfor = kubernetsServ.GetK8sPods()
	c.JSON(200, getNodesInfor)
}
