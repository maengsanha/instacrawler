package main

import (
	"github.com/maengsanha/instacrawler/middleware/meta"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	engine.POST("/api/meta", meta.Search())

	engine.Run(":3000")
}
