package services

import (
	"gok8s/handlers"

	"github.com/gin-gonic/gin"
)

func ServicesController(c *gin.Context) {
	//var getServicesResult = handlers.GetServicesHandler
	c.JSON(200, handlers.GetServicesHandler())
}
