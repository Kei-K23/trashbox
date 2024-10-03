//go:build !windows && !darwin && !linux
// +build !windows,!darwin,!linux

package trashbox

import "errors"

func MoveToTrashWindows(path string) error {
	return errors.New("move to trash is not supported on this platform")
}

func MoveToTrashMacOS(path string) error {
	return errors.New("move to trash is not supported on this platform")
}

func MoveToTrashLinux(path string) error {
	return errors.New("move to trash is not supported on this platform")
}
