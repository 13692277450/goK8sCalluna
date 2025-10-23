package system

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type NamspaceController struct {
}

func (sc NamspaceController) NamespaceController(c *gin.Context) {
	var getNamespaceInfo = kubernetsServ.GetNameSpaceList()
	c.JSON(200, getNamespaceInfo)
}
