package crawler

import "fmt"

// InstaPost is an Instagram post type.
type InstaPost struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text"`
	URL  string `json:"url,omitempty"`
	SRC  string `json:"src,omitempty"`
	Like int    `json:"like,omitempty"`
}

// String implements fmt.Stringer interface.
func (i InstaPost) String() string {
	return fmt.Sprintf("ID: %s\nText: %s\nURL: https://www.instagram.com/p/%s/\nSRC: %s\nLikes: %d\n", i.ID, i.Text, i.URL, i.SRC, i.Like)
}

// TagPage is a GraphQL endpoint json type.
type TagPage struct {
	GraphQL graphQL `json:"graphql,omitempty"`
}

type graphQL struct {
	Hashtag hashtag `json:"hashtag,omitempty"`
}

type hashtag struct {
	EdgeHashtagToMedia edgeHashtagToMedia `json:"edge_hashtag_to_media,omitempty"`
}

type edgeHashtagToMedia struct {
	Count    int      `json:"count,omitempty"`
	PageInfo pageInfo `json:"page_info,omitempty"`
	Edges    []edge   `json:"edges,omitempty"`
}

type pageInfo struct {
	HasNextPage bool   `json:"has_next_page,omitempty"`
	EndCursor   string `json:"end_cursor,omitempty"`
}

type edge struct {
	Node node `json:"node,omitempty"`
}

type node struct {
	ID                 string             `json:"id,omitempty"`
	EdgeMediaToCaption edgeMediaToCaption `json:"edge_media_to_caption,omitempty"`
	Shortcode          string             `json:"shortcode,omitempty"`
	DisplayURL         string             `json:"display_url,omitempty"`
	EdgeLikedBy        edgeLikedBy        `json:"edge_liked_by,omitempty"`
}

type edgeMediaToCaption struct {
	InEdges []inEdge `json:"edges,omitempty"`
}

type inEdge struct {
	InNode inNode `json:"node,omitempty"`
}

type inNode struct {
	Text string `json:"text"`
}

type edgeLikedBy struct {
	Count int `json:"count,omitempty"`
}