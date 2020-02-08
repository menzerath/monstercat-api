package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelease_Download(t *testing.T) {
	data, err := Release{}.Download("", ReleaseDownloadFormatMP3)
	assert.Error(t, err)
	assert.Empty(t, data)
}
