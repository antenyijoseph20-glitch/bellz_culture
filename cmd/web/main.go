package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antenyijoseph20-glitch/bellz_culture/internal/app"
	"github.com/antenyijoseph20-glitch/bellz_culture/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Printf("application stopped with an error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}

	server, err := app.NewServer(cfg)
	if err != nil {
		return fmt.Errorf("initialize application: %w", err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Bellz Culture server starting on port %s", cfg.Port)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("HTTP server: %w", err)
		}
		return nil

	case <-ctx.Done():
		log.Println("Shutdown signal received; stopping server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("graceful shutdown: %w", err)
		}

		if err := <-serverErrors; err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("HTTP server during shutdown: %w", err)
		}

		log.Println("Bellz Culture server stopped cleanly")
		return nil
	}
}
