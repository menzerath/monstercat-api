package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DownloadRelease(t *testing.T) {
	client := NewClient()

	err := client.DownloadRelease(CatalogItem{}, ReleaseDownloadFormatMP3, "./file.mp3")
	assert.Error(t, err)
	assert.Equal(t, ErrorClientNotLoggedIn, err)
}
