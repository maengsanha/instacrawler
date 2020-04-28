// Package top implements the Instagram top search crawling logic.
package top

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/joshua-dev/instacrawler/src/core"
)

const (
	requestPrefix string = "https://www.instagram.com/web/search/topsearch/?context=blended&query="
	reelOption    string = "include_reel=true"
)

// Search implements top search on Instagram with a given query.
func Search(query string) (hashtags []core.Hashtag, err error) {
	var requestURL string = fmt.Sprintf("%s%s&%s", requestPrefix, url.QueryEscape(query), reelOption)

	resp, err := http.Get(requestURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var coreSpriteHashtag core.SpriteHashtag
	if err = json.NewDecoder(resp.Body).Decode(&coreSpriteHashtag); err != nil {
		return
	}

	for _, hashtag := range coreSpriteHashtag.Hashtags {
		hashtags = append(hashtags, hashtag.Hashtag)
	}
	return
}
