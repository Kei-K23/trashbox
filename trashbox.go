package trashbox

import (
	"errors"
	"runtime"
)

func MoveToTrash(path string) error {
	switch runtime.GOOS {
	case "windows":
		// Handle for windows
		return nil
	case "darwin":
		// Handle for Mac OS
		return moveToTrashMacOS(path)
	case "linux":
		// Handle for Linux
		return nil
	default:
		return errors.New("unsupported platform")
	}
}
