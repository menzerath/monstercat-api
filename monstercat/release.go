package monstercat

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// ReleaseType describes what kind of release we are looking at
type ReleaseType string

// define all known release types
const (
	ReleaseTypeAlbum       ReleaseType = "Album"
	ReleaseTypeCompilation ReleaseType = "Compilation"
	ReleaseTypeEP          ReleaseType = "EP"
	ReleaseTypeMixes       ReleaseType = "Mixes"
	ReleaseTypePodcast     ReleaseType = "Podcast"
	ReleaseTypeSingle      ReleaseType = "Single"
)

// CatalogItem represents a single release from Monstercat API
type CatalogItem struct {
	ID           string    `json:"Id"`
	Title        string    `json:"Title"`
	ArtistsTitle string    `json:"ArtistsTitle"`
	DebutDate    time.Time `json:"DebutDate"`
	Release      Release   `json:"Release"`

	GenrePrimary   string `json:"GenrePrimary"`
	GenreSecondary string `json:"GenreSecondary"`

	Downloadable  bool `json:"Downloadable"`
	InEarlyAccess bool `json:"InEarlyAccess"`
	Streamable    bool `json:"Streamable"`
}

type Release struct {
	CatalogID   string      `json:"CatalogId"`
	ReleaseDate time.Time   `json:"ReleaseDate"`
	Type        ReleaseType `json:"Type"`
}

// DownloadFormat describes in what kind of formats we can download a release
type DownloadFormat string

// define all known download formats
const (
	ReleaseDownloadFormatFlac = "flac"
	ReleaseDownloadFormatMP3  = "mp3_320"
	ReleaseDownloadFormatWAV  = "wav"
)

// DownloadRelease downloads the given release as ZIP file in the requested format and stores it at the given path
func (client Client) DownloadRelease(catalogItem CatalogItem, downloadFormat DownloadFormat, downloadPath string) error {
	if !client.IsLoggedIn() {
		return ErrorClientNotLoggedIn
	}

	request, err := http.NewRequest("GET", fmt.Sprintf(endpointReleaseDownload, catalogItem.ID, downloadFormat), http.NoBody)
	if err != nil {
		return err
	}
	if client.IsLoggedIn() {
		request.Header.Set("Cookie", fmt.Sprintf("%s=%s", authenticationCookieName, client.authenticationCookie))
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("http error %d", response.StatusCode)
		}
		return fmt.Errorf("http error %d: %s", response.StatusCode, message)
	}

	// create and save file
	targetFile, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(targetFile, response.Body)
	if err != nil {
		return err
	}

	return nil
}
