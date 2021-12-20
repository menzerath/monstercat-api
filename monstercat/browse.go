package monstercat

import (
	"fmt"
	"net/url"
)

// BrowseOption is a function that can be used to modify the query parameters of the catalog endpoint.
type BrowseOption func(parameters *url.Values)

// Default returns the default query parameters for the catalog endpoint.
func Default() url.Values {
	defaultOptions := []BrowseOption{
		IncludeGold(true),
		IncludeUnreleased(true),
		WithSort("-date"),
		WithLimit(10),
		WithOffset(0),
	}

	parameters := url.Values{}
	for _, option := range defaultOptions {
		option(&parameters)
	}
	return parameters
}

// WithSearch limits the results to the tracks, albums or artists matching the given search query.
func WithSearch(query string) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Set("search", query)
	}
}

// WithBrand limits the results to the tracks released under the given brands.
// 1: Uncaged
// 2: Instinct
// 3: Call of the Wild
// 4: Silk
// 5: Silk Showcase
func WithBrand(brandID int) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Add("brands[]", fmt.Sprintf("%d", brandID))
	}
}

// WithGenre limits the results to the tracks matching the given genre.
// Examples are "dubstep", "acoustic" or "melodic house & techno". See the Monstercat Player for a full list of genres.
func WithGenre(genre string) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Add("genres[]", genre)
	}
}

// WithReleaseType limits the results to the tracks matching the given release type.
// Either "Single", "EP" or "Album".
func WithReleaseType(releaseType string) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Add("types[]", releaseType)
	}
}

// WithTag limits the results to the tracks matching the given tag.
// Examples are "chill", "badass" or "funky". See the Monstercat Player for a full list of genres.
func WithTag(tag string) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Add("tags[]", tag)
	}
}

// IncludeGold decides whether to include gold tracks or not.
func IncludeGold(include bool) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Set("nogold", fmt.Sprintf("%t", !include))
	}
}

// IncludeUnreleased decides whether to include unreleased tracks or not.
func IncludeUnreleased(include bool) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Set("onlyReleased", fmt.Sprintf("%t", !include))
	}
}

// WithSort sorts the results by the given field.
// Adding a minus sign to the field name will sort in descending order.
// Either "title", "artists", "release", "genre", "date", "bpm", "duration" or "brand".
func WithSort(sortBy string) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Set("sort", sortBy)
	}
}

// WithLimit limits the results to the given number of items.
func WithLimit(limit int) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Set("limit", fmt.Sprintf("%d", limit))
	}
}

// WithOffset allows to skip the given number of items.
func WithOffset(offset int) BrowseOption {
	return func(parameters *url.Values) {
		parameters.Set("offset", fmt.Sprintf("%d", offset))
	}
}
