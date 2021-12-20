package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowseCatalog_DefaultOptions(t *testing.T) {
	client := NewClient()

	catalog, err := client.BrowseCatalog()
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.Len(t, catalog.Data, 10)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestBrowseCatalog_WithOption(t *testing.T) {
	client := NewClient()

	catalog, err := client.BrowseCatalog(WithSearch("mix"))
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.Len(t, catalog.Data, 10)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestBrowseCatalog_WithOption_NoResults(t *testing.T) {
	client := NewClient()

	catalog, err := client.BrowseCatalog(WithSearch("xxx not found xxx"))
	assert.NoError(t, err)
	assert.Empty(t, catalog.Data)
	assert.Equal(t, 0, catalog.Total)
	assert.False(t, catalog.HasNextPage())
}

func TestBrowseCatalog_WithOptions(t *testing.T) {
	client := NewClient()

	catalog, err := client.BrowseCatalog(WithSearch("mix"), WithReleaseType(string(ReleaseTypeCompilation)))
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
}

func TestBrowseCatalog_WithLimitAndOffset(t *testing.T) {
	client := NewClient()

	catalog, err := client.BrowseCatalog(WithLimit(5), WithOffset(3))
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.Len(t, catalog.Data, 5)
	assert.NotEqual(t, 0, catalog.Total)
	assert.Equal(t, 5, catalog.Limit)
	assert.Equal(t, 3, catalog.Offset)
}

func TestBrowseCatalogUntilEnd(t *testing.T) {
	client := NewClient()

	catalog, err := client.BrowseCatalog(WithLimit(100))
	assert.NoError(t, err)

	for catalog.HasNextPage() {
		catalog, err = client.BrowseCatalog(WithLimit(catalog.Limit), WithOffset(catalog.Offset+catalog.Limit))
		assert.NoError(t, err)
	}
}

func TestBrowseCatalogEqualsCatalog(t *testing.T) {
	const (
		search      = "mix"
		releaseType = string(ReleaseTypePodcast)
		limit       = 3
		offset      = 5
	)
	client := NewClient()

	browseCatalog, err := client.BrowseCatalog(WithSearch(search), WithReleaseType(releaseType), WithLimit(limit), WithOffset(offset))
	assert.NoError(t, err)
	assert.NotEmpty(t, browseCatalog.Data)

	catalog, err := client.Catalog(search, releaseType, limit, offset)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)

	assert.Equal(t, browseCatalog.Data, catalog.Data)
}
