package deploy

import (
	"gok8s/handlers"

	"github.com/gin-gonic/gin"
)

type DeployNamespaceControllers struct {
}

func (dn DeployNamespaceControllers) DeployNamespaceController(c *gin.Context) {
	handlers.DeployNamespaceHandler(c.Writer, c.Request)
}
