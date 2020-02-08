package monstercat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Login performs a login request into your Monstercat account and stores the created authentication cookie on success for further requests
func (client *Client) Login(email string, password string) error {
	payload, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return err
	}

	response, err := http.Post(endpointLogin, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return ErrorInvalidCredentials
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == authenticationCookieName {
			client.authenticationCookie = cookie.Value
			return nil
		}
	}
	return fmt.Errorf("cookie not found")
}
