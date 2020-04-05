package searcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/joshua-dev/instacrawler/src/controllers/checker"
	"github.com/joshua-dev/instacrawler/src/core"
)

const (
	requestBaseURL = "https://www.instagram.com/web/search/topsearch/?context=blended&query="
	reelOption     = "include_reel=true"
)

// Searcher is a struct type that perfroms a search.
type Searcher struct {
	Hashtags core.Hashtags `json:"hashtags,omitempty"`
}

// New returns a new Searcher.
func New() *Searcher {
	return &Searcher{}
}

// TopSearch performs a top search in Instagram with a given query.
func (s *Searcher) TopSearch(query string) error {
	var requestURL string = fmt.Sprintf("%s%s&%s", requestBaseURL, url.QueryEscape(query), reelOption)

	resp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(s)
	if err != nil {
		return err
	}

	return nil
}
