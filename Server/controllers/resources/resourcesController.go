package resources

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type ResourcesController struct {
}

func (rc ResourcesController) GetResources(c *gin.Context) {
	c.JSON(200, kubernetsServ.GetK8sResources())
}
