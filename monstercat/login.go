package monstercat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpointLogin = "https://connect.monstercat.com/v2/signin"

	authenticationCookieName = "connect.sid"
)

// ErrorInvalidCredentials is returned when a login fails because of invalid credentials supplied by the user
var ErrorInvalidCredentials = fmt.Errorf("invalid credentials")

// Login performs a login request into your Monstercat account and returns the created authentication cookie on success
func Login(email string, password string) (string, error) {
	payload, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, endpointLogin, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := getHTTPClient().Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return "", ErrorInvalidCredentials
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == authenticationCookieName {
			return cookie.Value, nil
		}
	}
	return "", fmt.Errorf("cookie not found")
}
