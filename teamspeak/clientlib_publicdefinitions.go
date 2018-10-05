package teamspeak

/*
#cgo CFLAGS: -I../pluginsdk/include -I.
#include "teamspeak/clientlib_publicdefinitions.h"
*/
import "C"

/******************************************************************************
From clientlib_publicdefinitions.h
*******************************************************************************/

type Visibility uint8

const (
	EnterVisibility  = Visibility(C.ENTER_VISIBILITY)
	RetainVisibility = Visibility(C.RETAIN_VISIBILITY)
	LeaveVisibility  = Visibility(C.LEAVE_VISIBILITY)
)

type ConnectStatus uint8

func (value ConnectStatus) toC() C.enum_ConnectStatus {
	return C.enum_ConnectStatus(value)
}

const (
	//There is no activity to the server, this is the default value
	StatusDisconnected = ConnectStatus(C.STATUS_DISCONNECTED)
	//We are trying to connect, we haven't got a clientID yet, we haven't been accepted by the server
	StatusConnecting = ConnectStatus(C.STATUS_CONNECTING)
	//The server has accepted us, we can talk and hear and we got a clientID, but we don't have the channels and clients yet, we can get server infos (welcome msg etc.)
	StatusConnected = ConnectStatus(C.STATUS_CONNECTED)
	//we are CONNECTED and we are visible
	StatusConnectionEstablishing = ConnectStatus(C.STATUS_CONNECTION_ESTABLISHING)
	//we are CONNECTED and we have the client and channels available
	StatusConnectionEstablished = ConnectStatus(C.STATUS_CONNECTION_ESTABLISHED)
)

type LocalTestMode uint8

const (
	TestModeOff                 = LocalTestMode(C.TEST_MODE_OFF)
	TestModeVoiceLocalOnly      = LocalTestMode(C.TEST_MODE_VOICE_LOCAL_ONLY)
	TestModeVoiceLocalAndRemote = LocalTestMode(C.TEST_MODE_VOICE_LOCAL_AND_REMOTE)
)
