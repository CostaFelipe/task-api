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
	if err := Validate(name, email, password); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:      name,
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user, nil
}

func Validate(name, email, password string) error {
	if name == "" {
		return errNameEmpty
	}

	if email == "" {
		return errEmailEmpty
	}

	if password == "" {
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

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
