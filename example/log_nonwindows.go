//+build !windows

package main

import (
	"log"
	"syscall"
)

func MessageBox(caption, text string, style uintptr) (result int, callErr syscall.Errno) {
	log.Printf("%s: %s", caption, text)
	return
}

func freeMessageBoxLibraries() {
}
