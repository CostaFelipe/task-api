package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CostaFelipe/task-api/internal/dto"
	"github.com/CostaFelipe/task-api/internal/repository"
)

type TaskHandler struct {
	taskRepo *repository.TaskRepository
}

func NewTaskHandler(taskRepo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		taskRepo: taskRepo,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
		return
	}
}
