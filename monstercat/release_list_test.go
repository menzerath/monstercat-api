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

func TestReleases(t *testing.T) {
	client := NewClient()

	releaseList, err := client.Releases("", "", 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.NotEqual(t, 0, releaseList.Total)
	assert.True(t, releaseList.HasNextPage())
}

func TestReleases_Search(t *testing.T) {
	client := NewClient()

	releaseList, err := client.Releases("mix", "", 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.NotEqual(t, 0, releaseList.Total)
	assert.True(t, releaseList.HasNextPage())
}

func TestReleases_Search_NoResults(t *testing.T) {
	client := NewClient()

	releaseList, err := client.Releases("xxx not found xxx", "", 10, 0)
	assert.NoError(t, err)
	assert.Empty(t, releaseList.Results)
	assert.Equal(t, 0, releaseList.Total)
	assert.False(t, releaseList.HasNextPage())
}

func TestReleases_Type(t *testing.T) {
	client := NewClient()

	releaseList, err := client.Releases("", string(ReleaseTypeSingle), 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.NotEqual(t, 0, releaseList.Total)
	assert.True(t, releaseList.HasNextPage())
}

func TestReleases_Type_NoResults(t *testing.T) {
	client := NewClient()

	releaseList, err := client.Releases("", "xxx not found", 10, 0)
	assert.NoError(t, err)
	assert.Empty(t, releaseList.Results)
	assert.Equal(t, 0, releaseList.Total)
	assert.False(t, releaseList.HasNextPage())
}

func TestReleases_Search_Type(t *testing.T) {
	client := NewClient()

	releaseList, err := client.Releases("mix", string(ReleaseTypeCompilation), 5, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.NotEqual(t, 0, releaseList.Total)
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
