package ts3plugin

/*
#cgo CFLAGS: -I${SRCDIR}/pluginsdk/include

#include <stdio.h> // uint types
#include <stdlib.h> // free

// #include "teamspeak/clientlib_publicdefinitions.h"
#include "teamspeak/public_definitions.h"
// #include "teamspeak/public_errors.h"
// #include "teamspeak/public_errors_rare.h"
// #include "teamspeak/public_rare_definitions.h"

#include "plugin_definitions.h"
#include "ts3_functions.h"

typedef struct TS3Functions TS3Functions;

*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/icedream/go-ts3plugin/teamlog"
	"github.com/icedream/go-ts3plugin/teamspeak"
)

const (
	logTimeLayout = "2006-01-02 15:04:05.999999"
)

type Samples struct {
	channels     int
	samplesSlice []C.short
	sampleCount  int
	edited       bool
}

func newSamples(samplesArrayPointer *C.short, sampleCount C.int, channels C.int) *Samples {
	// Convert samples C array pointer to Go slice
	length := int(sampleCount) * int(channels)
	samplesSlice := (*[1 << 28]C.short)(unsafe.Pointer(samplesArrayPointer))[:length:length]

	return &Samples{
		channels:     int(channels),
		sampleCount:  int(sampleCount),
		samplesSlice: samplesSlice,
	}
}

func (s *Samples) Channels() int {
	return s.channels
}

func (s *Samples) SampleCount() int {
	return s.sampleCount
}

func (s *Samples) IsEdited() bool {
	return s.edited
}

// GetSamples returns an array containing the requested amount of samples
// for all channels. The resulting slice will be the requested sample count
// multiplied by the channel count.
// The array is a copy of all the values. To store the samples value, use
// SetSamples.
func (s *Samples) GetSamples(offset int, sampleCount int) (samples []int16) {
	samples = make([]int16, sampleCount*s.channels)
	for i := 0; i < len(samples); i++ {
		samples[i] = int16(s.samplesSlice[i])
	}
	return
}

// SetSamples writes the passed samples array to the original memory at the given
// offset.
func (s *Samples) SetSamples(offset int, samples []int16) {
	if len(samples)%s.channels != 0 {
		panic("Samples array must have a sample for each channel but does not.")
	}
	for i := offset; i < offset+len(samples); i++ {
		s.samplesSlice[i] = C.short(samples[i-offset])
	}
	s.edited = true
}

func convertAnyIDToGo(value C.anyID) teamspeak.AnyID {
	return teamspeak.AnyID(uint16(value))
}

func notYetImplemented(name string) {
	Functions().LogMessage(
		fmt.Sprintf("NOT YET IMPLEMENTED: %s!", name),
		teamlog.Warning, Name, 0)
}

//export ts3plugin_name
func ts3plugin_name() *C.char {
	return C.CString(Name)
}

//export ts3plugin_version
func ts3plugin_version() *C.char {
	return C.CString(Version)
}

//export ts3plugin_apiVersion
func ts3plugin_apiVersion() C.int {
	return C.int(ApiVersion)
}

//export ts3plugin_author
func ts3plugin_author() *C.char {
	return C.CString(Author)
}

//export ts3plugin_description
func ts3plugin_description() *C.char {
	return C.CString(Description)
}

//export ts3plugin_setFunctionPointers
func ts3plugin_setFunctionPointers(funcs C.TS3Functions) {
	functions = &TS3Functions{
		nativeFunctions: funcs,
	}
}

//export ts3plugin_init
func ts3plugin_init() C.int {
	if Init == nil {
		return 0
	}

	if !Init() {
		return 1
	}

	return 0
}

//export ts3plugin_shutdown
func ts3plugin_shutdown() {
	if Shutdown == nil {
		return
	}

	Shutdown()
}

/*
Optional functions
*/

//export ts3plugin_freeMemory
func ts3plugin_freeMemory(data unsafe.Pointer) {
	C.free(data)
}

//export ts3plugin_offersConfigure
func ts3plugin_offersConfigure() C.int {
	if OffersConfigure == nil {
		return C.int(PLUGIN_OFFERS_NO_CONFIGURE)
	}
	return C.int(OffersConfigure())
}

//export ts3plugin_configure
func ts3plugin_configure(handle *C.char, qParentWidget *C.char) {
	notYetImplemented("ts3plugin_configure")
}

//export ts3plugin_registerPluginID
func ts3plugin_registerPluginID(id *C.char) {
	pluginID = C.GoString(id)
}

//export ts3plugin_commandKeyword
func ts3plugin_commandKeyword() *C.char {
	return C.CString(CommandKeyword)
}

//export ts3plugin_processCommand
func ts3plugin_processCommand(serverConnectionHandlerID C.uint64, command *C.char) C.int {
	if ProcessCommand(uint64(serverConnectionHandlerID), C.GoString(command)) {
		return 0
	}
	return 1
}

//export ts3plugin_currentServerConnectionChanged
func ts3plugin_currentServerConnectionChanged(serverConnectionHandlerID C.uint64) {
	if CurrentServerConnectionChanged != nil {
		CurrentServerConnectionChanged(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_infoTitle
func ts3plugin_infoTitle() *C.char {
	return C.CString(InfoTitle)
}

//export ts3plugin_infoData
// func ts3plugin_infoData(serverConnectionHandlerID C.uint64, id C.uint64, pluginItemType C.PluginItemType, data *C.char) {
// }

//export ts3plugin_requestAutoload
func ts3plugin_requestAutoload() C.int {
	if RequestAutoload {
		return 1
	}
	return 0
}

//export ts3plugin_initMenus
// func ts3plugin_initMenus(menuItems ***C.PluginMenuItem, menuIcon *C.char) {
// 	notYetImplemented("ts3plugin_initMenus")
// }

//export ts3plugin_initHotkeys
// func ts3plugin_initHotkeys(hotkeys ***C.PluginHotkey) {
// 	notYetImplemented("ts3plugin_initHotkeys")
// }

/* Clientlib */
//export ts3plugin_onConnectStatusChangeEvent
func ts3plugin_onConnectStatusChangeEvent(serverConnectionHandlerID C.uint64, newStatus C.int, errorNumber C.uint) {
	if OnConnectedStatusChangeEvent != nil {
		OnConnectedStatusChangeEvent(
			uint64(serverConnectionHandlerID),
			int(newStatus),
			uint(errorNumber),
		)
	}
}

//export ts3plugin_onNewChannelEvent
func ts3plugin_onNewChannelEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, channelParentID C.uint64) {
	if OnNewChannelEvent != nil {
		OnNewChannelEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
			uint64(channelParentID),
		)
	}
}

//export ts3plugin_onNewChannelCreatedEvent
func ts3plugin_onNewChannelCreatedEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, channelParentID C.uint64, invokerID C.anyID, invokerName *C.char, invokerUniqueIdentifier *C.char) {
	if OnNewChannelCreatedEvent != nil {
		OnNewChannelCreatedEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
			uint64(channelParentID),
			convertAnyIDToGo(invokerID),
			C.GoString(invokerName),
			C.GoString(invokerUniqueIdentifier),
		)
	}
}

//export ts3plugin_onDelChannelEvent
func ts3plugin_onDelChannelEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, invokerID C.anyID, invokerName *C.char, invokerUniqueIdentifier *C.char) {
	if OnDelChannelEvent != nil {
		OnDelChannelEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
			convertAnyIDToGo(invokerID),
			C.GoString(invokerName),
			C.GoString(invokerUniqueIdentifier),
		)
	}
}

//export ts3plugin_onChannelMoveEvent
func ts3plugin_onChannelMoveEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, newChannelParentID C.uint64, invokerID C.anyID, invokerName *C.char, invokerUniqueIdentifier *C.char) {
	if OnChannelMoveEvent != nil {
		OnChannelMoveEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
			uint64(newChannelParentID),
			convertAnyIDToGo(invokerID),
			C.GoString(invokerName),
			C.GoString(invokerUniqueIdentifier),
		)
	}
}

//export ts3plugin_onUpdateChannelEvent
func ts3plugin_onUpdateChannelEvent(serverConnectionHandlerID C.uint64, channelID C.uint64) {
	if OnUpdateChannelEvent != nil {
		OnUpdateChannelEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
		)
	}
}

//export ts3plugin_onUpdateChannelEditedEvent
func ts3plugin_onUpdateChannelEditedEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, invokerID C.anyID, invokerName *C.char, invokerUniqueIdentifier *C.char) {
	if OnUpdateChannelEditedEvent != nil {
		OnUpdateChannelEditedEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
			convertAnyIDToGo(invokerID),
			C.GoString(invokerName),
			C.GoString(invokerUniqueIdentifier),
		)
	}
}

//export ts3plugin_onUpdateClientEvent
func ts3plugin_onUpdateClientEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, invokerID C.anyID, invokerName *C.char, invokerUniqueIdentifier *C.char) {
	if OnUpdateClientEvent != nil {
		OnUpdateClientEvent(
			uint64(serverConnectionHandlerID),
			convertAnyIDToGo(clientID),
			convertAnyIDToGo(invokerID),
			C.GoString(invokerName),
			C.GoString(invokerUniqueIdentifier),
		)
	}
}

//export ts3plugin_onClientMoveEvent
func ts3plugin_onClientMoveEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int, moveMessage *C.char) {
	if OnClientMoveEvent != nil {
		OnClientMoveEvent(
			uint64(serverConnectionHandlerID),
			convertAnyIDToGo(clientID),
			uint64(oldChannelID),
			uint64(newChannelID),
			teamspeak.Visibility(int(visibility)),
			C.GoString(moveMessage),
		)
	}
}

//export ts3plugin_onClientMoveSubscriptionEvent
func ts3plugin_onClientMoveSubscriptionEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int) {
	if OnClientMoveSubscriptionEvent != nil {
		OnClientMoveSubscriptionEvent(
			uint64(serverConnectionHandlerID),
			convertAnyIDToGo(clientID),
			uint64(oldChannelID),
			uint64(newChannelID),
			teamspeak.Visibility(int(visibility)),
		)
	}
}

//export ts3plugin_onClientMoveTimeoutEvent
func ts3plugin_onClientMoveTimeoutEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int, timeoutMessage *C.char) {
	notYetImplemented("ts3plugin_onClientMoveTimeoutEvent")
}

//export ts3plugin_onClientMoveMovedEvent
func ts3plugin_onClientMoveMovedEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int, moverID C.anyID, moverName *C.char, moverUniqueIdentifier *C.char, moveMessage *C.char) {
	notYetImplemented("ts3plugin_onClientMoveMovedEvent")
}

//export ts3plugin_onClientKickFromChannelEvent
func ts3plugin_onClientKickFromChannelEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int, kickerID C.anyID, kickerName *C.char, kickerUniqueIdentifier *C.char, kickMessage *C.char) {
	notYetImplemented("ts3plugin_onClientKickFromChannelEvent")
}

//export ts3plugin_onClientKickFromServerEvent
func ts3plugin_onClientKickFromServerEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int, kickerID C.anyID, kickerName *C.char, kickerUniqueIdentifier *C.char, kickMessage *C.char) {
	notYetImplemented("ts3plugin_onClientKickFromServerEvent")
}

//export ts3plugin_onClientIDsEvent
func ts3plugin_onClientIDsEvent(serverConnectionHandlerID C.uint64, uniqueClientIdentifier *C.char, clientID C.anyID, clientName *C.char) {
	notYetImplemented("ts3plugin_onClientIDsEvent")
}

//export ts3plugin_onClientIDsFinishedEvent
func ts3plugin_onClientIDsFinishedEvent(serverConnectionHandlerID C.uint64) {
	if OnClientIDsFinishedEvent != nil {
		OnClientIDsFinishedEvent(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_onServerEditedEvent
func ts3plugin_onServerEditedEvent(serverConnectionHandlerID C.uint64, editerID C.anyID, editerName *C.char, editerUniqueIdentifier *C.char) {
	if OnServerEditedEvent != nil {
		OnServerEditedEvent(
			uint64(serverConnectionHandlerID),
			convertAnyIDToGo(editerID),
			C.GoString(editerName),
			C.GoString(editerUniqueIdentifier),
		)
	}
}

//export ts3plugin_onServerUpdatedEvent
func ts3plugin_onServerUpdatedEvent(serverConnectionHandlerID C.uint64) {
	if OnServerUpdatedEvent != nil {
		OnServerUpdatedEvent(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_onServerErrorEvent
func ts3plugin_onServerErrorEvent(serverConnectionHandlerID C.uint64, errorMessage *C.char, errorCode C.uint, returnCode *C.char, extraMessage *C.char) C.int {
	if OnServerErrorEvent != nil {
		return C.int(OnServerErrorEvent(
			uint64(serverConnectionHandlerID),
			C.GoString(errorMessage),
			uint(errorCode),
			C.GoString(returnCode),
			C.GoString(extraMessage),
		))
	}
	return 0
}

//export ts3plugin_onServerStopEvent
func ts3plugin_onServerStopEvent(serverConnectionHandlerID C.uint64, shutdownMessage *C.char) {
	if OnServerStopEvent != nil {
		OnServerStopEvent(
			uint64(serverConnectionHandlerID),
			C.GoString(shutdownMessage),
		)
	}
}

//export ts3plugin_onTextMessageEvent
func ts3plugin_onTextMessageEvent(serverConnectionHandlerID C.uint64, targetMode C.anyID, toID C.anyID, fromID C.anyID, fromName *C.char, fromUniqueIdentifier *C.char, message *C.char, ffIgnored C.int) C.int {
	if OnTextMessageEvent != nil {
		return C.int(OnTextMessageEvent(
			uint64(serverConnectionHandlerID),
			convertAnyIDToGo(targetMode),
			convertAnyIDToGo(toID),
			convertAnyIDToGo(fromID),
			C.GoString(fromName),
			C.GoString(fromUniqueIdentifier),
			C.GoString(message),
			int(ffIgnored) != 0,
		))
	}
	return 0
}

//export ts3plugin_onTalkStatusChangeEvent
func ts3plugin_onTalkStatusChangeEvent(serverConnectionHandlerID C.uint64, status C.int, isReceivedWhisper C.int, clientID C.anyID) {
	if OnTalkStatusChangeEvent != nil {
		OnTalkStatusChangeEvent(
			uint64(serverConnectionHandlerID),
			int(status),
			int(isReceivedWhisper),
			convertAnyIDToGo(clientID),
		)
	}
}

//export ts3plugin_onConnectionInfoEvent
func ts3plugin_onConnectionInfoEvent(serverConnectionHandlerID C.uint64, clientID C.anyID) {
	if OnConnectionInfoEvent != nil {
		OnConnectionInfoEvent(
			uint64(serverConnectionHandlerID),
			convertAnyIDToGo(clientID),
		)
	}
}

//export ts3plugin_onServerConnectionInfoEvent
func ts3plugin_onServerConnectionInfoEvent(serverConnectionHandlerID C.uint64) {
	if OnServerConnectionInfoEvent != nil {
		OnServerConnectionInfoEvent(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_onChannelSubscribeEvent
func ts3plugin_onChannelSubscribeEvent(serverConnectionHandlerID C.uint64, channelID C.uint64) {
	if OnChannelSubscribeEvent != nil {
		OnChannelSubscribeEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
		)
	}
}

//export ts3plugin_onChannelSubscribeFinishedEvent
func ts3plugin_onChannelSubscribeFinishedEvent(serverConnectionHandlerID C.uint64) {
	if OnChannelSubscribeFinishedEvent != nil {
		OnChannelSubscribeFinishedEvent(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_onChannelUnsubscribeEvent
func ts3plugin_onChannelUnsubscribeEvent(serverConnectionHandlerID C.uint64, channelID C.uint64) {
	if OnChannelUnsubscribeEvent != nil {
		OnChannelUnsubscribeEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
		)
	}
}

//export ts3plugin_onChannelUnsubscribeFinishedEvent
func ts3plugin_onChannelUnsubscribeFinishedEvent(serverConnectionHandlerID C.uint64) {
	if OnChannelUnsubscribeFinishedEvent != nil {
		OnChannelUnsubscribeFinishedEvent(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_onChannelDescriptionUpdateEvent
func ts3plugin_onChannelDescriptionUpdateEvent(serverConnectionHandlerID C.uint64, channelID C.uint64) {
	if OnChannelDescriptionUpdateEvent != nil {
		OnChannelDescriptionUpdateEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
		)
	}
}

//export ts3plugin_onChannelPasswordChangedEvent
func ts3plugin_onChannelPasswordChangedEvent(serverConnectionHandlerID C.uint64, channelID C.uint64) {
	if OnChannelPasswordChangedEvent != nil {
		OnChannelPasswordChangedEvent(
			uint64(serverConnectionHandlerID),
			uint64(channelID),
		)
	}
}

//export ts3plugin_onPlaybackShutdownCompleteEvent
func ts3plugin_onPlaybackShutdownCompleteEvent(serverConnectionHandlerID C.uint64) {
	if OnPlaybackShutdownCompleteEvent != nil {
		OnPlaybackShutdownCompleteEvent(
			uint64(serverConnectionHandlerID),
		)
	}
}

//export ts3plugin_onSoundDeviceListChangedEvent
func ts3plugin_onSoundDeviceListChangedEvent(modeID *C.char, playOrCap C.int) {
	notYetImplemented("ts3plugin_onSoundDeviceListChangedEvent")
}

//export ts3plugin_onEditPlaybackVoiceDataEvent
func ts3plugin_onEditPlaybackVoiceDataEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, samples *C.short, sampleCount C.int, channels C.int) {
	notYetImplemented("ts3plugin_onEditPlaybackVoiceDataEvent")
}

//export ts3plugin_onEditPostProcessVoiceDataEvent
func ts3plugin_onEditPostProcessVoiceDataEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, samples *C.short, sampleCount C.int, channels C.int, channelSpeakerArray *C.uint, channelFillMask *C.uint) {
	notYetImplemented("ts3plugin_onEditPostProcessVoiceDataEvent")
}

//export ts3plugin_onEditMixedPlaybackVoiceDataEvent
func ts3plugin_onEditMixedPlaybackVoiceDataEvent(serverConnectionHandlerID C.uint64, samples *C.short, sampleCount C.int, channels C.int, channelSpeakerArray *C.uint, channelFillMask *C.uint) {
	if OnEditMixedPlaybackVoiceDataEvent != nil {
		channelsGo := int(channels)
		channelSpeakerSlice := make([]uint, channelsGo)
		for i, value := range (*[1 << 28]C.uint)(unsafe.Pointer(channelSpeakerArray))[:channelsGo:channelsGo] {
			channelSpeakerSlice[i] = uint(value)
		}

		channelFillMaskGo := uint(*channelFillMask)

		samplesGo := newSamples(samples, sampleCount, channels)

		OnEditMixedPlaybackVoiceDataEvent(
			uint64(serverConnectionHandlerID),
			samplesGo,
			channelSpeakerSlice,
			&channelFillMaskGo,
		)

		*channelFillMask = C.uint(channelFillMaskGo)
	}
}

//export ts3plugin_onEditCapturedVoiceDataEvent
func ts3plugin_onEditCapturedVoiceDataEvent(serverConnectionHandlerID C.uint64, samples *C.short, sampleCount C.int, channels C.int, edited *C.uint) {
	if OnEditCapturedVoiceDataEvent != nil {
		samplesGo := newSamples(samples, sampleCount, channels)

		shouldMute := OnEditCapturedVoiceDataEvent(
			uint64(serverConnectionHandlerID),
			samplesGo,
			(*edited&2) == 0, // check if bit 2 is unset (whether audio is muted)
		)

		if shouldMute {
			*edited &= ^C.uint(2) // clear bit 2
		} else {
			*edited |= 2 // set bit 2
		}

		if samplesGo.IsEdited() {
			*edited |= 1 // set bit 1
		}
	}
}

//export ts3plugin_onCustom3dRolloffCalculationClientEvent
func ts3plugin_onCustom3dRolloffCalculationClientEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, distance C.float, volume *C.float) {
	notYetImplemented("ts3plugin_onCustom3dRolloffCalculationClientEvent")
}

//export ts3plugin_onCustom3dRolloffCalculationWaveEvent
func ts3plugin_onCustom3dRolloffCalculationWaveEvent(serverConnectionHandlerID C.uint64, waveHandle C.uint64, distance C.float, volume *C.float) {
	notYetImplemented("ts3plugin_onCustom3dRolloffCalculationWaveEvent")
}

//export ts3plugin_onUserLoggingMessageEvent
func ts3plugin_onUserLoggingMessageEvent(logMessage *C.char, logLevel C.int, logChannel *C.char, logID C.uint64, logTime *C.char, completeLogString *C.char) {
	logTimeParsed, _ := time.Parse(logTimeLayout, C.GoString(logTime))
	if OnUserLoggingMessageEvent != nil {
		OnUserLoggingMessageEvent(
			C.GoString(logMessage),
			teamlog.LogLevel(int(logLevel)),
			C.GoString(logChannel),
			uint64(logID),
			logTimeParsed,
			C.GoString(completeLogString),
		)
	}
}

/* Clientlib rare */
//export ts3plugin_onClientBanFromServerEvent
func ts3plugin_onClientBanFromServerEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, oldChannelID C.uint64, newChannelID C.uint64, visibility C.int, kickerID C.anyID, kickerName *C.char, kickerUniqueIdentifier *C.char, time C.uint64, kickMessage *C.char) {
	notYetImplemented("ts3plugin_onClientBanFromServerEvent")
}

//export  ts3plugin_onClientPokeEvent
func ts3plugin_onClientPokeEvent(serverConnectionHandlerID C.uint64, fromClientID C.anyID, pokerName *C.char, pokerUniqueIdentity *C.char, message *C.char, ffIgnored C.int) C.int {
	notYetImplemented("ts3plugin_onClientPokeEvent")
	return 0
}

//export ts3plugin_onClientSelfVariableUpdateEvent
func ts3plugin_onClientSelfVariableUpdateEvent(serverConnectionHandlerID C.uint64, flag C.int, oldValue *C.char, newValue *C.char) {
	notYetImplemented("ts3plugin_onClientSelfVariableUpdateEvent")
}

//export ts3plugin_onFileListEvent
func ts3plugin_onFileListEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, path *C.char, name *C.char, size C.uint64, datetime C.uint64, typeID C.int, incompletesize C.uint64, returnCode *C.char) {
	notYetImplemented("ts3plugin_onFileListEvent")
}

//export ts3plugin_onFileListFinishedEvent
func ts3plugin_onFileListFinishedEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, path *C.char) {
	notYetImplemented("ts3plugin_onFileListFinishedEvent")
}

//export ts3plugin_onFileInfoEvent
func ts3plugin_onFileInfoEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, name *C.char, size C.uint64, datetime C.uint64) {
	notYetImplemented("ts3plugin_onFileInfoEvent")
}

//export ts3plugin_onServerGroupListEvent
func ts3plugin_onServerGroupListEvent(serverConnectionHandlerID C.uint64, serverGroupID C.uint64, name *C.char, typeID C.int, iconID C.int, saveDB C.int) {
	notYetImplemented("ts3plugin_onServerGroupListEvent")
}

//export ts3plugin_onServerGroupListFinishedEvent
func ts3plugin_onServerGroupListFinishedEvent(serverConnectionHandlerID C.uint64) {
	notYetImplemented("ts3plugin_onServerGroupListFinishedEvent")
}

//export ts3plugin_onServerGroupByClientIDEvent
func ts3plugin_onServerGroupByClientIDEvent(serverConnectionHandlerID C.uint64, name *C.char, serverGroupList C.uint64, clientDatabaseID C.uint64) {
	notYetImplemented("ts3plugin_onServerGroupByClientIDEvent")
}

//export ts3plugin_onServerGroupPermListEvent
func ts3plugin_onServerGroupPermListEvent(serverConnectionHandlerID C.uint64, serverGroupID C.uint64, permissionID C.uint, permissionValue C.int, permissionNegated C.int, permissionSkip C.int) {
	notYetImplemented("ts3plugin_onServerGroupPermListEvent")
}

//export ts3plugin_onServerGroupPermListFinishedEvent
func ts3plugin_onServerGroupPermListFinishedEvent(serverConnectionHandlerID C.uint64, serverGroupID C.uint64) {
	notYetImplemented("ts3plugin_onServerGroupPermListFinishedEvent")
}

//export ts3plugin_onServerGroupClientListEvent
func ts3plugin_onServerGroupClientListEvent(serverConnectionHandlerID C.uint64, serverGroupID C.uint64, clientDatabaseID C.uint64, clientNameIdentifier *C.char, clientUniqueID *C.char) {
	notYetImplemented("ts3plugin_onServerGroupClientListEvent")
}

//export ts3plugin_onChannelGroupListEvent
func ts3plugin_onChannelGroupListEvent(serverConnectionHandlerID C.uint64, channelGroupID C.uint64, name *C.char, typeID C.int, iconID C.int, saveDB C.int) {
	notYetImplemented("ts3plugin_onChannelGroupListEvent")
}

//export ts3plugin_onChannelGroupListFinishedEvent
func ts3plugin_onChannelGroupListFinishedEvent(serverConnectionHandlerID C.uint64) {
	notYetImplemented("ts3plugin_onChannelGroupListFinishedEvent")
}

//export ts3plugin_onChannelGroupPermListEvent
func ts3plugin_onChannelGroupPermListEvent(serverConnectionHandlerID C.uint64, channelGroupID C.uint64, permissionID C.uint, permissionValue C.int, permissionNegated C.int, permissionSkip C.int) {
	notYetImplemented("ts3plugin_onChannelGroupPermListEvent")
}

//export ts3plugin_onChannelGroupPermListFinishedEvent
func ts3plugin_onChannelGroupPermListFinishedEvent(serverConnectionHandlerID C.uint64, channelGroupID C.uint64) {
	notYetImplemented("ts3plugin_onChannelGroupPermListFinishedEvent")
}

//export ts3plugin_onChannelPermListEvent
func ts3plugin_onChannelPermListEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, permissionID C.uint, permissionValue C.int, permissionNegated C.int, permissionSkip C.int) {
	notYetImplemented("ts3plugin_onChannelPermListEvent")
}

//export ts3plugin_onChannelPermListFinishedEvent
func ts3plugin_onChannelPermListFinishedEvent(serverConnectionHandlerID C.uint64, channelID C.uint64) {
	notYetImplemented("ts3plugin_onChannelPermListFinishedEvent")
}

//export ts3plugin_onClientPermListEvent
func ts3plugin_onClientPermListEvent(serverConnectionHandlerID C.uint64, clientDatabaseID C.uint64, permissionID C.uint, permissionValue C.int, permissionNegated C.int, permissionSkip C.int) {
	notYetImplemented("ts3plugin_onClientPermListEvent")
}

//export ts3plugin_onClientPermListFinishedEvent
func ts3plugin_onClientPermListFinishedEvent(serverConnectionHandlerID C.uint64, clientDatabaseID C.uint64) {
	notYetImplemented("ts3plugin_onClientPermListFinishedEvent")
}

//export ts3plugin_onChannelClientPermListEvent
func ts3plugin_onChannelClientPermListEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, clientDatabaseID C.uint64, permissionID C.uint, permissionValue C.int, permissionNegated C.int, permissionSkip C.int) {
	notYetImplemented("ts3plugin_onChannelClientPermListEvent")
}

//export ts3plugin_onChannelClientPermListFinishedEvent
func ts3plugin_onChannelClientPermListFinishedEvent(serverConnectionHandlerID C.uint64, channelID C.uint64, clientDatabaseID C.uint64) {
	notYetImplemented("ts3plugin_onChannelClientPermListFinishedEvent")
}

//export ts3plugin_onClientChannelGroupChangedEvent
func ts3plugin_onClientChannelGroupChangedEvent(serverConnectionHandlerID C.uint64, channelGroupID C.uint64, channelID C.uint64, clientID C.anyID, invokerClientID C.anyID, invokerName *C.char, invokerUniqueIdentity *C.char) {
	notYetImplemented("ts3plugin_onClientChannelGroupChangedEvent")
}

//export  ts3plugin_onServerPermissionErrorEvent
func ts3plugin_onServerPermissionErrorEvent(serverConnectionHandlerID C.uint64, errorMessage *C.char, error C.uint, returnCode *C.char, failedPermissionID C.uint) C.int {
	notYetImplemented("ts3plugin_onServerPermissionErrorEvent")
	return 0
}

//export ts3plugin_onPermissionListGroupEndIDEvent
func ts3plugin_onPermissionListGroupEndIDEvent(serverConnectionHandlerID C.uint64, groupEndID C.uint) {
	notYetImplemented("ts3plugin_onPermissionListGroupEndIDEvent")
}

//export ts3plugin_onPermissionListEvent
func ts3plugin_onPermissionListEvent(serverConnectionHandlerID C.uint64, permissionID C.uint, permissionName *C.char, permissionDescription *C.char) {
	notYetImplemented("ts3plugin_onPermissionListEvent")
}

//export ts3plugin_onPermissionListFinishedEvent
func ts3plugin_onPermissionListFinishedEvent(serverConnectionHandlerID C.uint64) {
	notYetImplemented("ts3plugin_onPermissionListFinishedEvent")
}

//export ts3plugin_onPermissionOverviewEvent
func ts3plugin_onPermissionOverviewEvent(serverConnectionHandlerID C.uint64, clientDatabaseID C.uint64, channelID C.uint64, overviewType C.int, overviewID1 C.uint64, overviewID2 C.uint64, permissionID C.uint, permissionValue C.int, permissionNegated C.int, permissionSkip C.int) {
	notYetImplemented("ts3plugin_onPermissionOverviewEvent")
}

//export ts3plugin_onPermissionOverviewFinishedEvent
func ts3plugin_onPermissionOverviewFinishedEvent(serverConnectionHandlerID C.uint64) {
	notYetImplemented("ts3plugin_onPermissionOverviewFinishedEvent")
}

//export ts3plugin_onServerGroupClientAddedEvent
func ts3plugin_onServerGroupClientAddedEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, clientName *C.char, clientUniqueIdentity *C.char, serverGroupID C.uint64, invokerClientID C.anyID, invokerName *C.char, invokerUniqueIdentity *C.char) {
	notYetImplemented("ts3plugin_onServerGroupClientAddedEvent")
}

//export ts3plugin_onServerGroupClientDeletedEvent
func ts3plugin_onServerGroupClientDeletedEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, clientName *C.char, clientUniqueIdentity *C.char, serverGroupID C.uint64, invokerClientID C.anyID, invokerName *C.char, invokerUniqueIdentity *C.char) {
	notYetImplemented("ts3plugin_onServerGroupClientDeletedEvent")
}

//export ts3plugin_onClientNeededPermissionsEvent
func ts3plugin_onClientNeededPermissionsEvent(serverConnectionHandlerID C.uint64, permissionID C.uint, permissionValue C.int) {
	notYetImplemented("ts3plugin_onClientNeededPermissionsEvent")
}

//export ts3plugin_onClientNeededPermissionsFinishedEvent
func ts3plugin_onClientNeededPermissionsFinishedEvent(serverConnectionHandlerID C.uint64) {
	notYetImplemented("ts3plugin_onClientNeededPermissionsFinishedEvent")
}

//export ts3plugin_onFileTransferStatusEvent
func ts3plugin_onFileTransferStatusEvent(transferID C.anyID, status C.uint, statusMessage *C.char, remotefileSize C.uint64, serverConnectionHandlerID C.uint64) {
	notYetImplemented("ts3plugin_onFileTransferStatusEvent")
}

//export ts3plugin_onClientChatClosedEvent
func ts3plugin_onClientChatClosedEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, clientUniqueIdentity *C.char) {
	notYetImplemented("ts3plugin_onClientChatClosedEvent")
}

//export ts3plugin_onClientChatComposingEvent
func ts3plugin_onClientChatComposingEvent(serverConnectionHandlerID C.uint64, clientID C.anyID, clientUniqueIdentity *C.char) {
	notYetImplemented("ts3plugin_onClientChatComposingEvent")
}

//export ts3plugin_onServerLogEvent
func ts3plugin_onServerLogEvent(serverConnectionHandlerID C.uint64, logMsg *C.char) {
	notYetImplemented("ts3plugin_onServerLogEvent")
}

//export ts3plugin_onServerLogFinishedEvent
func ts3plugin_onServerLogFinishedEvent(serverConnectionHandlerID C.uint64, lastPos C.uint64, fileSize C.uint64) {
	notYetImplemented("ts3plugin_onServerLogFinishedEvent")
}

//export ts3plugin_onMessageListEvent
func ts3plugin_onMessageListEvent(serverConnectionHandlerID C.uint64, messageID C.uint64, fromClientUniqueIdentity *C.char, subject *C.char, timestamp C.uint64, flagRead C.int) {
	notYetImplemented("ts3plugin_onMessageListEvent")
}

//export ts3plugin_onMessageGetEvent
func ts3plugin_onMessageGetEvent(serverConnectionHandlerID C.uint64, messageID C.uint64, fromClientUniqueIdentity *C.char, subject *C.char, message *C.char, timestamp C.uint64) {
	notYetImplemented("ts3plugin_onMessageGetEvent")
}

//export ts3plugin_onClientDBIDfromUIDEvent
func ts3plugin_onClientDBIDfromUIDEvent(serverConnectionHandlerID C.uint64, uniqueClientIdentifier *C.char, clientDatabaseID C.uint64) {
	notYetImplemented("ts3plugin_onClientDBIDfromUIDEvent")
}

//export ts3plugin_onClientNamefromUIDEvent
func ts3plugin_onClientNamefromUIDEvent(serverConnectionHandlerID C.uint64, uniqueClientIdentifier *C.char, clientDatabaseID C.uint64, clientNickName *C.char) {
	notYetImplemented("ts3plugin_onClientNamefromUIDEvent")
}

//export ts3plugin_onClientNamefromDBIDEvent
func ts3plugin_onClientNamefromDBIDEvent(serverConnectionHandlerID C.uint64, uniqueClientIdentifier *C.char, clientDatabaseID C.uint64, clientNickName *C.char) {
	notYetImplemented("ts3plugin_onClientNamefromDBIDEvent")
}

//export ts3plugin_onComplainListEvent
func ts3plugin_onComplainListEvent(serverConnectionHandlerID C.uint64, targetClientDatabaseID C.uint64, targetClientNickName *C.char, fromClientDatabaseID C.uint64, fromClientNickName *C.char, complainReason *C.char, timestamp C.uint64) {
	notYetImplemented("ts3plugin_onComplainListEvent")
}

//export ts3plugin_onBanListEvent
func ts3plugin_onBanListEvent(serverConnectionHandlerID C.uint64, banid C.uint64, ip *C.char, name *C.char, uid *C.char, creationTime C.uint64, durationTime C.uint64, invokerName *C.char, invokercldbid C.uint64, invokeruid *C.char, reason *C.char, numberOfEnforcements C.int, lastNickName *C.char) {
	notYetImplemented("ts3plugin_onBanListEvent")
}

//export ts3plugin_onClientServerQueryLoginPasswordEvent
func ts3plugin_onClientServerQueryLoginPasswordEvent(serverConnectionHandlerID C.uint64, loginPassword *C.char) {
	notYetImplemented("ts3plugin_onClientServerQueryLoginPasswordEvent")
}

//export ts3plugin_onPluginCommandEvent
func ts3plugin_onPluginCommandEvent(serverConnectionHandlerID C.uint64, pluginName *C.char, pluginCommand *C.char) {
	notYetImplemented("ts3plugin_onPluginCommandEvent")
}

//export ts3plugin_onIncomingClientQueryEvent
func ts3plugin_onIncomingClientQueryEvent(serverConnectionHandlerID C.uint64, commandText *C.char) {
	notYetImplemented("ts3plugin_onIncomingClientQueryEvent")
}

//export ts3plugin_onServerTemporaryPasswordListEvent
func ts3plugin_onServerTemporaryPasswordListEvent(serverConnectionHandlerID C.uint64, clientNickname *C.char, uniqueClientIdentifier *C.char, description *C.char, password *C.char, timestampStart C.uint64, timestampEnd C.uint64, targetChannelID C.uint64, targetChannelPW *C.char) {
	notYetImplemented("ts3plugin_onServerTemporaryPasswordListEvent")
}

/* Client UI callbacks */
//export ts3plugin_onAvatarUpdated
func ts3plugin_onAvatarUpdated(serverConnectionHandlerID C.uint64, clientID C.anyID, avatarPath *C.char) {
	notYetImplemented("ts3plugin_onAvatarUpdated")
}

//export ts3plugin_onMenuItemEvent
// func ts3plugin_onMenuItemEvent(serverConnectionHandlerID C.uint64, typeID C.PluginMenuType, menuItemID C.int, selectedItemID C.uint64) {
// }

//export ts3plugin_onHotkeyEvent
func ts3plugin_onHotkeyEvent(keyword *C.char) {
	notYetImplemented("ts3plugin_onHotkeyEvent")
}

//export ts3plugin_onHotkeyRecordedEvent
func ts3plugin_onHotkeyRecordedEvent(keyword *C.char, key *C.char) {
	notYetImplemented("ts3plugin_onHotkeyRecordedEvent")
}

//export ts3plugin_onClientDisplayNameChanged
func ts3plugin_onClientDisplayNameChanged(serverConnectionHandlerID C.uint64, clientID C.anyID, displayName *C.char, uniqueClientIdentifier *C.char) {
	notYetImplemented("ts3plugin_onClientDisplayNameChanged")
}
