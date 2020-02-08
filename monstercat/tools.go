package monstercat

import (
	"net/http"
)

// getHTTPClient returns a default HTTP client
func getHTTPClient() *http.Client {
	return &http.Client{}
}
