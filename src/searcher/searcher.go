package searcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/joshua-dev/instacrawler/src/checker"
)

const requestBaseURL = "https://www.instagram.com/web/search/topsearch/?context=blended&query="
const reelOption = "include_reel=true"

// TopSearch returns the list of tags and the number of posts of typing query
// in the search field.
func TopSearch(query string) Hashtags {
	var requestURL string = fmt.Sprintf("%s%s&%s", requestBaseURL, url.QueryEscape(query), reelOption)

	resp, err := http.Get(requestURL)
	checker.CheckError(err)
	checker.CheckStatusCode(resp)

	defer resp.Body.Close()

	var result topSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	checker.CheckError(err)

	return result.Hashtags
}
