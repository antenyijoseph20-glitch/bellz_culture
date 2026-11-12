package web

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/antenyijoseph20-glitch/bellz_culture/internal/web/handlers"
)

var allowedImageTypes = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".webp": "image/webp",
	".gif":  "image/gif",
}

// NewRouter creates the application's HTTP routes.
func NewRouter(templateFS, cssFS fs.FS, imageDir string) (http.Handler, error) {
	if imageDir == "" {
		return nil, fmt.Errorf("image directory must be configured")
	}

	absoluteImageDir, err := filepath.Abs(imageDir)
	if err != nil {
		return nil, fmt.Errorf("resolve image directory: %w", err)
	}

	info, err := os.Stat(absoluteImageDir)
	if err != nil {
		return nil, fmt.Errorf("inspect image directory: %w", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("image path is not a directory: %s", absoluteImageDir)
	}

	homeHandler, err := handlers.NewHomeHandler(templateFS)
	if err != nil {
		return nil, fmt.Errorf("create homepage handler: %w", err)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", homeHandler)

	cssHandler := http.StripPrefix(
		"/static/css/",
		http.FileServer(http.FS(cssFS)),
	)
	mux.Handle("GET /static/css/", cssHandler)

	mux.Handle("GET /static/images/", newImageHandler(absoluteImageDir))

	return securityHeaders(mux), nil
}

// newImageHandler serves only approved image files from the configured directory.
func newImageHandler(imageDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/static/images/")

		if name == "" ||
			strings.ContainsAny(name, `/\`) ||
			strings.HasPrefix(name, ".") ||
			filepath.Base(name) != name {
			http.NotFound(w, r)
			return
		}

		contentType, allowed := allowedImageTypes[strings.ToLower(filepath.Ext(name))]
		if !allowed {
			http.NotFound(w, r)
			return
		}

		root, err := os.OpenRoot(imageDir)
		if err != nil {
			http.Error(w, "Image unavailable", http.StatusInternalServerError)
			return
		}
		defer root.Close()

		file, err := root.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil || !info.Mode().IsRegular() {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", contentType)
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Cache-Control", "public, max-age=3600")

		http.ServeContent(w, r, name, info.ModTime(), file)
	})
}
