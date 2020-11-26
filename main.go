package main

import (
	"github.com/maengsanha/instacrawler/middleware/greet"
	"github.com/maengsanha/instacrawler/middleware/meta"

	"github.com/gin-gonic/gin"
)

const metaPathPrefix = "/api/meta"

func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	api := engine.Group("/api")
	{
		api.POST(metaPathPrefix, meta.Search())
	}

	engine.GET("/", greet.Greet())

	engine.Run(":3000")
}
