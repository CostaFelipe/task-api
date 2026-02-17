package dto

import "github.com/CostaFelipe/task-api/internal/entity"

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string              `json:"token"`
	User  entity.UserResponse `json:"user"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TaskFilter struct {
	Completed *bool
	Priority  *entity.Priority
	Page      int
	Limit     int
}

type CreateTaskRequest struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Priority    entity.Priority `json:"priority"`
	DueDate     string          `json:"due_date,omitempty"`
}
