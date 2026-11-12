package assets

import (
	"embed"
	"io/fs"
)

//go:embed templates/*.html static/css/*.css
var embeddedFiles embed.FS

// Templates returns the embedded HTML templates.
func Templates() fs.FS {
	files, err := fs.Sub(embeddedFiles, "templates")
	if err != nil {
		panic("embedded templates are unavailable: " + err.Error())
	}

	return files
}

// CSS returns the embedded CSS assets.
func CSS() fs.FS {
	files, err := fs.Sub(embeddedFiles, "static/css")
	if err != nil {
		panic("embedded CSS assets are unavailable: " + err.Error())
	}

	return files
}
