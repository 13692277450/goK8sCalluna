package system

import (
	"fmt"
	"gok8s/serverServices"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NetworkInfoController(c *gin.Context) {
	networkInfoResult, err := serverServices.NewK8sNetworkCollector()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Networkinfo result:**********", networkInfoResult)
	c.JSON(http.StatusOK, networkInfoResult)
}
