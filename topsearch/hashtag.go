package topsearch

import "fmt"

type topSearchResult struct {
	Hashtags Hashtags `json:"hashtags,omitempty"`
}

// Hashtags is a content slice type.
type Hashtags []content

type content struct {
	Hashtag hashtag `json:"hashtag,omitempty"`
}

type hashtag struct {
	Name                 string `json:"name,omitempty"`
	MediaCount           int    `json:"media_count,omitempty"`
	SearchResultSubtitle string `json:"search_result_subtitle,omitempty"`
}

// String implements fmt.Stringer interface.
func (h hashtag) String() string {
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
