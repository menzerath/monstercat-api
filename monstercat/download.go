package monstercat

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFormat describes in what kind of formats we can download a release
type DownloadFormat string

// define all known download formats
const (
	DownloadReleaseFormatFlac = "flac"
	DownloadReleaseFormatMP3  = "mp3_320"
	DownloadReleaseFormatWAV  = "wav"
)

// DownloadRelease downloads the given release as ZIP file in the requested format and stores it at the given path
func (client Client) DownloadRelease(release Release, downloadFormat DownloadFormat, downloadPath string) error {
	if !client.IsLoggedIn() {
		return ErrorClientNotLoggedIn
	}

	request, err := http.NewRequest("GET", fmt.Sprintf(endpointDownloadRelease, release.ID, downloadFormat), http.NoBody)
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

// TODO: add download function for a single catalog item
