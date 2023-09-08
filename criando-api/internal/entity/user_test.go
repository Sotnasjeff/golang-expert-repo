package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jeff", "Jeff@gmail.com", "1526874")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jeff", user.Name)
	assert.Equal(t, "Jeff@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jeff", "Jeff@gmail.com", "1526874")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1526874"))
	assert.False(t, user.ValidatePassword("15870874"))
	assert.NotEqual(t, "1526874", user.Password)
}
