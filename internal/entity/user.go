package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	errNameEmpty     = errors.New("O nome não pode ser vazio")
	errEmailEmpty    = errors.New("O email não pode ser vazio")
	errPasswordEmpty = errors.New("O password não pode ser vazio")
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hash)
	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errNameEmpty
	}

	if u.Email == "" {
		return errEmailEmpty
	}

	if u.Password == "" {
		return errPasswordEmpty
	}

	return nil
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Created: u.CreatedAt,
	}
}
