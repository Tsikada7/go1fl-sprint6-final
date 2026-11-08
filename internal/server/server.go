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

func NewServer(l *log.Logger) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.PageHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	server := Server{
		Logger: l,
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     l,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
	return &server

}
