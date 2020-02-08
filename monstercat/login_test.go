package monstercat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	cookie, err := Login("fake@user.com", "test123")
	assert.Error(t, err)
	assert.Equal(t, ErrorInvalidCredentials, err)
	assert.Empty(t, cookie)
}
