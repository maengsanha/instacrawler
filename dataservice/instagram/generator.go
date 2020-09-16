package instagram

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maengsanha/instacrawler/model/instagram"
)

const (
	requestPrefix  string = "https://www.instagram.com/explore/tags/"
	requestSuffix  string = "/?__a=1"
	endpointPrefix string = "&max_id="
)

// PageParserGenerator generates an Instagram page source parser.
func PageParserGenerator(query, cache string) func() ([]instagram.Post, string, error) {
	var (
		hasNextPage = true
		endCursor   = cache
	)

	return func() ([]instagram.Post, string, error) {
		if !hasNextPage {
			return nil, endCursor, fmt.Errorf("reached end of pagination")
		}

		var (
			tagPage    instagram.TagPage
			requestURL string
		)

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
			return nil, endCursor, fmt.Errorf("request failed with status code %d", resp.StatusCode)
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
