package ts3plugin

//go:generate go run ./internal/generate/ts3_functions/main.go

/*
#cgo CFLAGS: -I./pluginsdk/include -I.

#include <stdlib.h> // free
#include "generated_ts3_function_wrappers.h"

*/
import "C"

import (
	"unsafe"

	"github.com/icedream/go-ts3plugin/teamlog"
	"github.com/icedream/go-ts3plugin/teamspeak"
)

/* ts3_functions.h */

type TS3Functions struct {
	nativeFunctions C.TS3Functions
}

func (this *TS3Functions) GetClientLibVersion() (result string, errorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(256))
	defer C.free(unsafe.Pointer(cResult))

	errorCode = uint32(C.getClientLibVersion(this.nativeFunctions, &cResult))
	if errorCode == 0 {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) GetClientLibVersionNumber() (result uint64, errorCode uint32) {
	var cResult C.uint64

	errorCode = uint32(C.getClientLibVersionNumber(this.nativeFunctions, &cResult))
	if errorCode == 0 {
		result = uint64(cResult)
	}

	return
}

func (this *TS3Functions) SpawnNewServerConnectionHandler(port int) (result uint64, errorCode uint32) {
	var cResult C.uint64

	errorCode = uint32(C.spawnNewServerConnectionHandler(this.nativeFunctions, C.int(port), &cResult))
	if errorCode == 0 {
		result = uint64(cResult)
	}

	return
}

func (this *TS3Functions) DestroyServerConnectionHandler(serverConnectionHandlerID uint64) (errorCode uint32) {
	errorCode = uint32(C.destroyServerConnectionHandler(this.nativeFunctions, C.uint64(serverConnectionHandlerID)))

	return
}

func (this *TS3Functions) GetErrorMessage(errorCode uint32) (result string, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getErrorMessage(this.nativeFunctions, C.uint(errorCode), &cResult))
	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) LogMessage(logMessage string, severity teamlog.LogLevel, channel string, logID uint64) (retErrorCode uint32) {
	cLogMessage := C.CString(logMessage)
	defer C.free(unsafe.Pointer(cLogMessage))

	cChannel := C.CString(channel)
	defer C.free(unsafe.Pointer(cChannel))

	retErrorCode = uint32(C.logMessage(this.nativeFunctions, cLogMessage,
		uint32(severity), cChannel, C.uint64(logID)))

	return
}

func (this *TS3Functions) PrintMessage(serverConnectionHandlerID uint64, message string, messageTarget PluginMessageTarget) {
	messageC := C.CString(message)
	defer C.free(unsafe.Pointer(messageC))

	C.printMessage(this.nativeFunctions, C.uint64(serverConnectionHandlerID), messageC, uint32(messageTarget))

	return
}

func (this *TS3Functions) PrintMessageToCurrentTab(message string) {
	messageC := C.CString(message)
	defer C.free(unsafe.Pointer(messageC))

	C.printMessageToCurrentTab(this.nativeFunctions, messageC)

	return
}

func (this *TS3Functions) GetPreProcessorConfigValue(serverConnectionHandlerID uint64, ident string) (result string, retErrorCode uint32) {
	identC := C.CString(ident)

	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getPreProcessorConfigValue(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		identC,
		&cResult,
	))
	if retErrorCode == 0 {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) SetPreProcessorConfigValue(serverConnectionHandlerID uint64, ident string, value string) (retErrorCode uint32) {
	identC := C.CString(ident)
	valueC := C.CString(value)

	retErrorCode = uint32(C.setPreProcessorConfigValue(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		identC,
		valueC,
	))

	return
}

func (this *TS3Functions) RequestSendPrivateTextMsg(serverConnectionHandlerID uint64, message string, targetClientID teamspeak.AnyID, customErrorIdentifier string) (retErrorCode uint32) {
	messageC := C.CString(message)
	var customErrorIdentifierC *C.char
	if len(customErrorIdentifier) > 0 {
		customErrorIdentifierC = C.CString(customErrorIdentifier)
	}

	retErrorCode = uint32(C.requestSendPrivateTextMsg(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		messageC,
		C.anyID(targetClientID),
		customErrorIdentifierC,
	))

	return
}

func (this *TS3Functions) RequestSendChannelTextMsg(serverConnectionHandlerID uint64, message string, targetChannelID uint64, customErrorIdentifier string) (retErrorCode uint32) {
	messageC := C.CString(message)
	var customErrorIdentifierC *C.char
	if len(customErrorIdentifier) > 0 {
		customErrorIdentifierC = C.CString(customErrorIdentifier)
	}

	retErrorCode = uint32(C.requestSendChannelTextMsg(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		messageC,
		C.uint64(targetChannelID),
		customErrorIdentifierC,
	))

	return
}
