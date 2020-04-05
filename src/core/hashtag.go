package core

import "fmt"

// Hashtags is a list of related hashtags type
// provided when searching at the top of Instagram.
type Hashtags []content

type content struct {
	Hashtag Hashtag `json:"hashtag,omitempty"`
}

// Hashtag is an Instagram hashtag type.
type Hashtag struct {
	Name                 string `json:"name,omitempty"`
	MediaCount           int    `json:"media_count,omitempty"`
	SearchResultSubtitle string `json:"search_result_subtitle,omitempty"`
}

// String implements fmt.Stringer interface.
func (h Hashtag) String() string {
	return fmt.Sprintf("#%s: %s(%d)", h.Name, h.SearchResultSubtitle, h.MediaCount)
}

// String implements fmt.Stringer interface.
func (h Hashtags) String() string {
	var str string
	for _, c := range h {
		str += fmt.Sprintf("%s\n", c.Hashtag.String())
	}
	return str
}
