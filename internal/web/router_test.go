package web

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/antenyijoseph20-glitch/bellz_culture/internal/assets"
)

func newTestRouter(t *testing.T) http.Handler {
	t.Helper()

	imageDir := t.TempDir()

	err := os.WriteFile(
		filepath.Join(imageDir, "test-image.jpeg"),
		[]byte("test image content"),
		0600,
	)
	if err != nil {
		t.Fatalf("create test image: %v", err)
	}

	router, err := NewRouter(
		assets.Templates(),
		assets.CSS(),
		imageDir,
	)
	if err != nil {
		t.Fatalf("create router: %v", err)
	}

	return router
}

func TestHomepage(t *testing.T) {
	router := newTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); !strings.Contains(got, "text/html") {
		t.Fatalf("expected HTML content type, got %q", got)
	}

	body := rec.Body.String()

	if !strings.Contains(body, "Bellz Culture") {
		t.Fatal("homepage does not contain the brand name")
	}

	if !strings.Contains(body, "Signature Striped Utility Tote") {
		t.Fatal("homepage does not contain the expected product")
	}
}

func TestHomepageHead(t *testing.T) {
	router := newTestRouter(t)

	req := httptest.NewRequest(http.MethodHead, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.Len() != 0 {
		t.Fatalf("HEAD response should have no body, got %d bytes", rec.Body.Len())
	}
}

func TestEmbeddedCSS(t *testing.T) {
	router := newTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/static/css/style.css", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "RESPONSIVE LAYOUT") {
		t.Fatal("CSS response does not contain the responsive styles")
	}
}

func TestProductImage(t *testing.T) {
	router := newTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/static/images/test-image.jpeg", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "test image content" {
		t.Fatalf("unexpected image response body: %q", rec.Body.String())
	}
}

func TestRouterRejectsInvalidImageDirectory(t *testing.T) {
	_, err := NewRouter(
		assets.Templates(),
		assets.CSS(),
		filepath.Join(t.TempDir(), "missing"),
	)

	if err == nil {
		t.Fatal("expected an error for a missing image directory")
	}
}

func TestRouterRejectsFileAsImageDirectory(t *testing.T) {
	file := filepath.Join(t.TempDir(), "not-a-directory")

	if err := os.WriteFile(file, []byte("not a directory"), 0600); err != nil {
		t.Fatalf("create test file: %v", err)
	}

	_, err := NewRouter(
		assets.Templates(),
		assets.CSS(),
		file,
	)

	if err == nil {
		t.Fatal("expected an error when image path is a file")
	}
}

// Ensure the router accepts a filesystem abstraction for embedded CSS.
var _ fs.FS = fstest.MapFS{}

func TestImageHandlerSecurityHeaders(t *testing.T) {
	imageDir := t.TempDir()

	if err := os.WriteFile(
		filepath.Join(imageDir, "photo.jpeg"),
		[]byte("image content"),
		0600,
	); err != nil {
		t.Fatalf("create image: %v", err)
	}

	router, err := NewRouter(assets.Templates(), assets.CSS(), imageDir)
	if err != nil {
		t.Fatalf("create router: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/static/images/photo.jpeg", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); got != "image/jpeg" {
		t.Errorf("expected image/jpeg, got %q", got)
	}

	if got := rec.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("expected nosniff, got %q", got)
	}

	if got := rec.Header().Get("Cache-Control"); got != "public, max-age=3600" {
		t.Errorf("unexpected Cache-Control: %q", got)
	}
}

func TestImageHandlerRejectsUnsafeFilesAndPaths(t *testing.T) {
	imageDir := t.TempDir()

	files := []string{
		"notes.txt",
		"vector.svg",
		".secret.jpeg",
	}

	for _, name := range files {
		if err := os.WriteFile(
			filepath.Join(imageDir, name),
			[]byte("private content"),
			0600,
		); err != nil {
			t.Fatalf("create test file %q: %v", name, err)
		}
	}

	nestedDir := filepath.Join(imageDir, "nested")
	if err := os.Mkdir(nestedDir, 0700); err != nil {
		t.Fatalf("create nested directory: %v", err)
	}

	if err := os.WriteFile(
		filepath.Join(nestedDir, "image.jpeg"),
		[]byte("nested content"),
		0600,
	); err != nil {
		t.Fatalf("create nested image: %v", err)
	}

	router, err := NewRouter(assets.Templates(), assets.CSS(), imageDir)
	if err != nil {
		t.Fatalf("create router: %v", err)
	}

	paths := []string{
		"/static/images/notes.txt",
		"/static/images/vector.svg",
		"/static/images/.secret.jpeg",
		"/static/images/nested/image.jpeg",
		"/static/images/",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			if rec.Code == http.StatusOK {
				t.Fatalf("unsafe path %q was served", path)
			}

			if strings.Contains(rec.Body.String(), "private content") ||
				strings.Contains(rec.Body.String(), "nested content") {
				t.Fatalf("private file content exposed for %q", path)
			}
		})
	}
}

func TestImageHandlerRejectsSymlinkEscape(t *testing.T) {
	imageDir := t.TempDir()
	outsideDir := t.TempDir()

	outsideFile := filepath.Join(outsideDir, "private.jpeg")
	if err := os.WriteFile(
		outsideFile,
		[]byte("private external content"),
		0600,
	); err != nil {
		t.Fatalf("create external file: %v", err)
	}

	linkPath := filepath.Join(imageDir, "outside.jpeg")
	if err := os.Symlink(outsideFile, linkPath); err != nil {
		t.Skipf("symbolic links unavailable: %v", err)
	}

	router, err := NewRouter(assets.Templates(), assets.CSS(), imageDir)
	if err != nil {
		t.Fatalf("create router: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/static/images/outside.jpeg", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal("handler served a file outside the image directory")
	}

	if strings.Contains(rec.Body.String(), "private external content") {
		t.Fatal("external file content was exposed")
	}
}

func TestRouterReturnsNotFoundForUnknownRoute(t *testing.T) {
	router := newTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/this-route-does-not-exist", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Security-Policy"); got == "" {
		t.Error("security headers should also be present on 404 responses")
	}
}

func TestRouterRejectsUnsupportedMethod(t *testing.T) {
	router := newTestRouter(t)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rec.Code)
	}

	allow := rec.Header().Get("Allow")

	if !strings.Contains(allow, "GET") {
		t.Errorf("Allow header should include GET, got %q", allow)
	}

	if !strings.Contains(allow, "HEAD") {
		t.Errorf("Allow header should include HEAD, got %q", allow)
	}

	if got := rec.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Errorf("expected nosniff security header, got %q", got)
	}
}
