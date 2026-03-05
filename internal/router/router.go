package router

import (
	"net/http"
	"time"

	"github.com/CostaFelipe/task-api/internal/handlers"
	"github.com/CostaFelipe/task-api/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(authHandler *handlers.AuthHandler, taskHandler *handlers.TaskHandler, authMiddleware *middleware.AuthMiddleware) http.Handler {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Timeout(60 * time.Minute))
	r.Use(corsMiddleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "OK", "message": "API is running"}`))
	})

	return r
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow", "Accept, Authorization, Content-Type, x-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
