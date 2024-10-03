package trashbox

import (
	"errors"
	"runtime"
)

//--------------------------------//
// Public package APIs functions  //
//--------------------------------//

func MoveToTrash(path string) error {
	switch runtime.GOOS {
	case "windows":
		// Handle for
		return moveToTrashWindows(path)
	case "darwin":
		// Handle for Mac OS
		return moveToTrashMacOS(path)
	case "linux":
		// Handle for Linux
		return moveToTrashLinux(path)
	default:
		return errors.New("unsupported platform")
	}
}

// -----------------------//
// Private lib functions //
// -----------------------//
