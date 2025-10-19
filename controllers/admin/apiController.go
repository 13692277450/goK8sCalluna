package admin

import "github.com/gin-gonic/gin"

type ApiResponse struct{}

func (con ApiResponse) Details(c *gin.Context) {
	c.String(200, "api detail")
}
func (con ApiResponse) Default(c *gin.Context) {
	c.String(200, "api default")
}
