package monstercat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient(t *testing.T) {
	httpClient := getHTTPClient()
	assert.NotNil(t, httpClient)
	assert.Equal(t, 5*time.Second, httpClient.Timeout)
}
