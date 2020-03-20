package crawler

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"

	"github.com/joshua-dev/instacrawler/src/checker"
)

const requestBaseURL string = "https://www.instagram.com/explore/tags/"

// Crawler is a struct type that perfroms crawling on Instagram.
type Crawler struct {
	Postings    []Posting `json:"postings,omitempty"`
	endCursor   string
	hasNextPage bool
}

// New returns a new Crawler.
func New() *Crawler {
	return &Crawler{}
}

// Init crawls page source from first Instagram hastag explore page with a given query.
func (c *Crawler) Init(query string) error {
	var requestURL string = requestBaseURL + url.QueryEscape(query) + "/"

	resp, err := http.Get(requestURL)
	checker.CheckError(err)
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checker.CheckError(err)

	return nil
}

// Next parses json struct from next Instagram hashtag explore page.
func (c *Crawler) Next() error {
	return nil
}
