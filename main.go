// Package main runs Instagram crawler api server.
package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joshua-dev/instacrawler/routers"
)

const (
	topSearchPathPrefix string = "/api/v1/topsearch"
	crawlPathPrefix     string = "/api/v1/crawl"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	engine.GET(topSearchPathPrefix, routers.HandleTopSearch)

	engine.POST(crawlPathPrefix, routers.HandleCrawl)

	engine.Run(":3000")
}