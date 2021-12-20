package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowseOptions_Default(t *testing.T) {
	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", Default().Encode())
}

func TestBrowseOptions_Complete(t *testing.T) {
	parameters := Default()
	WithSearch("mix")(&parameters)
	WithBrand(1)(&parameters)
	WithBrand(2)(&parameters)
	WithGenre("dubstep")(&parameters)
	WithGenre("acoustic")(&parameters)
	WithReleaseType("Single")(&parameters)
	WithReleaseType("EP")(&parameters)
	WithTag("chill")(&parameters)
	WithTag("badass")(&parameters)
	IncludeGold(false)(&parameters)
	IncludeUnreleased(false)(&parameters)
	WithSort("title")(&parameters)
	WithLimit(10)(&parameters)
	WithOffset(5)(&parameters)

	assert.Equal(t, "brands%5B%5D=1&brands%5B%5D=2&genres%5B%5D=dubstep&genres%5B%5D=acoustic&limit=10&nogold=true&offset=5&onlyReleased=true&search=mix&sort=title&tags%5B%5D=chill&tags%5B%5D=badass&types%5B%5D=Single&types%5B%5D=EP", parameters.Encode())
}

func TestBrowseOptions_WithSearch(t *testing.T) {
	parameters := Default()
	WithSearch("mix")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&search=mix&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithSearchMultiple(t *testing.T) {
	parameters := Default()
	WithSearch("mix")(&parameters)
	WithSearch("contest")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&search=contest&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithBrand(t *testing.T) {
	parameters := Default()
	WithBrand(1)(&parameters)

	assert.Equal(t, "brands%5B%5D=1&limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithBrand_Multiple(t *testing.T) {
	parameters := Default()
	WithBrand(1)(&parameters)
	WithBrand(2)(&parameters)

	assert.Equal(t, "brands%5B%5D=1&brands%5B%5D=2&limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithGenre(t *testing.T) {
	parameters := Default()
	WithGenre("dubstep")(&parameters)

	assert.Equal(t, "genres%5B%5D=dubstep&limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithGenre_Multiple(t *testing.T) {
	parameters := Default()
	WithGenre("dubstep")(&parameters)
	WithGenre("acoustic")(&parameters)

	assert.Equal(t, "genres%5B%5D=dubstep&genres%5B%5D=acoustic&limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithReleaseType(t *testing.T) {
	parameters := Default()
	WithReleaseType("Single")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date&types%5B%5D=Single", parameters.Encode())
}

func TestBrowseOptions_WithReleaseType_Multiple(t *testing.T) {
	parameters := Default()
	WithReleaseType("Single")(&parameters)
	WithReleaseType("EP")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date&types%5B%5D=Single&types%5B%5D=EP", parameters.Encode())
}

func TestBrowseOptions_WithTag(t *testing.T) {
	parameters := Default()
	WithTag("chill")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date&tags%5B%5D=chill", parameters.Encode())
}

func TestBrowseOptions_WithTag_Multiple(t *testing.T) {
	parameters := Default()
	WithTag("chill")(&parameters)
	WithTag("badass")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date&tags%5B%5D=chill&tags%5B%5D=badass", parameters.Encode())
}

func TestBrowseOptions_IncludeGold(t *testing.T) {
	parameters := Default()
	IncludeGold(false)(&parameters)

	assert.Equal(t, "limit=10&nogold=true&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_IncludeGold_Multiple(t *testing.T) {
	parameters := Default()
	IncludeGold(false)(&parameters)
	IncludeGold(true)(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_IncludeUnreleased(t *testing.T) {
	parameters := Default()
	IncludeUnreleased(false)(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=true&sort=-date", parameters.Encode())
}

func TestBrowseOptions_IncludeUnreleased_Multiple(t *testing.T) {
	parameters := Default()
	IncludeUnreleased(false)(&parameters)
	IncludeUnreleased(true)(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithSort(t *testing.T) {
	parameters := Default()
	WithSort("title")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=title", parameters.Encode())
}

func TestBrowseOptions_WithSort_Multiple(t *testing.T) {
	parameters := Default()
	WithSort("title")(&parameters)
	WithSort("-artist")(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=0&onlyReleased=false&sort=-artist", parameters.Encode())
}

func TestBrowseOptions_WithLimit(t *testing.T) {
	parameters := Default()
	WithLimit(5)(&parameters)

	assert.Equal(t, "limit=5&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithLimit_Multiple(t *testing.T) {
	parameters := Default()
	WithLimit(5)(&parameters)
	WithLimit(50)(&parameters)

	assert.Equal(t, "limit=50&nogold=false&offset=0&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithOffset(t *testing.T) {
	parameters := Default()
	WithOffset(5)(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=5&onlyReleased=false&sort=-date", parameters.Encode())
}

func TestBrowseOptions_WithOffset_Multiple(t *testing.T) {
	parameters := Default()
	WithOffset(5)(&parameters)
	WithOffset(50)(&parameters)

	assert.Equal(t, "limit=10&nogold=false&offset=50&onlyReleased=false&sort=-date", parameters.Encode())
}
