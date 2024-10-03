//go:build linux
// +build linux

/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/

package trashbox

import (
	"errors"
	"os"
	"path/filepath"
)

// MoveToTrash moves the specified file or directory to the Linux Trash.
//
// This function takes the path of a file or directory as an argument,
// converts it to an absolute path, and then moves it to the user's
// Trash directory located at ~/.local/share/Trash/files. If the
// provided path does not exist, or if the Trash directory does not
// exist, an error will be returned.
//
// If the operation is successful, the file or directory will no
// longer exist at the original path and will be relocated to the
// Trash for potential recovery.
//
// Parameters:
//   - path: The path of the file or directory to be moved to Trash.
//
// Returns:
//   - error: Returns nil on success. If an error occurs during the
//     process (e.g., if the file does not exist or the move fails),
//     an error will be returned explaining the reason for failure.
//     Specifically, if the Trash directory does not exist, an
//     error indicating that will be returned.
//
// Example:
//
//	err := MoveToTrash("path/to/your/file.txt")
//	if err != nil {
//	    log.Fatalf("Failed to move to Trash: %v", err)
//	}
func MoveToTrash(path string) error {
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
