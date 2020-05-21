// Package core defines models for Instagram crawling.
package core

import "fmt"

// SpriteHashtag is a type of associated search term
// recommended when searching on top of Instagram.
type SpriteHashtag struct {
	Hashtags []struct {
		Hashtag Hashtag `json:"hashtag,omitempty"`
	} `json:"hashtags,omitempty"`
}

// Hashtag is an Instagram hashtag type.
type Hashtag struct {
	Name                 string `json:"name,omitempty"`
	MediaCount           int    `json:"media_count,omitempty"`
	SearchResultSubtitle string `json:"search_result_subtitle,omitempty"`
}

// String implements fmt.Stringer interface.
func (s SpriteHashtag) String() (str string) {
	for _, hashtag := range s.Hashtags {
		str += fmt.Sprintln(hashtag.Hashtag.String())
	}
	return
}

// String implements fmt.Stringer interface.
func (h Hashtag) String() string {
	return fmt.Sprintf("#%s %s %d", h.Name, h.SearchResultSubtitle, h.MediaCount)
}
