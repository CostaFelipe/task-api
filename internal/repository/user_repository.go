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

func (u *UserRepository) Create(ctx context.Context, user *entity.User) error {
	var exists bool
	err := u.DB.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", user.Email).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return ErrEmailExists
	}

	query := `INSERT INTO users (name, email, password) VALUES (?,?,?)`

	result, err := u.DB.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
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

func (u *UserRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?`

	user := &entity.User{}

	err := u.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)

	return user, nil
}
