package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	user, err := NewUser("John Doe", "jhonny@doe.com", "123456")

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "jhonny@doe.com", user.Email)
	assert.NotZero(t, user.CreatedAt)
	assert.NotZero(t, user.UpdatedAt)
}

func TestEmptyNameUser(t *testing.T) {
	user, err := NewUser("", "jhonny@doe.com", "123456")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, errNameEmpty, err)
}
