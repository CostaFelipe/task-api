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
	errIDIsRequired          = errors.New("id is required")
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
	return &Task{
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
		UserID:      userId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
