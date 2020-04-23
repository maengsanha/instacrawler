// Package top implements the top search crawling logic on Instagram.
package top

import (
	"fmt"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"

	"github.com/joshua-dev/instacrawler/src/core"
)

const (
	requestPrefix string = "https://www.instagram.com/web/search/topsearch/?context=blended&query="
	reelOption    string = "include_reel=true"
)

var json jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary

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
