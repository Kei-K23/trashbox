package trashbox

import (
	"os"
	"testing"
)

func TestMoveToTrash(t *testing.T) {
	testFile, err := os.CreateTemp("", "this_is_test_trash_new")
	if err != nil {
		t.Fatalf("unable to create temp file: %v", err)
	}
	defer os.Remove(testFile.Name())
	err = MoveToTrashMacOS(testFile.Name())
	// err = MoveToTrash(testFile.Name())
	if err != nil {
		t.Fatalf("failed to move file to trash: %v", err)
	}
}
