package monstercat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Catalog represents a list of catalog items from Monstercat API
type Catalog struct {
	Data   []CatalogItem `json:"Data"`
	Total  int           `json:"Total"`
	Limit  int           `json:"Limit"`
	Offset int           `json:"Offset"`
}

// Catalog returns a set of catalog items containing the given search query, matching the given release type and being within the given range.
// While limit and offset are required, you may leave search and releaseType empty to ignore those filters.
// TODO: add brands[], tags[] and genres[]
func (client Client) Catalog(search string, releaseType string, limit int, offset int) (Catalog, error) {
	catalog := Catalog{}

	urlParameters := url.Values{}
	urlParameters.Add("search", search)
	urlParameters.Add("types[]", releaseType)
	urlParameters.Add("sort", "-date")         // TODO: make configurable + add values
	urlParameters.Add("nogold", "false")       // TODO: make configurable
	urlParameters.Add("onlyReleased", "false") // TODO: make configurable
	urlParameters.Add("limit", fmt.Sprintf("%d", limit))
	urlParameters.Add("offset", fmt.Sprintf("%d", offset))

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

// HasNextPage returns true if the catalog list contains more pages, false otherwise.
func (catalog Catalog) HasNextPage() bool {
	return (catalog.Offset + catalog.Limit) < catalog.Total
}