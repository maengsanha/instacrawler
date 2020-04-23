// Package checker implements checking logics for error and http response.
package checker

import "net/http"

// CheckError reports whether a given error occurs.
func CheckError(err error) bool {
	return err != nil
}

// CheckStatusCode reports whether the status code of a given response is not 200 or 429.
func CheckStatusCode(resp *http.Response) bool {
	return resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTooManyRequests
}
