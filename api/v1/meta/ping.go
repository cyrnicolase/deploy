package meta

import (
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	c.String(200, "Hello")
}