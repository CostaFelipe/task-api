package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CostaFelipe/task-api/internal/dto"
	"github.com/CostaFelipe/task-api/internal/entity"
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

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
		return
	}

	if err = h.userRepo.Create(r.Context(), user); err != nil {
		if errors.Is(err, repository.ErrEmailExists) {
			responseJSON(w, http.StatusBadRequest, map[string]string{"error": "Email já cadastrado"})
			return
		}
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "erro ao criar o usuário"})
		return
	}

	token, err := h.authMiddleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "Erro ao gerar token"})
		return
	}

	response := dto.AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}

	responseJSON(w, http.StatusCreated, response)
}

func responseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
