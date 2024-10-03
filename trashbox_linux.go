//go:build linux
// +build linux

package trashbox

import (
	"errors"
	"os"
	"path/filepath"
)

func MoveToTrashLinux(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Get the file path to trash file
	trashDir := filepath.Join(os.Getenv("HOME"), ".local/share/Trash/files")
	if _, err = os.Stat(trashDir); os.IsNotExist(err) {
		return errors.New("trash directory does not exist")
	}

	trashPath := filepath.Join(trashDir, filepath.Base(absPath))
	// Move the file to trash directory
	return os.Rename(absPath, trashPath)
}
