package internal

import (
	"errors"
	"strings"
)

func IsValidModelFileName(filename string) bool {
	// basic validation of extension
	l := strings.ToLower(filename)
	return strings.HasSuffix(l, ".glb") || strings.HasSuffix(l, ".gltf")
}

var ErrInvalidModel = errors.New("invalid model file; allowed: .glb, .gltf")
