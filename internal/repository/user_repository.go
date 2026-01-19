package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CostaFelipe/task-api/internal/entity"
)

var ErrUserNotFound = errors.New("usuário não encontrado")
var ErrEmailExists = errors.New("email já cadastrado")

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepositoy(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?`

	user := &entity.User{}
	err := u.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
