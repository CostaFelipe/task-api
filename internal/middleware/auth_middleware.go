package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/CostaFelipe/task-api/config"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type JWTClaims struct {
	UserID int    `json:"user_id"`
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

func (h *AuthMiddleware) GenerateToken(userId int, email string) (string, error) {
	claims := &JWTClaims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(h.jwtConfig.JWTExpiresIn) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.jwtConfig.JWTSecret))
}

func (h *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("authentication")
		if authHeader == "" {
			http.Error(w, `{"error": "dados inválidos"}`, http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"error": "formato do token inválido"}`, http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.jwtConfig.JWTSecret), nil
		})

		if err != nil {
			http.Error(w, `{"error": "token inválido"}`, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok || !token.Valid {
			http.Error(w, `{"error": "token inválido"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
