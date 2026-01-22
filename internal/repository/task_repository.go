package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CostaFelipe/task-api/internal/dto"
	"github.com/CostaFelipe/task-api/internal/entity"
)

var (
	errTaskNotFound = errors.New("Task not found!")
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

func (t *TaskRepository) FindByID(ctx context.Context, id, userID int) (*entity.Task, error) {

	query := `SELECT id, title, description, completed, priority, due_date, user_id, created_at, updated_at FROM users WHERE id=? AND user_id=?`

	task := &entity.Task{}
	var dueDate sql.NullTime

	err := t.DB.QueryRowContext(ctx, query, id, userID).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		dueDate,
		&task.Priority,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errTaskNotFound
		}
	}

	if dueDate.Valid {
		task.DueDate = &dueDate.Time
	}

	return task, nil
}

func (t *TaskRepository) FindAllByUserID(ctx context.Context, userId int, filter *dto.TaskFilter) (*[]entity.Task, int, error) {
	queryBase := "FROM users WHERE user_id = ?"
	args := []interface{}{userId}

	if filter.Completed != nil {
		queryBase += " AND completed = ?"
		args = append(args, *filter.Completed)
	}

	if filter.Priority != nil {
		queryBase += " AND priority = ?"
		args = append(args, *filter.Priority)
	}

	var total int
	queryCount := "SELECT COUNT(*) " + queryBase
	if err := t.DB.QueryRowContext(ctx, queryCount, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf(`
					SELECT id, title, description, completed, priority, due_date,
								 user_id, created_at, update_at
					%s
					ORDER BY created_at DESC
					LIMIT ? OFFSET = ?`,
		queryBase)

	offset := (filter.Page - 1) * filter.Limit
	args = append(args, filter.Limit, offset)

	row, err := t.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer row.Close()

	var tasks []entity.Task

	for row.Next() {
		var task entity.Task
		var dueDate sql.NullTime
		if err := row.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
			&dueDate,
			&task.UserID,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, 0, nil
		}

		if dueDate.Valid {
			task.DueDate = &dueDate.Time
		}

		tasks = append(tasks, task)
	}

	return &tasks, total, row.Err()

}

func (t *TaskRepository) Update(ctx context.Context, task *entity.Task) error {
	var setParts []string
	var args []interface{}

	if task.Title != "" {
		setParts = append(setParts, "title = ?")
		args = append(args, task.Title)
	}

	if task.Description != "" {
		setParts = append(setParts, "description = ?")
		args = append(args, task.Description)
	}

	setParts = append(setParts, "completed = ?")
	args = append(args, task.Completed)

	if task.Priority != "" {
		setParts = append(setParts, "priority = ?")
		args = append(args, task.Priority)
	}

	if task.DueDate != nil {
		setParts = append(setParts, "due_date = ?")
		args = append(args, task.DueDate)
	}

	setParts = append(setParts, "updated_at = NOW()")

	query := fmt.Sprintf(`
	         UPDATE tasks
					 SET %s
					 WHERE id = ? AND user_id = ?
	`, strings.Join(setParts, ", "))

	args = append(args, task.ID, task.UserID)

	result, err := t.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
