package monstercat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ReleaseList represents a list of releases from Monstercat API
type ReleaseList struct {
	Results []Release `json:"results"`
	Total   int       `json:"total"`
	Limit   int       `json:"limit"`
	Skip    int       `json:"skip"`
}

// ReleaseList returns the first batch of entries in the current release list (max. limit of 50 entries)
func (client Client) ReleaseList() (ReleaseList, error) {
	return client.ReleaseListAtPosition(50, 0)
}

// ReleaseListAtPosition returns the requested batch of entries in the current release list
func (client Client) ReleaseListAtPosition(limit int, offset int) (ReleaseList, error) {
	releaseList := ReleaseList{}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s?limit=%d&skip=%d", endpointReleaseList, limit, offset), http.NoBody)
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

// HasNextPage returns true if the release list contains more pages, false otherwise
func (releaseList ReleaseList) HasNextPage() bool {
	return (releaseList.Skip + releaseList.Limit) < releaseList.Total
}
