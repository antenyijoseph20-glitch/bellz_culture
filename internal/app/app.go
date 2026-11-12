package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antenyijoseph20-glitch/bellz_culture/internal/assets"
	"github.com/antenyijoseph20-glitch/bellz_culture/internal/config"
	"github.com/antenyijoseph20-glitch/bellz_culture/internal/web"
)

// NewServer assembles the application and returns its HTTP server.
func NewServer(cfg config.Config) (*http.Server, error) {
	router, err := web.NewRouter(
		assets.Templates(),
		assets.CSS(),
		cfg.ImageDir,
	)
	if err != nil {
		return nil, fmt.Errorf("create application router: %w", err)
	}

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MiB
	}

	return server, nil
}
