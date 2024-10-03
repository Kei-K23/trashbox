//go:build windows
// +build windows

package trashbox

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	FO_DELETE          = 3    // File operation: delete
	FOF_ALLOWUNDO      = 0x40 // Allow to move to Recycle Bin
	FOF_NOCONFIRMATION = 0x10 // No confirmation dialog
)

// SHFILEOPSTRUCT represents the structure used in SHFileOperationW.
type SHFILEOPSTRUCT struct {
	WFunc             uint32
	PFrom             *uint16
	PTo               *uint16
	FFlags            uint16
	AnyOps            bool
	HNameMap          uintptr
	LpszProgressTitle *uint16
}

var (
	shell32              = syscall.NewLazyDLL("shell32.dll")
	procSHFileOperationW = shell32.NewProc("SHFileOperationW")
)

func shFileOperation(op *SHFILEOPSTRUCT) error {
	ret, _, _ := procSHFileOperationW.Call(uintptr(unsafe.Pointer(op)))
	if ret != 0 {
		return fmt.Errorf("failed to move file to Recycle Bin, error code: %d", ret)
	}
	return nil
}

func moveToTrashWindows(path string) error {
	// Get the absolute file path of delete file
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	wPathPtr, err := windows.UTF16PtrFromString(absPath + "\x00")
	if err != nil {
		return err
	}

	op := &SHFILEOPSTRUCT{
		WFunc:  FO_DELETE,
		PFrom:  wPathPtr,
		FFlags: FOF_ALLOWUNDO | FOF_NOCONFIRMATION,
	}

	return shFileOperation(op)
}
