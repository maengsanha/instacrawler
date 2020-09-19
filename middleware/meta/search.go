package meta

import (
	"net/http"

	metasearch "github.com/maengsanha/instacrawler/model/meta"
	"github.com/maengsanha/instacrawler/usecase/meta"

	"github.com/gin-gonic/gin"
)

// Search returns a handler for /api/meta.
func Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req metasearch.Request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, nil)
		} else if len(req.HigherLayer) == len(req.HigherLayerCache) && len(req.LowerLayer) == len(req.LowerLayerCache) {
			c.JSON(http.StatusOK, meta.Search(req.HigherLayer, req.LowerLayer, req.HigherLayerCache, req.LowerLayerCache))
		} else {
			c.JSON(http.StatusBadRequest, nil)
		}
	}
}
