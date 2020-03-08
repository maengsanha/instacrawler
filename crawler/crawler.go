package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://instagram.com/explore/tags/"

type tag string

type tags []tag

// Posting is a posting of Instagram search result.
type Posting struct {
	URL string `json:"url,string,omitempty"`
	Img string `json:"image,string,omitempty"`
	Tag tags   `json:"tags,string,omitempty"`
}

// Crawl crawls results from searches on terms in Instagram.
func Crawl(term string) {
	var targetURL string = baseURL + term

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		// log.Println(err)
		fmt.Println(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	result := string(bytes)

	fmt.Println(result)
}
