package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ValidateModelPath(path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".glb" && ext != ".gltf" {
		return fmt.Errorf("invalid model format: only .glb or .gltf allowed")
	}
	if path == "" {
		return fmt.Errorf("model_path cannot be empty")
	}
	return nil
}
