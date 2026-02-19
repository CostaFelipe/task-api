package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/CostaFelipe/task-api/internal/dto"
	"github.com/CostaFelipe/task-api/internal/entity"
	"github.com/CostaFelipe/task-api/internal/middleware"
	"github.com/CostaFelipe/task-api/internal/repository"
	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	taskRepo *repository.TaskRepository
}

func NewTaskHandler(taskRepo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		taskRepo: taskRepo,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {

	userId := middleware.GetUserIDFromContext(r.Context())

	var taskDto dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
		return
	}

	task, err := entity.NewTask(taskDto.Title, taskDto.Description, taskDto.Priority, userId)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "erro ao criar task"})
		return
	}

	if taskDto.DueDate == "" {
		dueDate, err := time.Parse("2006-01-02", taskDto.DueDate)
		if err != nil {
			responseJSON(w, http.StatusBadRequest, map[string]string{"error": "Formato inválido"})
			return
		}

		task.DueDate = &dueDate
	}

	if err = h.taskRepo.Create(r.Context(), task); err != nil {
		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao criar tarefa"})
		return
	}

	responseJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserIDFromContext(r.Context())

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "ID inválido"})
		return
	}

	task, err := h.taskRepo.FindByID(r.Context(), id, userId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "task não encontrada"})
			return
		}
		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "erro ao buscar task"})
		return
	}

	responseJSON(w, http.StatusOK, task)
}
