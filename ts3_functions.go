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
	cResult := (*C.char)(C.malloc(256))
	defer C.free(unsafe.Pointer(cResult))

	code := C.getClientLibVersion(this.nativeFunctions, &cResult)
	if ok = code == 0; ok {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) GetClientLibVersionNumber() (result uint64, ok bool) {
	var cResult C.uint64

	code := C.getClientLibVersionNumber(this.nativeFunctions, &cResult)
	if ok = code == 0; ok {
		result = uint64(cResult)
	}

	return
}
