package system

import (
	"gok8s/kubernetsServ"
	"gok8s/models"

	"github.com/gin-gonic/gin"
)

type NamspaceController struct {
}

func (sc NamspaceController) NamespaceController(c *gin.Context) {
	var _ = models.NameSpaces{}
	var getNamespaceInfo = kubernetsServ.GetNameSpaceList()
	c.JSON(200, getNamespaceInfo)
}
