package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/expense-tracker/internal/service"
)

type Handler struct {
	ExpensesService *service.ExpensesService
}

func NewHandler(expensesService *service.ExpensesService) *Handler {
	return &Handler{
		ExpensesService: expensesService,
	}
}

func (h *Handler) HandleGetExpenses(w http.ResponseWriter, r *http.Request) {
	jsonExpenses, err := json.MarshalIndent(h.ExpensesService.List(), "", "	")
	if err != nil {
		http.Error(w, "ошибка при сериализации", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonExpenses)
}

func (h *Handler) HandlePostExpenses(w http.ResponseWriter, r *http.Request) {
	var expenceses service.Expenses

	if err := json.NewDecoder(r.Body).Decode(&expenceses); err != nil {
		http.Error(w, "ошибка при десериализации", http.StatusInternalServerError)
		return
	}

	if err := h.ExpensesService.Add(expenceses); err != nil {
		if errors.Is(err, service.ErrorEmptyTitle) ||
			errors.Is(err, service.ErrorIncorrectCategory) ||
			errors.Is(err, service.ErrorIncorrectAmount) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) HandleDeleteExpenses(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.ExpensesService.DeleteById(id); err != nil {
		if errors.Is(err, service.ErrorIdNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
