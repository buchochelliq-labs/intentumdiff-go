//go:build !windows

package intentumdiff

import "errors"

// nativeLib on non-Windows platforms: dlopen-based loading lands with the cgo (or
// purego) seam — the scaffold proves the ABI on Windows first; the API above is
// platform-independent.
type nativeLib struct{}

var errUnsupported = errors.New("intentumdiff: non-Windows loader not yet implemented (dlopen seam pending)")

func openNative(path string) (nativeLib, error) { return nativeLib{}, errUnsupported }

func (l nativeLib) call(name, argsJSON string) (string, error) { return "", errUnsupported }
