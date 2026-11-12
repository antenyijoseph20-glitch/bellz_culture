package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityHeaders(t *testing.T) {
	handler := securityHeaders(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	tests := map[string]string{
		"Content-Security-Policy": "default-src 'self'; script-src 'self'; style-src 'self' https://fonts.googleapis.com; font-src 'self' https://fonts.gstatic.com; img-src 'self' data:; connect-src 'self'; object-src 'none'; base-uri 'self'; form-action 'self'; frame-ancestors 'none'",
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":         "DENY",
		"Referrer-Policy":         "strict-origin-when-cross-origin",
		"Permissions-Policy":      "camera=(), microphone=(), geolocation=()",
	}

	for header, want := range tests {
		t.Run(header, func(t *testing.T) {
			if got := rec.Header().Get(header); got != want {
				t.Errorf("expected %s %q, got %q", header, want, got)
			}
		})
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

func TestRouterAppliesSecurityHeadersToAllRoutes(t *testing.T) {
	router := newTestRouter(t)

	paths := []string{
		"/",
		"/static/css/style.css",
		"/static/images/test-image.jpeg",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected status 200 for %s, got %d", path, rec.Code)
			}

			if got := rec.Header().Get("Content-Security-Policy"); got == "" {
				t.Errorf("Content-Security-Policy missing on %s", path)
			}

			if got := rec.Header().Get("X-Frame-Options"); got != "DENY" {
				t.Errorf("expected X-Frame-Options DENY on %s, got %q", path, got)
			}

			if got := rec.Header().Get("X-Content-Type-Options"); got != "nosniff" {
				t.Errorf("expected X-Content-Type-Options nosniff on %s, got %q", path, got)
			}
		})
	}
}
