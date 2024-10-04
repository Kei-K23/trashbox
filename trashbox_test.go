//go:build darwin
// +build darwin

/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/

// ! Note, I use Mac OS when developing this package, so for testing I use built constraint for Mac OS (darwin). Use your built constraint depending your OS to run test.
package trashbox

import (
	"os"
	"testing"
)

func TestMoveToTrash(t *testing.T) {
	testFile, err := os.CreateTemp("", "this_is_test_trash_new_for_recovery")
	if err != nil {
		t.Fatalf("unable to create temp file: %v", err)
	}
	defer os.Remove(testFile.Name())
	err = MoveToTrash(testFile.Name())
	// err = MoveToTrash(testFile.Name())
	if err != nil {
		t.Fatalf("failed to move file to trash: %v", err)
	}
}
