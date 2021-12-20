package monstercat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Catalog represents a list of catalog items from Monstercat API
type Catalog struct {
	Data   []CatalogItem `json:"Data"`
	Total  int           `json:"Total"`
	Limit  int           `json:"Limit"`
	Offset int           `json:"Offset"`
}

// BrowseCatalog returns a set of catalog items defined by the given BrowseOptions.
// All BrowseOptions will override their default values defined in Default().
func (client Client) BrowseCatalog(options ...BrowseOption) (Catalog, error) {
	catalog := Catalog{}

	urlParameters := Default()
	for _, option := range options {
		option(&urlParameters)
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", endpointCatalog, urlParameters.Encode()), http.NoBody)
	if err != nil {
		return catalog, err
	}
	if client.IsLoggedIn() {
		request.Header.Set("Cookie", fmt.Sprintf("%s=%s", authenticationCookieName, client.authenticationCookie))
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return catalog, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)
		if err != nil {
			return catalog, fmt.Errorf("http error %d", response.StatusCode)
		}
		return catalog, fmt.Errorf("http error %d: %s", response.StatusCode, message)
	}

	err = json.NewDecoder(response.Body).Decode(&catalog)
	if err != nil {
		return catalog, err
	}

	return catalog, nil
}

// Catalog returns a set of catalog items containing the given search query, matching the given release type and being within the given range.
// While limit and offset are required, you may leave search and releaseType empty to ignore those filters.
// Deprecated: Use the BrowseCatalog function instead.
func (client Client) Catalog(search string, releaseType string, limit int, offset int) (Catalog, error) {
	return client.BrowseCatalog(WithSearch(search), WithReleaseType(releaseType), WithLimit(limit), WithOffset(offset))
}

// HasNextPage returns true if the catalog list contains more pages, false otherwise.
func (catalog Catalog) HasNextPage() bool {
	return (catalog.Offset + catalog.Limit) < catalog.Total
}
