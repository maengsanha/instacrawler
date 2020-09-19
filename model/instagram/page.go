package instagram

// TagPage represents a single Instagram page source.
type TagPage struct {
	GraphQL struct {
		Hashtag struct {
			EdgeHashtagToMedia struct {
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						ID                 string `json:"id"`
						EdgeMediaToCaption struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						Shortcode   string `json:"shortcode"`
						EdgeLikedBy struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						ThumbnailSRC string `json:"thumbnail_src"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_hashtag_to_media"`
		} `json:"hashtag"`
	} `json:"graphql"`
}
