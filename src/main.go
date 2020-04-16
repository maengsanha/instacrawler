// Package main runs Instagram crawler api server.
package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joshua-dev/instacrawler/src/routers"
)

const (
	topSearchPrefix string = "/api/v1/topsearch"
	crawlPrefix     string = "/api/v1/crawl"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := gin.Default()

	router.GET(topSearchPrefix, routers.HandleTopSearch)

	router.POST(crawlPrefix, routers.HandleCrawl)

	router.Run()
}
