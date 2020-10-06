package ts3plugin

//go:generate go run ./internal/generate/ts3_functions/main.go

/*
#cgo CFLAGS: -I./pluginsdk/include -I.

#include <stdlib.h> // free
#include "generated_ts3_function_wrappers.h"

*/
import "C"

import (
	"math"
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

func (this *TS3Functions) GetPlaybackDeviceList(modeID string) (result [][]string, retErrorCode uint32) {
	cModeID := C.CString(modeID)
	defer C.free(unsafe.Pointer(cModeID))

	// actually [][]*C.char
	cResult := (***C.char)(nil)

	retErrorCode = uint32(C.getPlaybackDeviceList(this.nativeFunctions,
		cModeID, &cResult))

	if retErrorCode == teamspeak.ErrorOK {
		// ***C.char => []**C.char
		// find correct length
		for len := 1; len < math.MaxInt32; len++ {
			/*
				var theCArray *C.YourType = C.getTheArray()
				length := C.getTheArrayLength()
				slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]

				C.YourType = **C.char
			*/
			cResultSlice := (*[1 << 28]**C.char)(unsafe.Pointer(cResult))[:len:len]
			if cResultSlice[len-1] != nil {
				continue
			}
			len--
			cResultSlice = cResultSlice[0:len]

			// **C.char => []*C.char
			result = make([][]string, len)
			for i, cResultItem := range cResultSlice {
				// assume 2 *C.char here due to how device lists work
				cResultItemSlice := (*[1 << 28]*C.char)(unsafe.Pointer(cResultItem))[:2:2]
				// *C.char => string
				result[i] = []string{
					C.GoString(cResultItemSlice[0]),
					C.GoString(cResultItemSlice[1]),
				}
			}
			break
		}
	}

	return
}

func (this *TS3Functions) GetPlaybackModeList() (result []string, retErrorCode uint32) {
	// actually []*C.char
	cResult := (**C.char)(nil)

	retErrorCode = uint32(C.getPlaybackModeList(this.nativeFunctions,
		&cResult))

	if retErrorCode == teamspeak.ErrorOK {
		// **C.char => []*C.char
		// find correct length
		for len := 1; len < math.MaxInt32; len++ {
			/*
				var theCArray *C.YourType = C.getTheArray()
				length := C.getTheArrayLength()
				slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]

				C.YourType = *C.char
			*/
			cResultSlice := (*[1 << 28]*C.char)(unsafe.Pointer(cResult))[:len:len]
			if cResultSlice[len-1] != nil {
				continue
			}
			len--
			cResultSlice = cResultSlice[0:len]

			// *C.char => string
			result = make([]string, len)
			for i, cResultItem := range cResultSlice {
				result[i] = C.GoString(cResultItem)
			}
			break
		}
	}

	return
}

func (this *TS3Functions) GetCaptureDeviceList(modeID string) (result [][]string, retErrorCode uint32) {
	cModeID := C.CString(modeID)
	defer C.free(unsafe.Pointer(cModeID))

	// actually [][]*C.char
	cResult := (***C.char)(nil)

	retErrorCode = uint32(C.getCaptureDeviceList(this.nativeFunctions,
		cModeID, &cResult))

	if retErrorCode == teamspeak.ErrorOK {
		// ***C.char => []**C.char
		// find correct length
		for len := 1; len < math.MaxInt32; len++ {
			/*
				var theCArray *C.YourType = C.getTheArray()
				length := C.getTheArrayLength()
				slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]

				C.YourType = **C.char
			*/
			cResultSlice := (*[1 << 28]**C.char)(unsafe.Pointer(cResult))[:len:len]
			if cResultSlice[len-1] != nil {
				continue
			}
			len--
			cResultSlice = cResultSlice[0:len]

			// **C.char => []*C.char
			result = make([][]string, len)
			for i, cResultItem := range cResultSlice {
				// assume 2 *C.char here due to how device lists work
				cResultItemSlice := (*[1 << 28]*C.char)(unsafe.Pointer(cResultItem))[:2:2]
				// *C.char => string
				result[i] = []string{
					C.GoString(cResultItemSlice[0]),
					C.GoString(cResultItemSlice[1]),
				}
			}
			break
		}
	}

	return
}

func (this *TS3Functions) GetCaptureModeList() (result []string, retErrorCode uint32) {
	// actually []*C.char
	cResult := (**C.char)(nil)

	retErrorCode = uint32(C.getCaptureModeList(this.nativeFunctions,
		&cResult))

	if retErrorCode == teamspeak.ErrorOK {
		// **C.char => []*C.char
		// find correct length
		for len := 1; len < math.MaxInt32; len++ {
			/*
				var theCArray *C.YourType = C.getTheArray()
				length := C.getTheArrayLength()
				slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]

				C.YourType = *C.char
			*/
			cResultSlice := (*[1 << 28]*C.char)(unsafe.Pointer(cResult))[:len:len]
			if cResultSlice[len-1] != nil {
				continue
			}
			len--
			cResultSlice = cResultSlice[0:len]

			// *C.char => string
			result = make([]string, len)
			for i, cResultItem := range cResultSlice {
				result[i] = C.GoString(cResultItem)
			}
			break
		}
	}

	return
}

func (this *TS3Functions) GetDefaultPlaybackDevice(modeID string) (result []string, retErrorCode uint32) {
	cModeID := C.CString(modeID)
	defer C.free(unsafe.Pointer(cModeID))

	// actually []*C.char
	cResult := (**C.char)(nil)

	retErrorCode = uint32(C.getDefaultPlaybackDevice(this.nativeFunctions,
		cModeID, &cResult))

	if retErrorCode == teamspeak.ErrorOK {
		// **C.char => []*C.char
		// assume length of 2 for device description
		/*
			var theCArray *C.YourType = C.getTheArray()
			length := C.getTheArrayLength()
			slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]

			C.YourType = *C.char
		*/
		cResultSlice := (*[1 << 28]*C.char)(unsafe.Pointer(cResult))[:2:2]

		// *C.char => string
		result = make([]string, len(cResultSlice))
		for i, cResultItem := range cResultSlice {
			result[i] = C.GoString(cResultItem)
		}
	}

	return
}

func (this *TS3Functions) GetDefaultPlayBackMode() (result string, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getDefaultPlayBackMode(this.nativeFunctions, &cResult))
	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) GetDefaultCaptureDevice(modeID string) (result []string, retErrorCode uint32) {
	cModeID := C.CString(modeID)
	defer C.free(unsafe.Pointer(cModeID))

	// actually []*C.char
	cResult := (**C.char)(nil)

	retErrorCode = uint32(C.getDefaultCaptureDevice(this.nativeFunctions,
		cModeID, &cResult))

	if retErrorCode == teamspeak.ErrorOK {
		// **C.char => []*C.char
		// assume length of 2 for device description
		/*
			var theCArray *C.YourType = C.getTheArray()
			length := C.getTheArrayLength()
			slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]

			C.YourType = *C.char
		*/
		cResultSlice := (*[1 << 28]*C.char)(unsafe.Pointer(cResult))[:2:2]

		// *C.char => string
		result = make([]string, len(cResultSlice))
		for i, cResultItem := range cResultSlice {
			result[i] = C.GoString(cResultItem)
		}
	}

	return
}

func (this *TS3Functions) GetDefaultCaptureMode() (result string, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getDefaultCaptureMode(this.nativeFunctions, &cResult))
	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
	}

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
