package routers

import (
	"net/http"

	"github.com/joshua-dev/instacrawler/src/controllers/crawler"

	"github.com/gin-gonic/gin"
	"github.com/joshua-dev/instacrawler/src/controllers/meta"
	"github.com/joshua-dev/instacrawler/src/controllers/searcher"
)

// HandleTopSearch handles /api/v1/topsearch
func HandleTopSearch(c *gin.Context) {
	query := c.Query("query")
	s := searcher.New()
	hashtags, err := s.TopSearch(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": hashtags,
	})
}

// HandleCrawl handles /api/v1/crawl
func HandleCrawl(c *gin.Context) {
	var req crawler.Request
	c.BindJSON(&req)
	secondLayer := req.SecondLayer
	thirdLayer := req.ThirdLayer
	resp, err := meta.Search(secondLayer, thirdLayer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
