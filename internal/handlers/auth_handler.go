package handlers

import (
	"github.com/CostaFelipe/task-api/internal/middleware"
	"github.com/CostaFelipe/task-api/internal/repository"
)

type AuthHandler struct {
	userRepo       *repository.UserRepository
	authMiddleware *middleware.AuthMiddleware
}

func NewAuthHandler(userRepo *repository.UserRepository, authMiddleware *middleware.AuthMiddleware) *AuthHandler {
	return &AuthHandler{
		userRepo:       userRepo,
		authMiddleware: authMiddleware,
	}
}
