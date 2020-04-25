package routers

import (
	"net/http"

	"github.com/joshua-dev/instacrawler/src/controllers/crawler"
	"github.com/joshua-dev/instacrawler/src/controllers/top"

	"github.com/gin-gonic/gin"
	"github.com/joshua-dev/instacrawler/src/controllers/meta"
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
	resp := meta.Search(secondLayer, thirdLayer)
	c.JSON(http.StatusOK, resp)
}