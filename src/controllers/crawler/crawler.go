// Package crawler implements Instagram crawling logic.
package crawler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	jsoniter "github.com/json-iterator/go"

	"github.com/joshua-dev/instacrawler/src/controllers/checker"
	"github.com/joshua-dev/instacrawler/src/core"
)

const requestPrefix string = "https://www.instagram.com/explore/tags/"

var json jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary

// Crawler is an Instagram crawler type.
type Crawler struct {
	count       int
	endCursor   string
	hasNextPage bool
	InstaPosts  []core.InstaPost
}

// Request is a crawling request body type.
type Request struct {
	SecondLayer []string `json:"second_layer,omitempty"`
	ThirdLayer  []string `json:"third_layer,omitempty"`
	EndCursor   string   `json:"end_cursor"`
}

// Response is a crawling response body type.
type Response struct {
	SecondLayer []core.InstaPost `json:"second_layer,omitempty"`
	ThirdLayer  []core.InstaPost `json:"third_layer,omitempty"`
	EndCursor   string           `json:"end_cursor"`
}

// New returns a new Crawler.
func New() *Crawler {
	return &Crawler{}
}

// String implements fmt.Stringer interface.
func (c *Crawler) String() string {
	return fmt.Sprintf("count: %d\nend_cursor: %s\nhas_next_page: %t", c.count, c.endCursor, c.hasNextPage)
}

// init crawls page source from first Instagram hastag explore page with a given query.
func (c *Crawler) init(query string) error {
	var requestURL string = fmt.Sprintf("%s%s/?__a=1", requestPrefix, url.QueryEscape(query))

	resp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = c.update(data); err != nil {
		return err
	}

	return nil
}

// next parses json struct from next pagination.
func (c *Crawler) next(query string) error {
	if !c.hasNextPage {
		return errors.New("Reach end of GraphQL endpoint")
	}
	var requestURL string = fmt.Sprintf("%s%s/?__a=1&max_id=%s", requestPrefix, query, c.endCursor)

	resp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = c.update(data); err != nil {
		return err
	}

	return nil
}

// update updates crawler with a given json.
func (c *Crawler) update(data []byte) error {
	var syncer sync.WaitGroup
	var tagPage core.TagPage

	if err := json.Unmarshal(data, &tagPage); err != nil {
		return err
	}

	c.count = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Count
	c.endCursor = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
	c.hasNextPage = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.HasNextPage

	for _, edge := range tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges {
		syncer.Add(1)

		go func(c *Crawler, edge core.Edge) {
			defer syncer.Done()
			instaPost := core.InstaPost{
				ID:   edge.Node.ID,
				URL:  edge.Node.Shortcode,
				SRC:  edge.Node.DisplayURL,
				Like: edge.Node.EdgeLikedBy.Count,
			}
			if len(edge.Node.EdgeMediaToCaption.Edges) > 0 {
				instaPost.Text = edge.Node.EdgeMediaToCaption.Edges[0].Node.Text
			}
			c.InstaPosts = append(c.InstaPosts, instaPost)
		}(c, edge)
	}

	return nil
}

// Crawl implements crawling on Instagram with init and repeated next.
func (c *Crawler) Crawl(query string) {
	err := c.init(query)
	checker.CheckError(err)

	for c.hasNextPage {
		err = c.next(query)
		checker.CheckError(err)
	}
}
