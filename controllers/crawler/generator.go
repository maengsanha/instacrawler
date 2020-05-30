// Package crawler implements Instagram crawling logic.
package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joshua-dev/instacrawler/core/instagram"
)

const (
	requestPrefix  string = "https://www.instagram.com/explore/tags/"
	requestSuffix  string = "/?__a=1"
	endpointPrefix string = "&max_id="
)

// Request is a crawling request body type.
type Request struct {
	HigherLayer      []string `json:"higher_layer"`
	LowerLayer       []string `json:"lower_layer"`
	HigherLayerCache []string `json:"higher_layer_cache"`
	LowerLayerCache  []string `json:"lower_layer_cache"`
}

// Response is a crawling response body type.
type Response struct {
	HigherLayer      []instagram.Post `json:"higher_layer"`
	LowerLayer       []instagram.Post `json:"lower_layer"`
	HigherLayerCache []string         `json:"higher_layer_cache"`
	LowerLayerCache  []string         `json:"lower_layer_cache"`
}

// Generator generates an Instagram crawler.
func Generator(query, cache string) func() ([]instagram.Post, string, error) {
	var hasNextPage bool = true
	var endCursor string = cache

	return func() ([]instagram.Post, string, error) {
		if !hasNextPage {
			return nil, endCursor, fmt.Errorf("Reached end of pagination")
		}

		var tagPage instagram.TagPage
		var requestURL string

		if endCursor == "" {
			requestURL = fmt.Sprintf("%s%s%s", requestPrefix, query, requestSuffix)
		} else {
			requestURL = fmt.Sprintf("%s%s%s%s%s", requestPrefix, query, requestSuffix, endpointPrefix, endCursor)
		}

		resp, err := http.Get(requestURL)
		if err != nil {
			return nil, endCursor, err
		}
		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTooManyRequests {
			return nil, endCursor, fmt.Errorf("Request failed with status code %d", resp.StatusCode)
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(&tagPage); err != nil {
			return nil, endCursor, err
		}

		posts := make([]instagram.Post, len(tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges))

		for idx, edge := range tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges {
			post := instagram.Post{
				ID:   edge.Node.ID,
				URL:  edge.Node.Shortcode,
				SRC:  edge.Node.ThumbnailSRC,
				Like: edge.Node.EdgeLikedBy.Count,
			}
			if len(edge.Node.EdgeMediaToCaption.Edges) > 0 {
				post.Text = edge.Node.EdgeMediaToCaption.Edges[0].Node.Text
			}
			posts[idx] = post
		}

		hasNextPage = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.HasNextPage
		endCursor = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
		return posts, endCursor, nil
	}
}
