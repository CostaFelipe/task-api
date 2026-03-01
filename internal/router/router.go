package router

import (
	"net/http"

	"github.com/CostaFelipe/task-api/internal/handlers"
	"github.com/CostaFelipe/task-api/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(authHandler *handlers.AuthHandler, taskHandler *handlers.TaskHandler, authMiddleware *middleware.AuthMiddleware) http.Handler {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)

	return r
}
