package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatalog(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("", "", 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalog_Search(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("mix", "", 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalog_Search_NoResults(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("xxx not found xxx", "", 10, 0)
	assert.NoError(t, err)
	assert.Empty(t, catalog.Data)
	assert.Equal(t, 0, catalog.Total)
	assert.False(t, catalog.HasNextPage())
}

func TestCatalog_Type(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("", string(ReleaseTypeSingle), 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalog_Type_NoResults(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("", "xxx not found", 10, 0)
	assert.NoError(t, err)
	assert.Empty(t, catalog.Data)
	assert.Equal(t, 0, catalog.Total)
	assert.False(t, catalog.HasNextPage())
}

func TestCatalog_Search_Type(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("mix", string(ReleaseTypeCompilation), 5, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, catalog.Data)
	assert.NotEqual(t, 0, catalog.Total)
	assert.True(t, catalog.HasNextPage())
}

func TestCatalogUntilEnd(t *testing.T) {
	client := NewClient()

	catalog, err := client.Catalog("", "", 100, 0)
	assert.NoError(t, err)

	for catalog.HasNextPage() {
		catalog, err = client.Catalog("", "", catalog.Limit, catalog.Offset+catalog.Limit)
		assert.NoError(t, err)
	}
}
