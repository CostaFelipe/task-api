package dto

import "github.com/CostaFelipe/task-api/internal/entity"

type TaskFilter struct {
	Completed *bool
	Priority  *entity.Priority
	Page      int
	Limit     int
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string              `json:"token"`
	User  entity.UserResponse `json:"user"`
}
