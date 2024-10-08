package server

import (
	"net/http"

	"github.com/ignatiusratemo26/go-backend/pkg/services/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func New() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/appointments", handlers.CreateAppointment).Methods("POST")
	s.router.HandleFunc("/appointments/{id}", handlers.GetAppointment).Methods("GET")
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.router)
}
