package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsLoggedIn(t *testing.T) {
	client := NewClient()
	assert.False(t, client.IsLoggedIn())

	client.authenticationCookie = "xxx"
	assert.True(t, client.IsLoggedIn())
}
