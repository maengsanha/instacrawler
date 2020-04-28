// Package crawler implements Instagram crawling logic.
package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joshua-dev/instacrawler/src/core"
)

const (
	requestPrefix  string = "https://www.instagram.com/explore/tags/"
	requestSuffix  string = "/?__a=1"
	endpointPrefix string = "&max_id="
)

// Request is a crawling request body type.
type Request struct {
	SecondLayer      []string `json:"second_layer"`
	ThirdLayer       []string `json:"third_layer"`
	SecondLayerCache []string `json:"second_layer_cache"`
	ThirdLayerCache  []string `json:"third_layer_cache"`
}

// Response is a crawling response body type.
type Response struct {
	SecondLayer      []core.InstaPost `json:"second_layer"`
	ThirdLayer       []core.InstaPost `json:"third_layer"`
	SecondLayerCache []string         `json:"second_layer_cache"`
	ThirdLayerCache  []string         `json:"third_layer_cache"`
}

// Generator generates an Instagram crawler.
func Generator(query, cache string) func() ([]core.InstaPost, string, error) {
	var hasNextPage bool = true
	var endCursor string = cache
	return func() ([]core.InstaPost, string, error) {
		if !hasNextPage {
			return nil, endCursor, fmt.Errorf("Reached end of pagination")
		}
		var tagPage core.TagPage
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

		posts := make([]core.InstaPost, len(tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges))

		for idx, edge := range tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges {
			post := core.InstaPost{
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
