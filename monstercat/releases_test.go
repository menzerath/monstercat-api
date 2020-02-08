package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DownloadRelease(t *testing.T) {
	client := NewClient()

	data, err := client.DownloadRelease(Release{}, ReleaseDownloadFormatMP3)
	assert.Error(t, err)
	assert.Equal(t, ErrorClientNotLoggedIn, err)
	assert.Empty(t, data)
}
