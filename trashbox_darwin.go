//go:build darwin
// +build darwin

/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/

package trashbox

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Metadata struct {
	OriginalPath string `json:"original_path"`
}

// MoveToTrash moves the specified file or directory to the macOS Trash.
//
// This function takes the path of a file or directory as an argument,
// converts it to an absolute path, and then moves it to the user's
// Trash directory (located at ~/.Trash). If the provided path does
// not exist or cannot be accessed, an error is returned.
//
// It is important to note that this function does not handle any
// conflicts that may arise if a file with the same name already exists
// in the Trash. If the operation is successful, the file or directory
// will no longer exist at the original path and will be relocated to
// the Trash for potential recovery.
//
// Parameters:
//   - path: The path of the file or directory to be moved to Trash.
//
// Returns:
//   - error: Returns nil on success. If an error occurs during the
//     process (e.g., if the file does not exist or the move fails),
//     an error will be returned explaining the reason for failure.
//
// Example:
//
//	err := MoveToTrash("path/to/your/file.txt")
//	if err != nil {
//	    log.Fatalf("Failed to move to Trash: %v", err)
//	}
func MoveToTrash(path string) error {
	// Get the absolute file path of delete file
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Get the trash file path to move to .Trash directory
	trashPath := filepath.Join(os.Getenv("HOME"), ".Trash", filepath.Base(path))
	// Move the file to .Trash directory
	err = os.Rename(absPath, trashPath)
	if err != nil {
		return err
	}

	// Create metadata file for recovery the deleted file
	metadata := Metadata{OriginalPath: absPath}
	metadataPath := trashPath + ".metadata.json"

	// Create metadata file in Trash bin
	metadataFile, err := os.Create(metadataPath)
	if err != nil {
		return err
	}
	defer metadataFile.Close()

	encoder := json.NewEncoder(metadataFile)
	err = encoder.Encode(metadata)
	if err != nil {
		return err
	}

	// Process is success and return nill
	return nil
}
