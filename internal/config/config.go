package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port     string
	ImageDir string
}

func Load() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portNumber, err := strconv.Atoi(port)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return Config{}, fmt.Errorf("PORT must be a valid number between 1 and 65535")
	}

	imageDir := os.Getenv("IMAGE_DIR")
	if imageDir == "" {
		imageDir = "static/images"
	}

	return Config{
		Port:     port,
		ImageDir: imageDir,
	}, nil
}
