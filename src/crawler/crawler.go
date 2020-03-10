package crawler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joshua-dev/instacrawler/src/checker"
)

const requestBaseURL = "https://instagram.com/explore/tags/"

// Crawl returns a Posting slice containing the URL of each post, cover photo URL, tag information, and number of likes found in the tag search by query.
func Crawl(query string) []Posting {
	return nil
}

// CrawlNumCoreSpriteHashtag returns the number of posts in the tag search url.
func CrawlNumCoreSpriteHashtag(query string) int {
	var requestURL string = requestBaseURL + query

	resp, err := http.Get(requestURL)
	checker.CheckError(err)
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checker.CheckError(err)

	parsedData := strings.Split(doc.Text(), `"edge_hashtag_to_media":{"count":`)[1]
	numCoreSpriteHashtag, err := strconv.Atoi(strings.Split(parsedData, ",")[0])
	checker.CheckError(err)

	return numCoreSpriteHashtag
}

// CoreSpriteHashtag returns a Posting struct containing cover photo URL, tag information, and number of likes
func CoreSpriteHashtag(url string) Posting {
	return Posting{}
}

// CoreSpriteHashtagSrc returns cover photo URL.
func CoreSpriteHashtagSrc(url string) string {
	return ""
}

// CoreSpriteHashtagTagsInfo returns hashtag informations.
func CoreSpriteHashtagTagsInfo(url string) string {
	return ""
}

// CoreSpriteHashtagLike returns number of likes.
func CoreSpriteHashtagLike(url string) int {
	return 0
}
