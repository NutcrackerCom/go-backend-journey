package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/todo-list-server/internal/service"
)

type Handler struct {
	TaskService *service.TaskService
}

func NewHandler(taskService *service.TaskService) *Handler {
	return &Handler{
		TaskService: taskService,
	}
}

func (h *Handler) HandleGetMain(w http.ResponseWriter, r *http.Request) {
	jsonList, err := json.MarshalIndent(h.TaskService.List(), "", "	")
	if err != nil {
		http.Error(w, "ошибка при сериализации", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonList)
}

func (h *Handler) HandlePostMain(w http.ResponseWriter, r *http.Request) {
	var task service.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.TaskService.Add(task)
	if err != nil {
		if errors.Is(err, service.ErrorIncorrectStatus) ||
			errors.Is(err, service.ErrorEmptyTaskName) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) HandleDeleteMain(w http.ResponseWriter, r *http.Request) {
	text, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(string(text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.TaskService.DeleteById(id)
	if err != nil {
		if errors.Is(err, service.ErrorIdNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
