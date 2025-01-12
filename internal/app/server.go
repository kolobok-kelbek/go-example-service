package app

import (
	"context"
	"fmt"
	"github.com/kolobok-kelbek/tomato/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Server, mux *http.ServeMux) *Server {
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}

	return &Server{httpServer: srv}
}

func (s *Server) Up() error {
	fmt.Printf("Starting HTTP server on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Down(ctx context.Context) error {
	fmt.Println("HTTP server is shutting down...")
	return s.httpServer.Shutdown(ctx)
}
