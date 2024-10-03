//go:build !windows && !darwin && !linux
// +build !windows,!darwin,!linux

/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/

package trashbox

import "errors"

func MoveToTrash(path string) error {
	return errors.New("move to trash bin or recycle bin is not supported on this platform")
}
