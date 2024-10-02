//go:build darwin
// +build darwin

package trashbox

import (
	"os"
	"path/filepath"
)

func moveToTrashMacOS(path string) error {
	// Get the absolute file path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Get the trash file path to move to .Trash directory
	trashPath := filepath.Join(os.Getenv("HOME"), ".Trash", filepath.Base(path))
	// Move the file to .Trash directory
	return os.Rename(absPath, trashPath)
}
