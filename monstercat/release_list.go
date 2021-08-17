package monstercat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ReleaseList represents a list of releases from Monstercat API
type ReleaseList struct {
	Results []Release `json:"results"`
	Total   int       `json:"total"`
	Limit   int       `json:"limit"`
	Skip    int       `json:"skip"`
}

// ReleaseList returns the first batch of entries in the current release list (max. limit of 50 entries).
// Depreacted: use Releases instead.
func (client Client) ReleaseList() (ReleaseList, error) {
	return client.ReleaseListAtPosition(50, 0)
}

// ReleaseListAtPosition returns the requested batch of entries in the current release list.
// Depreacted: use Releases instead.
func (client Client) ReleaseListAtPosition(limit int, offset int) (ReleaseList, error) {
	return client.Releases("", "", limit, offset)
}

// Releases returns a set of releases containing the given search query, matching the given release type and being within the given range.
// While limit and offset are required, you may leave search and releaseType empty to ignore those filters.
func (client Client) Releases(search string, releaseType string, limit int, offset int) (ReleaseList, error) {
	releaseList := ReleaseList{}

	urlParameters := url.Values{}
	urlParameters.Add("search", search)
	urlParameters.Add("type", releaseType)
	urlParameters.Add("limit", fmt.Sprintf("%d", limit))
	urlParameters.Add("offset", fmt.Sprintf("%d", offset))

	request, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", endpointReleaseList, urlParameters.Encode()), http.NoBody)
	if err != nil {
		return releaseList, err
	}
	if client.IsLoggedIn() {
		request.Header.Set("Cookie", fmt.Sprintf("%s=%s", authenticationCookieName, client.authenticationCookie))
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return releaseList, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)
		if err != nil {
			return releaseList, fmt.Errorf("http error %d", response.StatusCode)
		}
		return releaseList, fmt.Errorf("http error %d: %s", response.StatusCode, message)
	}

	err = json.NewDecoder(response.Body).Decode(&releaseList)
	if err != nil {
		return releaseList, err
	}

	return releaseList, nil
}

// HasNextPage returns true if the release list contains more pages, false otherwise.
func (releaseList ReleaseList) HasNextPage() bool {
	return (releaseList.Skip + releaseList.Limit) < releaseList.Total
}
