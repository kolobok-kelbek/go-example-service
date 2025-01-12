package app

import (
	"context"
	"fmt"
	"github.com/kolobok-kelbek/tomato/internal/config"
	"github.com/kolobok-kelbek/tomato/internal/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	config *config.Config
	server *Server
	logger logger.Logger
}

func NewApp(config *config.Config, server *Server, logger logger.Logger) *App {
	return &App{
		config: config,
		server: server,
		logger: logger,
	}
}

func (a *App) Run() error {
	a.logger.Info("Start application")
	defer a.logger.Info("Application stopped")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := a.server.Up(); err != nil {
			fmt.Printf("HTTP server Start error: %v\n", err)
			stop <- syscall.SIGTERM
		}
	}()

	<-stop
	fmt.Println("Shutting down application...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := a.server.Down(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	return nil
}
