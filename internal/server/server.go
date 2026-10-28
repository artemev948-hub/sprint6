package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.Upload)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}
}

func (s *Server) Run() error {
	return s.Server.ListenAndServe()
}
