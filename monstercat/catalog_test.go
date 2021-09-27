package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatalog(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("", "", 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalog_Search(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("mix", "", 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalog_Search_NoResults(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("xxx not found xxx", "", 10, 0)
	assert.NoError(t, err)
	assert.Empty(t, catalog.Data)
	assert.Equal(t, 0, catalog.Total)
	assert.False(t, catalog.HasNextPage())
}

func TestCatalog_Type(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("", string(ReleaseTypeSingle), 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalog_Type_NoResults(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("", "xxx not found", 10, 0)
	assert.NoError(t, err)
	assert.Empty(t, catalog.Data)
	assert.Equal(t, 0, catalog.Total)
	assert.False(t, catalog.HasNextPage())
}

func TestCatalog_Search_Type(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("mix", string(ReleaseTypeCompilation), 5, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalogUntilEnd(t *testing.T) {
	client := NewClient()

	catalog, err := client.GetCatalog("", "", 100, 0)
	assert.NoError(t, err)

	for catalog.HasNextPage() {
		catalog, err = client.GetCatalog("", "", catalog.Limit, catalog.Offset+catalog.Limit)
		assert.NoError(t, err)
	}
}
