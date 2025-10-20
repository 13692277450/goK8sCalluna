package middlewares

import "github.com/gin-gonic/gin"

func MiddleWaresInit(c *gin.Context) {
	// middlewares init
	c.Set("Username", "Jacky")
	c.Next()
}
