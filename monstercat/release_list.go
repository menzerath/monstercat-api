package monstercat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const endpointReleaseList = "https://connect.monstercat.com/v2/releases"

// ReleaseList represents a list of releases from Monstercat API
type ReleaseList struct {
	Results []Release `json:"results"`
	Total   int       `json:"total"`
	Limit   int       `json:"limit"`
	Skip    int       `json:"skip"`
}

// GetReleaseList returns the first batch of entries in the current release list (max. limit of 50 entries)
func GetReleaseList() (ReleaseList, error) {
	return GetReleaseListAtPosition(50, 0)
}

// GetReleaseListAtPosition returns the requested batch of entries in the current release list
func GetReleaseListAtPosition(limit int, offset int) (ReleaseList, error) {
	releaseList := ReleaseList{}

	response, err := http.Get(fmt.Sprintf("%s?limit=%d&skip=%d", endpointReleaseList, limit, offset))
	if err != nil {
		return releaseList, err
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

// LoadNextPage loads the next batch of releases into the calling struct
func (releaseList *ReleaseList) LoadNextPage() error {
	newReleaseList, err := GetReleaseListAtPosition(releaseList.Limit, releaseList.Skip+releaseList.Limit)
	if err != nil {
		return err
	}

	releaseList.Results = newReleaseList.Results
	releaseList.Skip = newReleaseList.Skip
	return nil
}
