// Package instagram defines models for Instagram crawling.
package instagram

import "fmt"

// CoreSpriteHashtag is a type of associated search term
// recommended when searching on top of Instagram.
type CoreSpriteHashtag struct {
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
func (c CoreSpriteHashtag) String() (str string) {
	for _, hashtag := range c.Hashtags {
		str += fmt.Sprintln(hashtag.Hashtag.String())
	}
	return
}

// String implements fmt.Stringer interface.
func (h Hashtag) String() string {
	return fmt.Sprintf("#%s %s %d", h.Name, h.SearchResultSubtitle, h.MediaCount)
}
