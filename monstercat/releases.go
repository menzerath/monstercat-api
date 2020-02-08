package monstercat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const endpointReleaseDownload = "https://connect.monstercat.com/v2/release/%s/download?format=%s"

// ReleaseType describes what kind of release we are looking at
type ReleaseType string

// define all known release types
const (
	ReleaseTypeSingle  ReleaseType = "Single"
	ReleaseTypeAlbum   ReleaseType = "Album"
	ReleaseTypePodcast ReleaseType = "Podcast"
	ReleaseTypeEP      ReleaseType = "EP"
)

// Release represents a single release from Monstercat API
type Release struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Artist      string      `json:"artistsTitle"`
	CatalogID   string      `json:"catalogId"`
	ReleaseDate time.Time   `json:"releaseDate"`
	Type        ReleaseType `json:"type"`

	GenrePrimary   string `json:"genrePrimary"`
	GenreSecondary string `json:"genreSecondary"`

	Downloadable bool `json:"downloadable"`
	Streamable   bool `json:"streamable"`
}

// DownloadFormat describes in what kind of formats we can download a release
type DownloadFormat string

// define all known download formats
const (
	ReleaseDownloadFormatFlac = "flac"
	ReleaseDownloadFormatMP3  = "mp3_320"
	ReleaseDownloadFormatWAV  = "wav"
)

// Download downloads the current release as ZIP file in the given format and returns the retrieved file
func (release Release) Download(authenticationCookie string, downloadFormat DownloadFormat) ([]byte, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(endpointReleaseDownload, release.ID, downloadFormat), http.NoBody)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Cookie", fmt.Sprintf("%s=%s", authenticationCookieName, authenticationCookie))

	response, err := getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// check for correct status code
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("http status %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
