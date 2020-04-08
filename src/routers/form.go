// Package routers implements Instagram crawler api routers.
package routers

import (
	"github.com/joshua-dev/instacrawler/src/core"
)

// TopSearchRequest is a top search request body type.
// path: /api/v1/topsearch
type TopSearchRequest struct {
	Query string `json:"query,omitempty"`
}

// CrawlRequest is a crawling request body type.
type CrawlRequest struct {
	SecondLayer []string `json:"second_layer,omitempty"`
	ThirdLayer  []string `json:"third_layer,omitempty"`
}

// CrawlResponse is a crawling response body type.
type CrawlResponse struct {
	SecondLayer []core.InstaPost `json:"second_layer,omitempty"`
	ThirdLayer  []core.InstaPost `json:"third_layer,omitempty"`
}
