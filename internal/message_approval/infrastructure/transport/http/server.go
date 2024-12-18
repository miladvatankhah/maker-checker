package http

import (
	"fmt"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/presentation/http/v1"
	"net/http"
)

type Config struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type HTTPServer struct {
	cfg              Config
	messageHandlerV1 *v1.MessageHandler
	userHandlerV1    *v1.UserHandler
}

func NewHTTPServer(cfg Config, messageHandlerV1 *v1.MessageHandler, userHandlerV1 *v1.UserHandler) *HTTPServer {
	return &HTTPServer{cfg: cfg, messageHandlerV1: messageHandlerV1, userHandlerV1: userHandlerV1}
}

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()

	// API Version 1 routes
	apiV1 := http.NewServeMux()
	apiV1.HandleFunc("POST /messages", s.messageHandlerV1.CreateMessage)
	apiV1.HandleFunc("PATCH /messages/approve/{id}", s.messageHandlerV1.ApproveMessage)
	apiV1.HandleFunc("PATCH /messages/reject/{id}", s.messageHandlerV1.RejectMessage)
	apiV1.HandleFunc("POST /users", s.userHandlerV1.RegisterUser)

	// Group API V1 routes under /api/v1 prefix
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: mux,
	}

	return server.ListenAndServe()
}
