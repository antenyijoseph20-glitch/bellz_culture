package assets

import (
	"io/fs"
	"testing"
)

func TestEmbeddedAssetsExist(t *testing.T) {
	tests := []struct {
		name string
		fsys fs.FS
		file string
	}{
		{
			name: "homepage template",
			fsys: Templates(),
			file: "index.html",
		},
		{
			name: "stylesheet",
			fsys: CSS(),
			file: "style.css",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := fs.Stat(tt.fsys, tt.file)
			if err != nil {
				t.Fatalf("asset %q unavailable: %v", tt.file, err)
			}

			if info.IsDir() || info.Size() == 0 {
				t.Fatalf("asset %q is invalid", tt.file)
			}
		})
	}
}
