package http

import (
	"net/http"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account/internal/bank"
)

type Server struct {
	svc *bank.Service
}

func NewServer(svc *bank.Service) *Server {
	newServer := Server{svc: svc}
	return &newServer
}

func (s *Server) Handler() http.Handler {

}
