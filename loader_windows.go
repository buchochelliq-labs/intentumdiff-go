//go:build windows

package intentumdiff

import (
	"fmt"
	"syscall"
	"unsafe"
)

// nativeLib on Windows is stdlib-only: LoadLibrary + raw proc calls (no cgo).
type nativeLib struct {
	dll  *syscall.DLL
	call_ *syscall.Proc
	free *syscall.Proc
}

func openNative(path string) (nativeLib, error) {
	dll, err := syscall.LoadDLL(path)
	if err != nil {
		return nativeLib{}, err
	}
	callProc, err := dll.FindProc("intentumdiff_call")
	if err != nil {
		return nativeLib{}, fmt.Errorf("intentumdiff_call not exported: %w", err)
	}
	freeProc, err := dll.FindProc("intentumdiff_free")
	if err != nil {
		return nativeLib{}, fmt.Errorf("intentumdiff_free not exported: %w", err)
	}
	return nativeLib{dll: dll, call_: callProc, free: freeProc}, nil
}

func (l nativeLib) call(name, argsJSON string) (string, error) {
	cName, err := syscall.BytePtrFromString(name)
	if err != nil {
		return "", err
	}
	cArgs, err := syscall.BytePtrFromString(argsJSON)
	if err != nil {
		return "", err
	}
	ret, _, _ := l.call_.Call(
		uintptr(unsafe.Pointer(cName)),
		uintptr(unsafe.Pointer(cArgs)),
	)
	if ret == 0 {
		return "", fmt.Errorf("intentumdiff_call returned NULL (allocation/encoding failure)")
	}
	// Copy the C string out, then free the engine allocation.
	out := goString(ret)
	l.free.Call(ret)
	return out, nil
}

// goString copies a NUL-terminated C string at ptr into a Go string. The single
// uintptr->Pointer conversion below is a FOREIGN heap pointer returned by the engine
// (never Go memory), which go vet's unsafeptr heuristic cannot know — CI runs vet with
// -unsafeptr=false for exactly this file's sake; all walking uses unsafe.Add.
func goString(ptr uintptr) string {
	base := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	length := 0
	for {
		b := *(*byte)(unsafe.Add(base, length))
		if b == 0 {
			break
		}
		length++
	}
	return string(unsafe.Slice((*byte)(base), length))
}
