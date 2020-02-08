package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReleaseList(t *testing.T) {
	releaseList, err := GetReleaseList()
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.True(t, releaseList.HasNextPage())

	oldReleaseList := releaseList
	err = releaseList.LoadNextPage()
	assert.NoError(t, err)
	assert.NotEqual(t, oldReleaseList.Results, releaseList.Results)
}

func TestGetReleaseListAtPosition(t *testing.T) {
	releaseList, err := GetReleaseListAtPosition(5, 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, releaseList.Results)
	assert.True(t, releaseList.HasNextPage())

	oldReleaseList := releaseList
	err = releaseList.LoadNextPage()
	assert.NoError(t, err)
	assert.NotEqual(t, oldReleaseList.Results, releaseList.Results)
}

func TestGetReleaseListUntilEnd(t *testing.T) {
	releaseList, err := GetReleaseList()
	assert.NoError(t, err)

	for releaseList.HasNextPage() {
		assert.NoError(t, releaseList.LoadNextPage())
	}
}
