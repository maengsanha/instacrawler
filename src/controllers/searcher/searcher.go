// Package searcher implements logics for crawling Instagram top search.
package searcher

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/joshua-dev/instacrawler/src/controllers/checker"
	"github.com/joshua-dev/instacrawler/src/core"
)

const (
	requestBaseURL string = "https://www.instagram.com/web/search/topsearch/?context=blended&query="
	reelOption     string = "include_reel=true"
)

// Searcher is a struct type that perfroms a search.
type Searcher struct {
	Hashtags core.Hashtags `json:"hashtags,omitempty"`
}

// New returns a new Searcher.
func New() Searcher {
	return Searcher{}
}

// TopSearch implements top search on Instagram with a given query.
func (s Searcher) TopSearch(query string) ([]byte, error) {
	var requestURL string = fmt.Sprintf("%s%s&%s", requestBaseURL, url.QueryEscape(query), reelOption)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	var hashtags []core.Hashtag

	for _, c := range s.Hashtags {
		hashtags = append(hashtags, core.Hashtag{
			Name:                 c.Hashtag.Name,
			MediaCount:           c.Hashtag.MediaCount,
			SearchResultSubtitle: strings.Split(c.Hashtag.SearchResultSubtitle, " ")[0],
		})
	}

	return json.Marshal(hashtags)
}
