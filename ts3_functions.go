package ts3plugin

//go:generate go run ./internal/generate/ts3_functions/main.go

/*
#cgo CFLAGS: -I./pluginsdk/include -I.

#include <stdlib.h> // free
#include "generated_ts3_function_wrappers.h"

*/
import "C"
import "unsafe"

/* ts3_functions.h */

type TS3Functions struct {
	nativeFunctions C.TS3Functions
}

func (this *TS3Functions) GetClientLibVersion() (result string, ok bool) {
	// create buffer for teamspeak code to write to
	buf := (*C.char)(C.malloc(256))
	defer C.free(unsafe.Pointer(buf))

	code := C.getClientLibVersion(this.nativeFunctions, &buf)
	if ok = code == 0; ok {
		result = C.GoString(buf)
	}

	return
}
