package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CostaFelipe/task-api/internal/dto"
	"github.com/CostaFelipe/task-api/internal/entity"
	"github.com/CostaFelipe/task-api/internal/middleware"
	"github.com/CostaFelipe/task-api/internal/repository"
	"github.com/CostaFelipe/task-api/internal/util"
	"github.com/CostaFelipe/task-api/pkg/responses"
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
		util.ResponseJSON(w, http.StatusBadRequest, responses.ErrorResponse{Error: "dados inválidos"})
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		util.ResponseJSON(w, http.StatusBadRequest, responses.ErrorResponse{Error: "dados inválidos"})
		return
	}

	if err = h.userRepo.Create(r.Context(), user); err != nil {
		if errors.Is(err, repository.ErrEmailExists) {
			util.ResponseJSON(w, http.StatusBadRequest, responses.ErrorResponse{Error: "Email já cadastrado"})
			return
		}
		util.ResponseJSON(w, http.StatusBadRequest, responses.ErrorResponse{Error: "erro ao criar usuário"})
		return
	}

	token, err := h.authMiddleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		util.ResponseJSON(w, http.StatusBadRequest, responses.ErrorResponse{Error: "erro ao gerar token"})
		return
	}

	response := dto.AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}

	util.ResponseJSON(w, http.StatusCreated, response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		util.ResponseJSON(w, http.StatusBadRequest, responses.ErrorResponse{Error: "dados inválidos"})
		return
	}

	user, err := h.userRepo.FindByEmail(r.Context(), userDto.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			util.ResponseJSON(w, http.StatusNotFound, responses.ErrorResponse{Error: "usuário não encontrado"})
			return
		}
		util.ResponseJSON(w, http.StatusUnauthorized, responses.ErrorResponse{Error: "erro ao encontrar usuário"})
		return
	}

	err = user.ValidatePassword(userDto.Password)
	if err != nil {
		util.ResponseJSON(w, http.StatusUnauthorized, responses.ErrorResponse{Error: "senha errada"})
		return
	}

	token, err := h.authMiddleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		util.ResponseJSON(w, http.StatusUnauthorized, responses.ErrorResponse{Error: "erro ao gerar token"})
		return
	}

	response := dto.AuthResponse{
		Token: token,
		User:  user.ToResponse(),
	}

	util.ResponseJSON(w, http.StatusOK, response)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetUserIDFromContext(r.Context())

	user, err := h.userRepo.FindByID(r.Context(), id)
	if err != nil {
		util.ResponseJSON(w, http.StatusNotFound, responses.ErrorResponse{Error: "usuário não encontrado"})
		return
	}

	util.ResponseJSON(w, http.StatusOK, user.ToResponse())
}
