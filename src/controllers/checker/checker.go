// Package checker implements checking logics for error and http response.
package checker

import (
	"log"
	"net/http"
)

// CheckError reports whether a given error occurs.
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// CheckStatusCode reports whether the status code of a given response is not 200.
func CheckStatusCode(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with StatusCode: %d\n", resp.StatusCode)
	}
}
