package server

import (
	"context"
	"net/http"
	"time"

	"github.com/kokhno-nikolay/lets-go-chat/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.ServerPort,
			Handler:        handler,
			ReadTimeout:    time.Duration(cfg.ServerReadTimeout),
			WriteTimeout:   time.Duration(cfg.ServerWriteTimeout),
			MaxHeaderBytes: cfg.ServerMaxHeaderBytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
