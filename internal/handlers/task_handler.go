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

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserIDFromContext(r.Context())

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	filter := dto.TaskFilter{
		Page:  page,
		Limit: limit,
	}

	if completStr := r.URL.Query().Get("completed"); completStr != "" {
		completed := completStr == "true"
		filter.Completed = &completed
	}

	if priorityStr := r.URL.Query().Get("priority"); priorityStr != "" {
		priorityStr := entity.Priority(priorityStr)
		filter.Priority = &priorityStr
	}

	tasks, total, err := h.taskRepo.FindAllByUserID(r.Context(), userId, &filter)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, map[string]string{"err": "erro ao buscar tasks"})
		return
	}

	response := map[string]interface{}{
		"data":        tasks,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": (total + limit - 1) / limit,
	}

	responseJSON(w, http.StatusOK, response)

}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserIDFromContext(r.Context())
	taskId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "id inválido"})
		return
	}

	task, err := h.taskRepo.FindByID(r.Context(), taskId, userId)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			responseJSON(w, http.StatusNotFound, map[string]string{"error": "task não encontrada"})
			return
		}

		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "erro ao buscar tarefa"})
		return
	}

	var req dto.UpdateTask
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
	}
	if req.Title != nil {
		task.Title = *req.Title
	}

	if req.Description != nil {
		task.Description = *req.Description
	}

	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	if req.Priority != nil {
		task.Priority = *req.Priority
	}

	if req.DueDate != nil {
		if *req.DueDate != "" {
			task.DueDate = nil
		} else {
			dueDate, err := time.Parse("2006-01-02", *req.DueDate)
			if err != nil {
				responseJSON(w, http.StatusBadRequest, map[string]string{"error": "formato due date inválido"})
				return
			}

			task.DueDate = &dueDate
		}
	}

	if err := h.taskRepo.Update(r.Context(), task); err != nil {
		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar tasks"})
		return
	}

	taskUpdate, _ := h.taskRepo.FindByID(r.Context(), taskId, userId)
	responseJSON(w, http.StatusOK, taskUpdate)

}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserIDFromContext(r.Context())
	taskId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "id inválido"})
		return
	}

	if err := h.taskRepo.Delete(r.Context(), userId, taskId); err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "tarefa não encontrada"})
			return
		}
		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "erro ao deletar tarefa"})
		return
	}
}

func (h *TaskHandler) ToggleComplete(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetUserIDFromContext(r.Context())
	taskId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "id inválda"})
		return
	}

	task, err := h.taskRepo.ToggleComplete(r.Context(), taskId, userId)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			responseJSON(w, http.StatusNotFound, map[string]string{"error": "task não encontrada"})
			return
		}

		responseJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar tarefa"})
		return
	}

	responseJSON(w, http.StatusOK, task)
}
