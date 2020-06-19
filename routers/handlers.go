package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua-dev/instacrawler/controllers/meta"
	"github.com/joshua-dev/instacrawler/controllers/top"
)

// HandleTopSearch handles /api/v1/topsearch.
func HandleTopSearch(c *gin.Context) {
	query := c.Query("query")

	hashtags, err := top.Search(query)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, hashtags)
}

// HandleCrawl handles /api/v1/crawl.
func HandleCrawl(c *gin.Context) {
	var req meta.Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else if len(req.HigherLayer) == len(req.HigherLayerCache) && len(req.LowerLayer) == len(req.LowerLayerCache) {
		c.JSON(http.StatusOK, meta.Search(req.HigherLayer, req.LowerLayer, req.HigherLayerCache, req.LowerLayerCache))
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}
