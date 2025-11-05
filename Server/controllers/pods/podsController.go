package pods

import (
	"encoding/json"
	"gok8s/handlers"
	"gok8s/kubernetsServ"
	"gok8s/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PodsControllers struct{}

func (p PodsControllers) ResponsePodsListController(c *gin.Context) {
	c.JSON(200, kubernetsServ.GetK8sPods())
}

func (p PodsControllers) DeletePodController(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read Request body in delete pod handler",
			"details": err.Error(),
		})
		return
	}
	if len(bodyBytes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Empty pod content",
			"details": "The pod content cannot be empty",
		})
		return
	}
	err = json.Unmarshal(bodyBytes, &models.Request)
	if models.Request.PodName == "" {
		c.JSON(400, gin.H{"error": "podName is required"})
		return
	}
	if models.Request.NameSpace == "" {
		c.JSON(400, gin.H{"error": "namespace is required"})
		return
	}

	handlers.DeletePodHandler(c, models.Request.PodName, models.Request.NameSpace)

}
