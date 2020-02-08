package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	client := NewClient()

	err := client.Login("fake@user.com", "test123")
	assert.Error(t, err)
	assert.Equal(t, ErrorInvalidCredentials, err)

	assert.Empty(t, client.authenticationCookie)
	assert.False(t, client.IsLoggedIn())
}
