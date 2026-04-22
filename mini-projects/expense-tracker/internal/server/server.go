package server

import (
	"log"
	"net/http"
	"time"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/expense-tracker/internal/handlers"
	"github.com/NutcrackerCom/go-backend-journey/mini-projects/expense-tracker/internal/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	Http   http.Server
}

func NewServer(logger *log.Logger) *Server {
	r := chi.NewRouter()

	expensesService := &service.ExpensesService{
		ExpensesList: make(map[service.ExpenseType][]service.Expenses),
	}
	handler := handlers.NewHandler(expensesService)

	r.Get("/expenses", handler.HandleGetExpenses)
	r.Post("/expenses", handler.HandlePostExpenses)
	r.Delete("/expenses", handler.HandleDeleteExpenses)

	return &Server{
		Logger: logger,
		Http: http.Server{
			Addr:         ":8080",
			Handler:      r,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
