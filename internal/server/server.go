package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final-tp1/internal/handlers"
)

// Server структура, содержащая логгер и http-сервер.
type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// NewServer создает новый HTTP-сервер с заданным логгером и зарегистрированными обработчиками.
func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	// Регистрация обработчиков
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создание экземпляра http.Server
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}

// Start запускает HTTP-сервер.
func (s *Server) Start() error {
	s.Logger.Println("Server is starting on port 8080")
	return s.HTTPServer.ListenAndServe()
}
