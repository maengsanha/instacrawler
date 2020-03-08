package crawler

import "fmt"

type hashtag struct {
	Name                 string `json:"name,string,omitempty"`
	ID                   int    `json:"id,omitempty"`
	MediaCount           int    `json:"media_count,omitempty"`
	ProfilePicURL        string `json:"profile_pic_url,string,omitempty"`
	SearchResultSubtitle string `json:"search_result_subtitle,omitempty"`
}

type hashtags struct {
	Position int     `json:"position,omitempty"`
	Hashtag  hashtag `json:"hashtag,omitempty"`
}

type topSearchResult struct {
	Users    []interface{} `json:"users,omitempty"`
	Places   []interface{} `json:"places,omitempty"`
	Hashtags hashtags      `json:"hashtags,omitempty"`
}

func (h hashtag) String() string {
	return fmt.Sprintf("%s: %s", h.Name, h.SearchResultSubtitle)
}
