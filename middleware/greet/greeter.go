package greet

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Greet() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Hello, %s!", name)
	}
}
