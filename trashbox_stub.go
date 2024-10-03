//go:build !windows && !darwin && !linux
// +build !windows,!darwin,!linux

package trashbox

import "errors"

func moveToTrashWindows(path string) error {
	return errors.New("move to trash is not supported on this platform")
}

func moveToTrashMacOS(path string) error {
	return errors.New("move to trash is not supported on this platform")
}

func moveToTrashLinux(path string) error {
	return errors.New("move to trash is not supported on this platform")
}
