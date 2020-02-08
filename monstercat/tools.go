package monstercat

import (
	"net/http"
	"time"
)

// getHTTPClient returns a default HTTP client
func getHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}
