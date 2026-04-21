package server

import (
	"log"
	"net/http"
	"time"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/todo-list-server/internal/handlers"
	"github.com/NutcrackerCom/go-backend-journey/mini-projects/todo-list-server/internal/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	Http   http.Server
}

func NewServer(logger *log.Logger) Server {
	r := chi.NewRouter()

	taskService := &service.TaskService{}
	handler := handlers.NewHandler(taskService)

	r.Get("/", handler.HandleGetMain)
	r.Post("/", handler.HandlePostMain)
	r.Delete("/", handler.HandleDeleteMain)

	return Server{
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
