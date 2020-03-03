package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://instagram.com/explore/tags/"

// Scrape scrapes results from searches on terms in Instagram.
func Scrape(term string) {
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
