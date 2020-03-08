package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func checkError(err error) {
	if err != nil {
		// log.Println(err)
		fmt.Println(err)
	}
}

func checkStatusCode(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		// log.Printf("Request failed with StatusCode: %d\n", resp.StatusCode)
		fmt.Printf("Request failed with StatusCode: %d\n", resp.StatusCode)
	}
}

// TopSearch returns the list of tags and the number of posts of typing query
// in the search field.
func TopSearch(query string) {
	const requestBaseURL = "https://www.instagram.com/web/search/topsearch/?context=blended&query="
	const reelOption = "include_reel=true"

	var requestURL string = fmt.Sprintf("%s%s&%s", requestBaseURL, url.QueryEscape(query), reelOption)

	resp, err := http.Get(requestURL)
	checkError(err)
	checkStatusCode(resp)

	defer resp.Body.Close() // prevent memory lick

	var result topSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	checkError(err)

	fmt.Println(result)
}
