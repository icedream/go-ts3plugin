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
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItemSlice[0]))
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItemSlice[1]))
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItem))
			}
			C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResult))
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
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItem))
			}
			C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResult))
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
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItemSlice[0]))
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItemSlice[1]))
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItem))
			}
			C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResult))
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
				C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItem))
			}
			C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResult))
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
			C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItem))
		}
		C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResult))
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
			C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResultItem))
		}
		C.freeMemory(this.nativeFunctions, unsafe.Pointer(cResult))
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

func (this *TS3Functions) OpenPlaybackDevice(serverConnectionHandlerID uint64, modeID, playbackDevice string) (retErrorCode uint32) {
	cModeID := C.CString(modeID)
	defer C.free(unsafe.Pointer(cModeID))

	cPlaybackDevice := C.CString(playbackDevice)
	defer C.free(unsafe.Pointer(cPlaybackDevice))

	retErrorCode = uint32(C.openPlaybackDevice(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		cModeID,
		cPlaybackDevice))
	return
}

func (this *TS3Functions) OpenCaptureDevice(serverConnectionHandlerID uint64, modeID, playbackDevice string) (retErrorCode uint32) {
	cModeID := C.CString(modeID)
	defer C.free(unsafe.Pointer(cModeID))

	cPlaybackDevice := C.CString(playbackDevice)
	defer C.free(unsafe.Pointer(cPlaybackDevice))

	retErrorCode = uint32(C.openCaptureDevice(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		cModeID,
		cPlaybackDevice))
	return
}

func (this *TS3Functions) GetCurrentPlaybackDeviceName(serverConnectionHandlerID uint64) (result string, isDefault bool, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	cIsDefault := C.int(0)

	retErrorCode = uint32(C.getCurrentPlaybackDeviceName(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		&cResult,
		&cIsDefault))

	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
		isDefault = cIsDefault != 0
	}

	return
}

func (this *TS3Functions) GetCurrentPlayBackMode(serverConnectionHandlerID uint64) (result string, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getCurrentPlayBackMode(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		&cResult))
	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) GetCurrentCaptureDeviceName(serverConnectionHandlerID uint64) (result string, isDefault bool, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	cIsDefault := C.int(0)

	retErrorCode = uint32(C.getCurrentCaptureDeviceName(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		&cResult,
		&cIsDefault))

	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
		isDefault = cIsDefault != 0
	}

	return
}

func (this *TS3Functions) GetCurrentCaptureMode(serverConnectionHandlerID uint64) (result string, retErrorCode uint32) {
	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getCurrentCaptureMode(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		&cResult))
	if retErrorCode == teamspeak.ErrorOK {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) InitiateGracefulPlaybackShutdown(serverConnectionHandlerID uint64) (retErrorCode uint32) {
	retErrorCode = uint32(C.initiateGracefulPlaybackShutdown(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID)))
	return
}

func (this *TS3Functions) ClosePlaybackDevice(serverConnectionHandlerID uint64) (retErrorCode uint32) {
	retErrorCode = uint32(C.closePlaybackDevice(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID)))
	return
}

func (this *TS3Functions) CloseCaptureDevice(serverConnectionHandlerID uint64) (retErrorCode uint32) {
	retErrorCode = uint32(C.closeCaptureDevice(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID)))
	return
}

func (this *TS3Functions) ActivateCaptureDevice(serverConnectionHandlerID uint64) (retErrorCode uint32) {
	retErrorCode = uint32(C.activateCaptureDevice(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID)))
	return
}

func (this *TS3Functions) RegisterCustomDevice(deviceID, deviceDisplayName string, capFrequency, capChannels, playFrequency, playChannels int) (retErrorCode uint32) {
	cDeviceID := C.CString(deviceID)
	defer C.free(unsafe.Pointer(cDeviceID))

	cDeviceDisplayName := C.CString(deviceDisplayName)
	defer C.free(unsafe.Pointer(cDeviceDisplayName))

	retErrorCode = uint32(C.registerCustomDevice(this.nativeFunctions,
		cDeviceID,
		cDeviceDisplayName,
		C.int(capFrequency),
		C.int(capChannels),
		C.int(playFrequency),
		C.int(playChannels)))
	return
}

func (this *TS3Functions) UnregisterCustomDevice(deviceID string) (retErrorCode uint32) {
	cDeviceID := C.CString(deviceID)
	defer C.free(unsafe.Pointer(cDeviceID))

	retErrorCode = uint32(C.unregisterCustomDevice(this.nativeFunctions,
		cDeviceID))
	return
}

func (this *TS3Functions) ProcessCustomCaptureData(deviceName string, sampleSlice []int16) (retErrorCode uint32) {
	cDeviceName := C.CString(deviceName)
	defer C.free(unsafe.Pointer(cDeviceName))

	cBufferSlice := make([]C.short, len(sampleSlice), len(sampleSlice))
	for i, sample := range sampleSlice {
		cBufferSlice[i] = C.short(sample)
	}

	cSamples := C.int(len(cBufferSlice))

	retErrorCode = uint32(C.processCustomCaptureData(this.nativeFunctions,
		cDeviceName, &cBufferSlice[0], cSamples))

	return
}

func (this *TS3Functions) AcquireCustomPlaybackData(deviceName string, sampleSlice []int16) (retErrorCode uint32) {
	cDeviceName := C.CString(deviceName)
	defer C.free(unsafe.Pointer(cDeviceName))

	cBufferSlice := make([]C.short, len(sampleSlice), len(sampleSlice))

	cSamples := C.int(len(cBufferSlice))

	retErrorCode = uint32(C.acquireCustomPlaybackData(this.nativeFunctions,
		cDeviceName, &cBufferSlice[0], cSamples))

	for i, sample := range cBufferSlice {
		sampleSlice[i] = int16(sample)
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

func (this *TS3Functions) GetPreProcessorInfoValueFloat(serverConnectionHandlerID uint64, ident string) (result float64, retErrorCode uint32) {
	identC := C.CString(ident)

	cResult := C.float(0)

	retErrorCode = uint32(C.getPreProcessorInfoValueFloat(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		identC,
		&cResult,
	))
	if retErrorCode == 0 {
		result = float64(cResult)
	}

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

func (this *TS3Functions) GetEncodeConfigValue(serverConnectionHandlerID uint64, ident string) (result string, retErrorCode uint32) {
	identC := C.CString(ident)

	// create buffer for teamspeak code to write to
	cResult := (*C.char)(C.malloc(2 * 1024))
	defer C.free(unsafe.Pointer(cResult))

	retErrorCode = uint32(C.getEncodeConfigValue(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		identC,
		&cResult,
	))
	if retErrorCode == 0 {
		result = C.GoString(cResult)
	}

	return
}

func (this *TS3Functions) GetPlaybackConfigValueAsFloat(serverConnectionHandlerID uint64, ident string) (result float64, retErrorCode uint32) {
	identC := C.CString(ident)

	cResult := C.float(0)

	retErrorCode = uint32(C.getPlaybackConfigValueAsFloat(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		identC,
		&cResult,
	))
	if retErrorCode == 0 {
		result = float64(cResult)
	}

	return
}

func (this *TS3Functions) SetPlaybackConfigValue(serverConnectionHandlerID uint64, ident string, value string) (retErrorCode uint32) {
	identC := C.CString(ident)
	valueC := C.CString(value)

	retErrorCode = uint32(C.setPlaybackConfigValue(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		identC,
		valueC,
	))

	return
}

// TODO - setClientVolumeModifier

func (this *TS3Functions) StartVoiceRecording(serverConnectionHandlerID uint64) (retErrorCode uint32) {
	retErrorCode = uint32(C.startVoiceRecording(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID)))

	return
}

func (this *TS3Functions) StopVoiceRecording(serverConnectionHandlerID uint64) (retErrorCode uint32) {
	retErrorCode = uint32(C.stopVoiceRecording(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID)))

	return
}

// TODO - systemset3DListenerAttributes
// TODO - set3DWaveAttributes
// TODO - systemset3DSettings
// TODO - channelset3DAttributes

func (this *TS3Functions) StartConnection(
	serverConnectionHandlerID uint64,
	identity string,
	ip string,
	port uint32,
	nickname string,
	defaultChannelArray []string,
	defaultChannelPassword string,
	serverPassword string,
) (retErrorCode uint32) {
	cIdentity := C.CString(identity)

	cIP := C.CString(ip)

	cNickname := C.CString(nickname)

	cDefaultChannelArraySlice := make([]*C.char, len(defaultChannelArray)+1)
	for i, value := range defaultChannelArray {
		cDefaultChannelArraySlice[i] = C.CString(value)
	}
	cDefaultChannelArray := &cDefaultChannelArraySlice[0]

	cDefaultChannelPassword := C.CString(defaultChannelPassword)

	cServerPassword := C.CString(serverPassword)

	retErrorCode = uint32(C.startConnection(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		cIdentity,
		cIP,
		C.uint(port),
		cNickname,
		cDefaultChannelArray,
		cDefaultChannelPassword,
		cServerPassword))

	return
}

func (this *TS3Functions) StopConnection(
	serverConnectionHandlerID uint64,
	quitMessage string,
) (retErrorCode uint32) {
	cQuitMessage := C.CString(quitMessage)

	retErrorCode = uint32(C.stopConnection(this.nativeFunctions,
		C.uint64(serverConnectionHandlerID),
		cQuitMessage))

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
