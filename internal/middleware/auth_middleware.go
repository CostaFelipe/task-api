package middleware

import (
	"github.com/CostaFelipe/task-api/config"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type UserClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	jwtConfig config.Config
}

func NewAuthMiddleware(cfg config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		jwtConfig: cfg,
	}
}
