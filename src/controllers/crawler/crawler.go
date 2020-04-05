package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joshua-dev/instacrawler/src/controllers/checker"
	"github.com/joshua-dev/instacrawler/src/core"
)

const requestBaseURL string = "https://www.instagram.com/explore/tags/"

// Crawler is a struct type that perfroms crawling on Instagram.
type Crawler struct {
	Count       int    `json:"count,omitempty"`
	EndCursor   string `json:"end_cursor,omitempty"`
	HasNextPage bool   `json:"has_next_page,omitempty"`
	InstaPosts  []core.InstaPost
}

// New returns a new Crawler.
func New() *Crawler {
	return &Crawler{}
}

// String implements fmt.Stringer interface.
func (c *Crawler) String() string {
	var str string
	str += fmt.Sprintf("count: %d\n", c.Count)
	str += fmt.Sprintf("end_cursor: %s\n", c.EndCursor)
	str += fmt.Sprintf("has_next_page: %t\n\n", c.HasNextPage)
	for _, instaPost := range c.InstaPosts {
		str += fmt.Sprintf("%s\n", instaPost.String())
	}
	return str
}

// init crawls page source from first Instagram hastag explore page with a given query.
func (c *Crawler) init(query string) error {
	var requestURL string = fmt.Sprintf("%s%s/", requestBaseURL, url.QueryEscape(query))

	// fmt.Fprintf(os.Stdout, "Requesting to %s...\n", requestURL)

	resp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	parsedInstaSource := strings.Split((strings.Split(doc.Text(), `"TagPage":[`)[1]), `]},"hostname":`)[0]

	if err = c.update(parsedInstaSource); err != nil {
		return err
	}

	return nil
}

// next parses json struct from next pagination.
func (c *Crawler) next(query string) error {
	if !c.HasNextPage {
		return errors.New("Reach end of GraphQL endpoint")
	}
	var requestURL string = fmt.Sprintf("%s%s/?__a=1&max_id=%s", requestBaseURL, query, c.EndCursor)

	// fmt.Fprintf(os.Stdout, "Requesting to %s...\n", requestURL)

	resp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	if err = c.update(doc.Text()); err != nil {
		return err
	}

	return nil
}

// update updates c with a given json.
func (c *Crawler) update(jsonText string) error {
	tagPage := core.TagPage{}
	if err := json.Unmarshal([]byte(jsonText), &tagPage); err != nil {
		return err
	}
	c.Count = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Count
	c.EndCursor = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
	c.HasNextPage = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.HasNextPage

	for _, edge := range tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges {
		go c.collect(edge)
	}

	return nil
}

// collect appends an Instagram post to c with a given edge.
func (c *Crawler) collect(e core.Edge) {
	instaPost := core.InstaPost{
		ID:   e.Node.ID,
		URL:  e.Node.Shortcode,
		SRC:  e.Node.DisplayURL,
		Like: e.Node.EdgeLikedBy.Count,
	}
	if len(e.Node.EdgeMediaToCaption.InEdges) > 0 {
		instaPost.Text = e.Node.EdgeMediaToCaption.InEdges[0].InNode.Text
	}
	c.InstaPosts = append(c.InstaPosts, instaPost)
}

// Crawl completes crawling from Instagram through init and repeated next.
func (c *Crawler) Crawl(query string) {
	// use every CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	// fmt.Fprintf(os.Stdout, "CPU usage: %d\n", runtime.GOMAXPROCS(0))

	err := c.init(query)
	checker.CheckError(err)

	for c.HasNextPage && len(c.InstaPosts) < 1000 {
		err = c.next(query)
		checker.CheckError(err)
	}
}
