package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("IMAGE_DIR", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Port != "8080" {
		t.Errorf("expected port 8080, got %q", cfg.Port)
	}

	if cfg.ImageDir != "static/images" {
		t.Errorf("expected static/images, got %q", cfg.ImageDir)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	tests := []string{"abc", "0", "-1", "65536"}

	for _, port := range tests {
		t.Run(port, func(t *testing.T) {
			t.Setenv("PORT", port)

			_, err := Load()
			if err == nil {
				t.Fatalf("expected port %q to be rejected", port)
			}
		})
	}
}

func TestLoadCustomValues(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("IMAGE_DIR", "/tmp/bellz-images")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Port != "9090" {
		t.Errorf("expected port 9090, got %q", cfg.Port)
	}

	if cfg.ImageDir != "/tmp/bellz-images" {
		t.Errorf("unexpected image directory: %q", cfg.ImageDir)
	}
}
