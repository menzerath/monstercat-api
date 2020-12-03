package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DownloadRelease(t *testing.T) {
	client := NewClient()

	err := client.DownloadRelease(Release{}, ReleaseDownloadFormatMP3, "./file.zip")
	assert.Error(t, err)
	assert.Equal(t, ErrorClientNotLoggedIn, err)
}
