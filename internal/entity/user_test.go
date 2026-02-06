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
	assert.NotEmpty(t, user.Password)
	assert.NotZero(t, user.CreatedAt)
	assert.NotZero(t, user.UpdatedAt)
}

func TestUserValidation(t *testing.T) {
	t.Run("Empty Name", func(t *testing.T) {
		user, err := NewUser("", "jhonny@doe.com", "123456")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, errNameEmpty, err)
	})

	t.Run("Empty Email", func(t *testing.T) {
		user, err := NewUser("Joe", "", "123456")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, errEmailEmpty, err)
	})

	t.Run("Empty Password", func(t *testing.T) {
		user, err := NewUser("JTest", "teste@doe.com", "")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, errPasswordEmpty, err)
	})
}
