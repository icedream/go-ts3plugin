//+build windows

package main

import (
	"syscall"
	"unsafe"
)

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")

	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

func MessageBox(caption, text string, style uintptr) (result int, callErr syscall.Errno) {
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall9(uintptr(messageBox),
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0,
		0,
		0,
		0)
	if callErr != 0 {
		return
	}
	result = int(ret)
	return
}

/*func GetModuleHandle() (handle uintptr, err syscall.Errno) {
	var nargs uintptr = 0
	if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0); callErr != 0 {
		err = callErr
	} else {
		handle = ret
	}
	return
}
*/

func freeMessageBoxLibraries() {
	defer syscall.FreeLibrary(kernel32)
	defer syscall.FreeLibrary(user32)
}
