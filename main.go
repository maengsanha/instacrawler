package main

import (
	"runtime"

	"github.com/maengsanha/instacrawler/middleware/meta"

	"github.com/gin-gonic/gin"
)

const metaPathPrefix = "/api/meta"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	api := engine.Group("/api")

	api.POST(metaPathPrefix, meta.Search())

	engine.Run(":3000")
}
