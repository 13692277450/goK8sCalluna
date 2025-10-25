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
