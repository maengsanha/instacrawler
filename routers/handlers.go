package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua-dev/instacrawler/controllers/crawler"
	"github.com/joshua-dev/instacrawler/controllers/meta"
	"github.com/joshua-dev/instacrawler/controllers/top"
)

// HandleTopSearch handles /api/v1/topsearch
func HandleTopSearch(c *gin.Context) {
	query := c.Query("query")

	hashtags, err := top.Search(query)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, hashtags)
}

// HandleCrawl handles /api/v1/crawl
func HandleCrawl(c *gin.Context) {
	var req crawler.Request
	c.BindJSON(&req)

	secondLayer := req.SecondLayer
	thirdLayer := req.ThirdLayer
	secondLayerCache := req.SecondLayerCache
	thirdLayerCache := req.ThirdLayerCache

	if len(secondLayer) == len(secondLayerCache) && len(thirdLayer) == len(thirdLayerCache) {
		c.JSON(http.StatusOK, meta.Search(secondLayer, thirdLayer, secondLayerCache, thirdLayerCache))
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}
