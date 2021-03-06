package monstercat

import "fmt"

const (
	endpointLogin           = "https://connect.monstercat.com/v2/signin"
	endpointReleaseList     = "https://connect.monstercat.com/v2/releases"
	endpointReleaseDownload = "https://connect.monstercat.com/v2/release/%s/download?format=%s"

	authenticationCookieName = "cid"
)

var (
	// ErrorInvalidCredentials is returned when a login fails because of invalid credentials supplied by the user
	ErrorInvalidCredentials = fmt.Errorf("invalid credentials")

	// ErrorClientNotLoggedIn is returned when an action fails because of missing authentication
	ErrorClientNotLoggedIn = fmt.Errorf("client not logged in")
)

// Client is used to handle all operations, especially stateful ones like Login
type Client struct {
	authenticationCookie string
}

// NewClient creates and returns a new Client for interaction with the Monstercat API
func NewClient() Client {
	return Client{}
}

// IsLoggedIn returns true if the client is logged in, false otherwise
func (client Client) IsLoggedIn() bool {
	return client.authenticationCookie != ""
}
