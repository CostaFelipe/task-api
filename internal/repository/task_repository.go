package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/CostaFelipe/task-api/internal/entity"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (t *TaskRepository) Create(ctx context.Context, task *entity.Task) error {
	query := `INSERT INTO tasks (title, description, priority, due_date, user_id)`

	result, err := t.DB.ExecContext(ctx, query, task.Title, task.Description, task.Priority, task.DueDate, task.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = int(id)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return nil
}
