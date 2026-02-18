package entity

import (
	"errors"
	"time"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

var (
	errTitleIsRequired       = errors.New("title is required")
	errDescriptionIsRequired = errors.New("description is required")
	//errIDIsRequired          = errors.New("id is required")
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Priority    Priority   `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	UserID      int        `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(title, description string, priority Priority, userId int) (*Task, error) {
	task := &Task{
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    getValidPriority(priority),
		UserID:      userId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := task.Validation(); err != nil {
		return nil, err
	}

	return task, nil
}

func (t *Task) Validation() error {
	if t.Title == "" {
		return errTitleIsRequired
	}

	if t.Description == "" {
		return errDescriptionIsRequired
	}
	return nil
}

func getValidPriority(priority Priority) Priority {
	if priority == "" {
		priority = PriorityMedium
	}

	return priority
}
