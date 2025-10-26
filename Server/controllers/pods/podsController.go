package pods

import (
	"gok8s/kubernetsServ"

	"github.com/gin-gonic/gin"
)

type PodsControllers struct{}

// var podsJsonInfo models.PodInfo

func (p PodsControllers) ResponsePodsListController(c *gin.Context) {

	c.JSON(200, kubernetsServ.GetK8sPods())
}
