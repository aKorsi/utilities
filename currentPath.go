package utilities

import (
	"os"
	"path/filepath"
)

func CurrentPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}
