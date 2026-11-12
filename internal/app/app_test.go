package app

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/antenyijoseph20-glitch/bellz_culture/internal/config"
)

func TestNewServer(t *testing.T) {
	imageDir := t.TempDir()

	cfg := config.Config{
		Port:     "8080",
		ImageDir: imageDir,
	}

	server, err := NewServer(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if server.Addr != ":8080" {
		t.Errorf("expected address :8080, got %q", server.Addr)
	}

	if server.Handler == nil {
		t.Fatal("expected an application handler")
	}

	if server.ReadHeaderTimeout != 5*time.Second {
		t.Errorf("unexpected read header timeout: %v", server.ReadHeaderTimeout)
	}

	if server.ReadTimeout != 15*time.Second {
		t.Errorf("unexpected read timeout: %v", server.ReadTimeout)
	}

	if server.WriteTimeout != 30*time.Second {
		t.Errorf("unexpected write timeout: %v", server.WriteTimeout)
	}

	if server.IdleTimeout != 60*time.Second {
		t.Errorf("unexpected idle timeout: %v", server.IdleTimeout)
	}

	if server.MaxHeaderBytes != 1<<20 {
		t.Errorf("unexpected maximum header size: %d", server.MaxHeaderBytes)
	}
}

func TestNewServerRejectsMissingImageDirectory(t *testing.T) {
	cfg := config.Config{
		Port:     "8080",
		ImageDir: filepath.Join(t.TempDir(), "missing"),
	}

	_, err := NewServer(cfg)
	if err == nil {
		t.Fatal("expected an error for a missing image directory")
	}
}

func TestNewServerRejectsFileAsImageDirectory(t *testing.T) {
	imageFile := filepath.Join(t.TempDir(), "not-a-directory")

	if err := os.WriteFile(imageFile, []byte("test"), 0600); err != nil {
		t.Fatalf("create test file: %v", err)
	}

	cfg := config.Config{
		Port:     "8080",
		ImageDir: imageFile,
	}

	_, err := NewServer(cfg)
	if err == nil {
		t.Fatal("expected an error when image path is a file")
	}
}

func TestServerHandlerServesHomepage(t *testing.T) {
	cfg := config.Config{
		Port:     "8080",
		ImageDir: t.TempDir(),
	}

	server, err := NewServer(cfg)
	if err != nil {
		t.Fatalf("create server: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}

	recorder := &responseRecorder{header: make(http.Header)}

	server.Handler.ServeHTTP(recorder, req)

	if recorder.statusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.statusCode)
	}
}

type responseRecorder struct {
	header     http.Header
	statusCode int
	body       []byte
}

func (r *responseRecorder) Header() http.Header {
	return r.header
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func (r *responseRecorder) Write(data []byte) (int, error) {
	if r.statusCode == 0 {
		r.statusCode = http.StatusOK
	}

	r.body = append(r.body, data...)
	return len(data), nil
}
