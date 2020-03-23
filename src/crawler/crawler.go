package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/joshua-dev/instacrawler/src/checker"
)

const requestBaseURL string = "https://www.instagram.com/explore/tags/"

// Crawler is a struct type that perfroms crawling on Instagram.
type Crawler struct {
	Count       int    `json:"count,omitempty"`
	EndCursor   string `json:"end_cursor,omitempty"`
	HasNextPage bool   `json:"has_next_page,omitempty"`
	InstaPosts  []InstaPost
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

	tagPage := TagPage{}

	if err = json.Unmarshal([]byte(parsedInstaSource), &tagPage); err != nil {
		return err
	}

	c.Count = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Count
	c.EndCursor = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
	c.HasNextPage = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.HasNextPage

	for _, edge := range tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges {
		// TODO: Concurrency
		instaPost := InstaPost{
			ID:   edge.Node.ID,
			URL:  edge.Node.Shortcode,
			SRC:  edge.Node.DisplayURL,
			Like: edge.Node.EdgeLikedBy.Count,
		}
		if len(edge.Node.EdgeMediaToCaption.InEdges) > 0 {
			instaPost.Text = edge.Node.EdgeMediaToCaption.InEdges[0].InNode.Text
		}
		c.InstaPosts = append(c.InstaPosts, instaPost)
	}

	return nil
}

// next parses json struct from next pagination.
func (c *Crawler) next(query string) error {
	if !c.HasNextPage {
		return errors.New("Reach end of GraphQL endpoint")
	}
	var requestURL string = fmt.Sprintf("%s%s/?__a=1&max_id=%s", requestBaseURL, query, c.EndCursor)

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

	tagPage := TagPage{}

	if err = json.Unmarshal([]byte(doc.Text()), &tagPage); err != nil {
		return err
	}

	c.Count = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Count
	c.EndCursor = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
	c.HasNextPage = tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.PageInfo.HasNextPage

	for _, edge := range tagPage.GraphQL.Hashtag.EdgeHashtagToMedia.Edges {
		// TODO: Concurrency
		instaPost := InstaPost{
			ID:   edge.Node.ID,
			URL:  edge.Node.Shortcode,
			SRC:  edge.Node.DisplayURL,
			Like: edge.Node.EdgeLikedBy.Count,
		}
		if len(edge.Node.EdgeMediaToCaption.InEdges) > 0 {
			instaPost.Text = edge.Node.EdgeMediaToCaption.InEdges[0].InNode.Text
		}
		c.InstaPosts = append(c.InstaPosts, instaPost)
	}

	return nil
}

// Crawl completes crawling from Instagram through init and repeated next.
func (c *Crawler) Crawl(query string) {
	err := c.init(query)
	checker.CheckError(err)

	for c.HasNextPage {
		err = c.next(query)
		checker.CheckError(err)
	}
}
