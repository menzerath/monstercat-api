package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseList(t *testing.T) {
	client := NewClient()

	releaseList, err := client.ReleaseList()
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.True(t, releaseList.HasNextPage())
}

func TestReleaseListAtPosition(t *testing.T) {
	client := NewClient()

	releaseList, err := client.ReleaseListAtPosition(5, 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.True(t, releaseList.HasNextPage())
}

func TestReleaseListUntilEnd(t *testing.T) {
	client := NewClient()

	releaseList, err := client.ReleaseList()
	assert.NoError(t, err)

	for releaseList.HasNextPage() {
		releaseList, err = client.ReleaseListAtPosition(releaseList.Limit, releaseList.Skip+releaseList.Limit)
		assert.NoError(t, err)
	}
}
