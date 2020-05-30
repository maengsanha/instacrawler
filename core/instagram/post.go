// Package instagram defines models for Instagram crawling.
package instagram

import "fmt"

// Post is an Instagram post type.
type Post struct {
	ID   string `json:"-"`
	Text string `json:"text"`
	URL  string `json:"url,omitempty"`
	SRC  string `json:"src,omitempty"`
	Like int    `json:"like,omitempty"`
}

// PostMap is an Instagram post map type.
type PostMap map[string]Post

// String implements fmt.Stringer interface.
func (p Post) String() string {
	return fmt.Sprintf("%s\n%s\n%d\n%s\n", p.URL, p.SRC, p.Like, p.Text)
}

// TagPage is an Instagram json page source type.
type TagPage struct {
	GraphQL struct {
		Hashtag struct {
			EdgeHashtagToMedia struct {
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page,omitempty"`
					EndCursor   string `json:"end_cursor,omitempty"`
				} `json:"page_info,omitempty"`
				Edges []struct {
					Node struct {
						ID                 string `json:"id,omitempty"`
						EdgeMediaToCaption struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node,omitempty"`
							} `json:"edges,omitempty"`
						} `json:"edge_media_to_caption,omitempty"`
						Shortcode   string `json:"shortcode,omitempty"`
						EdgeLikedBy struct {
							Count int `json:"count,omitempty"`
						} `json:"edge_liked_by,omitempty"`
						ThumbnailSRC string `json:"thumbnail_src"`
					} `json:"node,omitempty"`
				} `json:"edges,omitempty"`
			} `json:"edge_hashtag_to_media,omitempty"`
		} `json:"hashtag,omitempty"`
	} `json:"graphql,omitempty"`
}
