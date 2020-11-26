package greet

import "github.com/gin-gonic/gin"

func Greet() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		c.String(200, "Hello, %s!", name)
	}
}
