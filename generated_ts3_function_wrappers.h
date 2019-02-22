/**
 * This file has been automatically generated.
 * Please use 'go generate' to update this file.
 */

#include "ts3_functions.h"

typedef struct TS3Functions TS3Functions;


// unsigned int getClientLibVersion(char**  result)
unsigned int getClientLibVersion(const struct TS3Functions funcs, char**  result) {
	
	return funcs.getClientLibVersion(result);
}

// unsigned int getClientLibVersionNumber(uint64*  result)
unsigned int getClientLibVersionNumber(const struct TS3Functions funcs, uint64*  result) {
	
	return funcs.getClientLibVersionNumber(result);
}

// unsigned int spawnNewServerConnectionHandler(int  port, uint64*  result)
unsigned int spawnNewServerConnectionHandler(const struct TS3Functions funcs, int  port, uint64*  result) {
	
	return funcs.spawnNewServerConnectionHandler(port, result);
}

// unsigned int destroyServerConnectionHandler(uint64  serverConnectionHandlerID)
unsigned int destroyServerConnectionHandler(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.destroyServerConnectionHandler(serverConnectionHandlerID);
}

// unsigned int getErrorMessage(unsigned int  errorCode, char**  error)
unsigned int getErrorMessage(const struct TS3Functions funcs, unsigned int  errorCode, char**  error) {
	
	return funcs.getErrorMessage(errorCode, error);
}

// unsigned int freeMemory(void*  pointer)
unsigned int freeMemory(const struct TS3Functions funcs, void*  pointer) {
	
	return funcs.freeMemory(pointer);
}

// unsigned int logMessage(const char*  logMessage, enum LogLevel  severity, const char*  channel, uint64  logID)
unsigned int logMessage(const struct TS3Functions funcs, const char*  logMessage, enum LogLevel  severity, const char*  channel, uint64  logID) {
	
	return funcs.logMessage(logMessage, severity, channel, logID);
}

// unsigned int getPlaybackDeviceList(const char*  modeID, char****  result)
unsigned int getPlaybackDeviceList(const struct TS3Functions funcs, const char*  modeID, char****  result) {
	
	return funcs.getPlaybackDeviceList(modeID, result);
}

// unsigned int getPlaybackModeList(char***  result)
unsigned int getPlaybackModeList(const struct TS3Functions funcs, char***  result) {
	
	return funcs.getPlaybackModeList(result);
}

// unsigned int getCaptureDeviceList(const char*  modeID, char****  result)
unsigned int getCaptureDeviceList(const struct TS3Functions funcs, const char*  modeID, char****  result) {
	
	return funcs.getCaptureDeviceList(modeID, result);
}

// unsigned int getCaptureModeList(char***  result)
unsigned int getCaptureModeList(const struct TS3Functions funcs, char***  result) {
	
	return funcs.getCaptureModeList(result);
}

// unsigned int getDefaultPlaybackDevice(const char*  modeID, char***  result)
unsigned int getDefaultPlaybackDevice(const struct TS3Functions funcs, const char*  modeID, char***  result) {
	
	return funcs.getDefaultPlaybackDevice(modeID, result);
}

// unsigned int getDefaultPlayBackMode(char**  result)
unsigned int getDefaultPlayBackMode(const struct TS3Functions funcs, char**  result) {
	
	return funcs.getDefaultPlayBackMode(result);
}

// unsigned int getDefaultCaptureDevice(const char*  modeID, char***  result)
unsigned int getDefaultCaptureDevice(const struct TS3Functions funcs, const char*  modeID, char***  result) {
	
	return funcs.getDefaultCaptureDevice(modeID, result);
}

// unsigned int getDefaultCaptureMode(char**  result)
unsigned int getDefaultCaptureMode(const struct TS3Functions funcs, char**  result) {
	
	return funcs.getDefaultCaptureMode(result);
}

// unsigned int openPlaybackDevice(uint64  serverConnectionHandlerID, const char*  modeID, const char*  playbackDevice)
unsigned int openPlaybackDevice(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  modeID, const char*  playbackDevice) {
	
	return funcs.openPlaybackDevice(serverConnectionHandlerID, modeID, playbackDevice);
}

// unsigned int openCaptureDevice(uint64  serverConnectionHandlerID, const char*  modeID, const char*  captureDevice)
unsigned int openCaptureDevice(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  modeID, const char*  captureDevice) {
	
	return funcs.openCaptureDevice(serverConnectionHandlerID, modeID, captureDevice);
}

// unsigned int getCurrentPlaybackDeviceName(uint64  serverConnectionHandlerID, char**  result, int*  isDefault)
unsigned int getCurrentPlaybackDeviceName(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, char**  result, int*  isDefault) {
	
	return funcs.getCurrentPlaybackDeviceName(serverConnectionHandlerID, result, isDefault);
}

// unsigned int getCurrentPlayBackMode(uint64  serverConnectionHandlerID, char**  result)
unsigned int getCurrentPlayBackMode(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, char**  result) {
	
	return funcs.getCurrentPlayBackMode(serverConnectionHandlerID, result);
}

// unsigned int getCurrentCaptureDeviceName(uint64  serverConnectionHandlerID, char**  result, int*  isDefault)
unsigned int getCurrentCaptureDeviceName(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, char**  result, int*  isDefault) {
	
	return funcs.getCurrentCaptureDeviceName(serverConnectionHandlerID, result, isDefault);
}

// unsigned int getCurrentCaptureMode(uint64  serverConnectionHandlerID, char**  result)
unsigned int getCurrentCaptureMode(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, char**  result) {
	
	return funcs.getCurrentCaptureMode(serverConnectionHandlerID, result);
}

// unsigned int initiateGracefulPlaybackShutdown(uint64  serverConnectionHandlerID)
unsigned int initiateGracefulPlaybackShutdown(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.initiateGracefulPlaybackShutdown(serverConnectionHandlerID);
}

// unsigned int closePlaybackDevice(uint64  serverConnectionHandlerID)
unsigned int closePlaybackDevice(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.closePlaybackDevice(serverConnectionHandlerID);
}

// unsigned int closeCaptureDevice(uint64  serverConnectionHandlerID)
unsigned int closeCaptureDevice(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.closeCaptureDevice(serverConnectionHandlerID);
}

// unsigned int activateCaptureDevice(uint64  serverConnectionHandlerID)
unsigned int activateCaptureDevice(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.activateCaptureDevice(serverConnectionHandlerID);
}

// unsigned int playWaveFileHandle(uint64  serverConnectionHandlerID, const char*  path, int  loop, uint64*  waveHandle)
unsigned int playWaveFileHandle(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  path, int  loop, uint64*  waveHandle) {
	
	return funcs.playWaveFileHandle(serverConnectionHandlerID, path, loop, waveHandle);
}

// unsigned int pauseWaveFileHandle(uint64  serverConnectionHandlerID, uint64  waveHandle, int  pause)
unsigned int pauseWaveFileHandle(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  waveHandle, int  pause) {
	
	return funcs.pauseWaveFileHandle(serverConnectionHandlerID, waveHandle, pause);
}

// unsigned int closeWaveFileHandle(uint64  serverConnectionHandlerID, uint64  waveHandle)
unsigned int closeWaveFileHandle(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  waveHandle) {
	
	return funcs.closeWaveFileHandle(serverConnectionHandlerID, waveHandle);
}

// unsigned int playWaveFile(uint64  serverConnectionHandlerID, const char*  path)
unsigned int playWaveFile(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  path) {
	
	return funcs.playWaveFile(serverConnectionHandlerID, path);
}

// unsigned int registerCustomDevice(const char*  deviceID, const char*  deviceDisplayName, int  capFrequency, int  capChannels, int  playFrequency, int  playChannels)
unsigned int registerCustomDevice(const struct TS3Functions funcs, const char*  deviceID, const char*  deviceDisplayName, int  capFrequency, int  capChannels, int  playFrequency, int  playChannels) {
	
	return funcs.registerCustomDevice(deviceID, deviceDisplayName, capFrequency, capChannels, playFrequency, playChannels);
}

// unsigned int unregisterCustomDevice(const char*  deviceID)
unsigned int unregisterCustomDevice(const struct TS3Functions funcs, const char*  deviceID) {
	
	return funcs.unregisterCustomDevice(deviceID);
}

// unsigned int processCustomCaptureData(const char*  deviceName, const short*  buffer, int  samples)
unsigned int processCustomCaptureData(const struct TS3Functions funcs, const char*  deviceName, const short*  buffer, int  samples) {
	
	return funcs.processCustomCaptureData(deviceName, buffer, samples);
}

// unsigned int acquireCustomPlaybackData(const char*  deviceName, short*  buffer, int  samples)
unsigned int acquireCustomPlaybackData(const struct TS3Functions funcs, const char*  deviceName, short*  buffer, int  samples) {
	
	return funcs.acquireCustomPlaybackData(deviceName, buffer, samples);
}

// unsigned int getPreProcessorInfoValueFloat(uint64  serverConnectionHandlerID, const char*  ident, float*  result)
unsigned int getPreProcessorInfoValueFloat(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ident, float*  result) {
	
	return funcs.getPreProcessorInfoValueFloat(serverConnectionHandlerID, ident, result);
}

// unsigned int getPreProcessorConfigValue(uint64  serverConnectionHandlerID, const char*  ident, char**  result)
unsigned int getPreProcessorConfigValue(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ident, char**  result) {
	
	return funcs.getPreProcessorConfigValue(serverConnectionHandlerID, ident, result);
}

// unsigned int setPreProcessorConfigValue(uint64  serverConnectionHandlerID, const char*  ident, const char*  value)
unsigned int setPreProcessorConfigValue(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ident, const char*  value) {
	
	return funcs.setPreProcessorConfigValue(serverConnectionHandlerID, ident, value);
}

// unsigned int getEncodeConfigValue(uint64  serverConnectionHandlerID, const char*  ident, char**  result)
unsigned int getEncodeConfigValue(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ident, char**  result) {
	
	return funcs.getEncodeConfigValue(serverConnectionHandlerID, ident, result);
}

// unsigned int getPlaybackConfigValueAsFloat(uint64  serverConnectionHandlerID, const char*  ident, float*  result)
unsigned int getPlaybackConfigValueAsFloat(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ident, float*  result) {
	
	return funcs.getPlaybackConfigValueAsFloat(serverConnectionHandlerID, ident, result);
}

// unsigned int setPlaybackConfigValue(uint64  serverConnectionHandlerID, const char*  ident, const char*  value)
unsigned int setPlaybackConfigValue(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ident, const char*  value) {
	
	return funcs.setPlaybackConfigValue(serverConnectionHandlerID, ident, value);
}

// unsigned int setClientVolumeModifier(uint64  serverConnectionHandlerID, anyID  clientID, float  value)
unsigned int setClientVolumeModifier(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, float  value) {
	
	return funcs.setClientVolumeModifier(serverConnectionHandlerID, clientID, value);
}

// unsigned int startVoiceRecording(uint64  serverConnectionHandlerID)
unsigned int startVoiceRecording(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.startVoiceRecording(serverConnectionHandlerID);
}

// unsigned int stopVoiceRecording(uint64  serverConnectionHandlerID)
unsigned int stopVoiceRecording(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.stopVoiceRecording(serverConnectionHandlerID);
}

// unsigned int systemset3DListenerAttributes(uint64  serverConnectionHandlerID, const TS3_VECTOR*  position, const TS3_VECTOR*  forward, const TS3_VECTOR*  up)
unsigned int systemset3DListenerAttributes(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const TS3_VECTOR*  position, const TS3_VECTOR*  forward, const TS3_VECTOR*  up) {
	
	return funcs.systemset3DListenerAttributes(serverConnectionHandlerID, position, forward, up);
}

// unsigned int set3DWaveAttributes(uint64  serverConnectionHandlerID, uint64  waveHandle, const TS3_VECTOR*  position)
unsigned int set3DWaveAttributes(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  waveHandle, const TS3_VECTOR*  position) {
	
	return funcs.set3DWaveAttributes(serverConnectionHandlerID, waveHandle, position);
}

// unsigned int systemset3DSettings(uint64  serverConnectionHandlerID, float  distanceFactor, float  rolloffScale)
unsigned int systemset3DSettings(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, float  distanceFactor, float  rolloffScale) {
	
	return funcs.systemset3DSettings(serverConnectionHandlerID, distanceFactor, rolloffScale);
}

// unsigned int channelset3DAttributes(uint64  serverConnectionHandlerID, anyID  clientID, const TS3_VECTOR*  position)
unsigned int channelset3DAttributes(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const TS3_VECTOR*  position) {
	
	return funcs.channelset3DAttributes(serverConnectionHandlerID, clientID, position);
}

// unsigned int startConnection(uint64  serverConnectionHandlerID, const char*  identity, const char*  ip, unsigned int  port, const char*  nickname, const char**  defaultChannelArray, const char*  defaultChannelPassword, const char*  serverPassword)
unsigned int startConnection(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  identity, const char*  ip, unsigned int  port, const char*  nickname, const char**  defaultChannelArray, const char*  defaultChannelPassword, const char*  serverPassword) {
	
	return funcs.startConnection(serverConnectionHandlerID, identity, ip, port, nickname, defaultChannelArray, defaultChannelPassword, serverPassword);
}

// unsigned int stopConnection(uint64  serverConnectionHandlerID, const char*  quitMessage)
unsigned int stopConnection(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  quitMessage) {
	
	return funcs.stopConnection(serverConnectionHandlerID, quitMessage);
}

// unsigned int requestClientMove(uint64  serverConnectionHandlerID, anyID  clientID, uint64  newChannelID, const char*  password, const char*  returnCode)
unsigned int requestClientMove(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, uint64  newChannelID, const char*  password, const char*  returnCode) {
	
	return funcs.requestClientMove(serverConnectionHandlerID, clientID, newChannelID, password, returnCode);
}

// unsigned int requestClientVariables(uint64  serverConnectionHandlerID, anyID  clientID, const char*  returnCode)
unsigned int requestClientVariables(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  returnCode) {
	
	return funcs.requestClientVariables(serverConnectionHandlerID, clientID, returnCode);
}

// unsigned int requestClientKickFromChannel(uint64  serverConnectionHandlerID, anyID  clientID, const char*  kickReason, const char*  returnCode)
unsigned int requestClientKickFromChannel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  kickReason, const char*  returnCode) {
	
	return funcs.requestClientKickFromChannel(serverConnectionHandlerID, clientID, kickReason, returnCode);
}

// unsigned int requestClientKickFromServer(uint64  serverConnectionHandlerID, anyID  clientID, const char*  kickReason, const char*  returnCode)
unsigned int requestClientKickFromServer(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  kickReason, const char*  returnCode) {
	
	return funcs.requestClientKickFromServer(serverConnectionHandlerID, clientID, kickReason, returnCode);
}

// unsigned int requestChannelDelete(uint64  serverConnectionHandlerID, uint64  channelID, int  force, const char*  returnCode)
unsigned int requestChannelDelete(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, int  force, const char*  returnCode) {
	
	return funcs.requestChannelDelete(serverConnectionHandlerID, channelID, force, returnCode);
}

// unsigned int requestChannelMove(uint64  serverConnectionHandlerID, uint64  channelID, uint64  newChannelParentID, uint64  newChannelOrder, const char*  returnCode)
unsigned int requestChannelMove(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, uint64  newChannelParentID, uint64  newChannelOrder, const char*  returnCode) {
	
	return funcs.requestChannelMove(serverConnectionHandlerID, channelID, newChannelParentID, newChannelOrder, returnCode);
}

// unsigned int requestSendPrivateTextMsg(uint64  serverConnectionHandlerID, const char*  message, anyID  targetClientID, const char*  returnCode)
unsigned int requestSendPrivateTextMsg(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  message, anyID  targetClientID, const char*  returnCode) {
	
	return funcs.requestSendPrivateTextMsg(serverConnectionHandlerID, message, targetClientID, returnCode);
}

// unsigned int requestSendChannelTextMsg(uint64  serverConnectionHandlerID, const char*  message, uint64  targetChannelID, const char*  returnCode)
unsigned int requestSendChannelTextMsg(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  message, uint64  targetChannelID, const char*  returnCode) {
	
	return funcs.requestSendChannelTextMsg(serverConnectionHandlerID, message, targetChannelID, returnCode);
}

// unsigned int requestSendServerTextMsg(uint64  serverConnectionHandlerID, const char*  message, const char*  returnCode)
unsigned int requestSendServerTextMsg(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  message, const char*  returnCode) {
	
	return funcs.requestSendServerTextMsg(serverConnectionHandlerID, message, returnCode);
}

// unsigned int requestConnectionInfo(uint64  serverConnectionHandlerID, anyID  clientID, const char*  returnCode)
unsigned int requestConnectionInfo(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  returnCode) {
	
	return funcs.requestConnectionInfo(serverConnectionHandlerID, clientID, returnCode);
}

// unsigned int requestClientSetWhisperList(uint64  serverConnectionHandlerID, anyID  clientID, const uint64*  targetChannelIDArray, const anyID*  targetClientIDArray, const char*  returnCode)
unsigned int requestClientSetWhisperList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const uint64*  targetChannelIDArray, const anyID*  targetClientIDArray, const char*  returnCode) {
	
	return funcs.requestClientSetWhisperList(serverConnectionHandlerID, clientID, targetChannelIDArray, targetClientIDArray, returnCode);
}

// unsigned int requestChannelSubscribe(uint64  serverConnectionHandlerID, const uint64*  channelIDArray, const char*  returnCode)
unsigned int requestChannelSubscribe(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const uint64*  channelIDArray, const char*  returnCode) {
	
	return funcs.requestChannelSubscribe(serverConnectionHandlerID, channelIDArray, returnCode);
}

// unsigned int requestChannelSubscribeAll(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestChannelSubscribeAll(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestChannelSubscribeAll(serverConnectionHandlerID, returnCode);
}

// unsigned int requestChannelUnsubscribe(uint64  serverConnectionHandlerID, const uint64*  channelIDArray, const char*  returnCode)
unsigned int requestChannelUnsubscribe(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const uint64*  channelIDArray, const char*  returnCode) {
	
	return funcs.requestChannelUnsubscribe(serverConnectionHandlerID, channelIDArray, returnCode);
}

// unsigned int requestChannelUnsubscribeAll(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestChannelUnsubscribeAll(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestChannelUnsubscribeAll(serverConnectionHandlerID, returnCode);
}

// unsigned int requestChannelDescription(uint64  serverConnectionHandlerID, uint64  channelID, const char*  returnCode)
unsigned int requestChannelDescription(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  returnCode) {
	
	return funcs.requestChannelDescription(serverConnectionHandlerID, channelID, returnCode);
}

// unsigned int requestMuteClients(uint64  serverConnectionHandlerID, const anyID*  clientIDArray, const char*  returnCode)
unsigned int requestMuteClients(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const anyID*  clientIDArray, const char*  returnCode) {
	
	return funcs.requestMuteClients(serverConnectionHandlerID, clientIDArray, returnCode);
}

// unsigned int requestUnmuteClients(uint64  serverConnectionHandlerID, const anyID*  clientIDArray, const char*  returnCode)
unsigned int requestUnmuteClients(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const anyID*  clientIDArray, const char*  returnCode) {
	
	return funcs.requestUnmuteClients(serverConnectionHandlerID, clientIDArray, returnCode);
}

// unsigned int requestClientPoke(uint64  serverConnectionHandlerID, anyID  clientID, const char*  message, const char*  returnCode)
unsigned int requestClientPoke(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  message, const char*  returnCode) {
	
	return funcs.requestClientPoke(serverConnectionHandlerID, clientID, message, returnCode);
}

// unsigned int requestClientIDs(uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, const char*  returnCode)
unsigned int requestClientIDs(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, const char*  returnCode) {
	
	return funcs.requestClientIDs(serverConnectionHandlerID, clientUniqueIdentifier, returnCode);
}

// unsigned int clientChatClosed(uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, anyID  clientID, const char*  returnCode)
unsigned int clientChatClosed(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, anyID  clientID, const char*  returnCode) {
	
	return funcs.clientChatClosed(serverConnectionHandlerID, clientUniqueIdentifier, clientID, returnCode);
}

// unsigned int clientChatComposing(uint64  serverConnectionHandlerID, anyID  clientID, const char*  returnCode)
unsigned int clientChatComposing(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  returnCode) {
	
	return funcs.clientChatComposing(serverConnectionHandlerID, clientID, returnCode);
}

// unsigned int requestServerTemporaryPasswordAdd(uint64  serverConnectionHandlerID, const char*  password, const char*  description, uint64  duration, uint64  targetChannelID, const char*  targetChannelPW, const char*  returnCode)
unsigned int requestServerTemporaryPasswordAdd(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  password, const char*  description, uint64  duration, uint64  targetChannelID, const char*  targetChannelPW, const char*  returnCode) {
	
	return funcs.requestServerTemporaryPasswordAdd(serverConnectionHandlerID, password, description, duration, targetChannelID, targetChannelPW, returnCode);
}

// unsigned int requestServerTemporaryPasswordDel(uint64  serverConnectionHandlerID, const char*  password, const char*  returnCode)
unsigned int requestServerTemporaryPasswordDel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  password, const char*  returnCode) {
	
	return funcs.requestServerTemporaryPasswordDel(serverConnectionHandlerID, password, returnCode);
}

// unsigned int requestServerTemporaryPasswordList(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestServerTemporaryPasswordList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestServerTemporaryPasswordList(serverConnectionHandlerID, returnCode);
}

// unsigned int getClientID(uint64  serverConnectionHandlerID, anyID*  result)
unsigned int getClientID(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID*  result) {
	
	return funcs.getClientID(serverConnectionHandlerID, result);
}

// unsigned int getClientSelfVariableAsInt(uint64  serverConnectionHandlerID, size_t  flag, int*  result)
unsigned int getClientSelfVariableAsInt(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, int*  result) {
	
	return funcs.getClientSelfVariableAsInt(serverConnectionHandlerID, flag, result);
}

// unsigned int getClientSelfVariableAsString(uint64  serverConnectionHandlerID, size_t  flag, char**  result)
unsigned int getClientSelfVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, char**  result) {
	
	return funcs.getClientSelfVariableAsString(serverConnectionHandlerID, flag, result);
}

// unsigned int setClientSelfVariableAsInt(uint64  serverConnectionHandlerID, size_t  flag, int  value)
unsigned int setClientSelfVariableAsInt(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, int  value) {
	
	return funcs.setClientSelfVariableAsInt(serverConnectionHandlerID, flag, value);
}

// unsigned int setClientSelfVariableAsString(uint64  serverConnectionHandlerID, size_t  flag, const char*  value)
unsigned int setClientSelfVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, const char*  value) {
	
	return funcs.setClientSelfVariableAsString(serverConnectionHandlerID, flag, value);
}

// unsigned int flushClientSelfUpdates(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int flushClientSelfUpdates(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.flushClientSelfUpdates(serverConnectionHandlerID, returnCode);
}

// unsigned int getClientVariableAsInt(uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, int*  result)
unsigned int getClientVariableAsInt(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, int*  result) {
	
	return funcs.getClientVariableAsInt(serverConnectionHandlerID, clientID, flag, result);
}

// unsigned int getClientVariableAsUInt64(uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, uint64*  result)
unsigned int getClientVariableAsUInt64(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, uint64*  result) {
	
	return funcs.getClientVariableAsUInt64(serverConnectionHandlerID, clientID, flag, result);
}

// unsigned int getClientVariableAsString(uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, char**  result)
unsigned int getClientVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, char**  result) {
	
	return funcs.getClientVariableAsString(serverConnectionHandlerID, clientID, flag, result);
}

// unsigned int getClientList(uint64  serverConnectionHandlerID, anyID**  result)
unsigned int getClientList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID**  result) {
	
	return funcs.getClientList(serverConnectionHandlerID, result);
}

// unsigned int getChannelOfClient(uint64  serverConnectionHandlerID, anyID  clientID, uint64*  result)
unsigned int getChannelOfClient(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, uint64*  result) {
	
	return funcs.getChannelOfClient(serverConnectionHandlerID, clientID, result);
}

// unsigned int getChannelVariableAsInt(uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, int*  result)
unsigned int getChannelVariableAsInt(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, int*  result) {
	
	return funcs.getChannelVariableAsInt(serverConnectionHandlerID, channelID, flag, result);
}

// unsigned int getChannelVariableAsUInt64(uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, uint64*  result)
unsigned int getChannelVariableAsUInt64(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, uint64*  result) {
	
	return funcs.getChannelVariableAsUInt64(serverConnectionHandlerID, channelID, flag, result);
}

// unsigned int getChannelVariableAsString(uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, char**  result)
unsigned int getChannelVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, char**  result) {
	
	return funcs.getChannelVariableAsString(serverConnectionHandlerID, channelID, flag, result);
}

// unsigned int getChannelIDFromChannelNames(uint64  serverConnectionHandlerID, char**  channelNameArray, uint64*  result)
unsigned int getChannelIDFromChannelNames(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, char**  channelNameArray, uint64*  result) {
	
	return funcs.getChannelIDFromChannelNames(serverConnectionHandlerID, channelNameArray, result);
}

// unsigned int setChannelVariableAsInt(uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, int  value)
unsigned int setChannelVariableAsInt(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, int  value) {
	
	return funcs.setChannelVariableAsInt(serverConnectionHandlerID, channelID, flag, value);
}

// unsigned int setChannelVariableAsUInt64(uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, uint64  value)
unsigned int setChannelVariableAsUInt64(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, uint64  value) {
	
	return funcs.setChannelVariableAsUInt64(serverConnectionHandlerID, channelID, flag, value);
}

// unsigned int setChannelVariableAsString(uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, const char*  value)
unsigned int setChannelVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, size_t  flag, const char*  value) {
	
	return funcs.setChannelVariableAsString(serverConnectionHandlerID, channelID, flag, value);
}

// unsigned int flushChannelUpdates(uint64  serverConnectionHandlerID, uint64  channelID, const char*  returnCode)
unsigned int flushChannelUpdates(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  returnCode) {
	
	return funcs.flushChannelUpdates(serverConnectionHandlerID, channelID, returnCode);
}

// unsigned int flushChannelCreation(uint64  serverConnectionHandlerID, uint64  channelParentID, const char*  returnCode)
unsigned int flushChannelCreation(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelParentID, const char*  returnCode) {
	
	return funcs.flushChannelCreation(serverConnectionHandlerID, channelParentID, returnCode);
}

// unsigned int getChannelList(uint64  serverConnectionHandlerID, uint64**  result)
unsigned int getChannelList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64**  result) {
	
	return funcs.getChannelList(serverConnectionHandlerID, result);
}

// unsigned int getChannelClientList(uint64  serverConnectionHandlerID, uint64  channelID, anyID**  result)
unsigned int getChannelClientList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, anyID**  result) {
	
	return funcs.getChannelClientList(serverConnectionHandlerID, channelID, result);
}

// unsigned int getParentChannelOfChannel(uint64  serverConnectionHandlerID, uint64  channelID, uint64*  result)
unsigned int getParentChannelOfChannel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, uint64*  result) {
	
	return funcs.getParentChannelOfChannel(serverConnectionHandlerID, channelID, result);
}

// unsigned int getServerConnectionHandlerList(uint64**  result)
unsigned int getServerConnectionHandlerList(const struct TS3Functions funcs, uint64**  result) {
	
	return funcs.getServerConnectionHandlerList(result);
}

// unsigned int getServerVariableAsInt(uint64  serverConnectionHandlerID, size_t  flag, int*  result)
unsigned int getServerVariableAsInt(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, int*  result) {
	
	return funcs.getServerVariableAsInt(serverConnectionHandlerID, flag, result);
}

// unsigned int getServerVariableAsUInt64(uint64  serverConnectionHandlerID, size_t  flag, uint64*  result)
unsigned int getServerVariableAsUInt64(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, uint64*  result) {
	
	return funcs.getServerVariableAsUInt64(serverConnectionHandlerID, flag, result);
}

// unsigned int getServerVariableAsString(uint64  serverConnectionHandlerID, size_t  flag, char**  result)
unsigned int getServerVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, size_t  flag, char**  result) {
	
	return funcs.getServerVariableAsString(serverConnectionHandlerID, flag, result);
}

// unsigned int requestServerVariables(uint64  serverConnectionHandlerID)
unsigned int requestServerVariables(const struct TS3Functions funcs, uint64  serverConnectionHandlerID) {
	
	return funcs.requestServerVariables(serverConnectionHandlerID);
}

// unsigned int getConnectionStatus(uint64  serverConnectionHandlerID, int*  result)
unsigned int getConnectionStatus(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, int*  result) {
	
	return funcs.getConnectionStatus(serverConnectionHandlerID, result);
}

// unsigned int getConnectionVariableAsUInt64(uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, uint64*  result)
unsigned int getConnectionVariableAsUInt64(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, uint64*  result) {
	
	return funcs.getConnectionVariableAsUInt64(serverConnectionHandlerID, clientID, flag, result);
}

// unsigned int getConnectionVariableAsDouble(uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, double*  result)
unsigned int getConnectionVariableAsDouble(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, double*  result) {
	
	return funcs.getConnectionVariableAsDouble(serverConnectionHandlerID, clientID, flag, result);
}

// unsigned int getConnectionVariableAsString(uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, char**  result)
unsigned int getConnectionVariableAsString(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, size_t  flag, char**  result) {
	
	return funcs.getConnectionVariableAsString(serverConnectionHandlerID, clientID, flag, result);
}

// unsigned int cleanUpConnectionInfo(uint64  serverConnectionHandlerID, anyID  clientID)
unsigned int cleanUpConnectionInfo(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID) {
	
	return funcs.cleanUpConnectionInfo(serverConnectionHandlerID, clientID);
}

// unsigned int requestClientDBIDfromUID(uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, const char*  returnCode)
unsigned int requestClientDBIDfromUID(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, const char*  returnCode) {
	
	return funcs.requestClientDBIDfromUID(serverConnectionHandlerID, clientUniqueIdentifier, returnCode);
}

// unsigned int requestClientNamefromUID(uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, const char*  returnCode)
unsigned int requestClientNamefromUID(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  clientUniqueIdentifier, const char*  returnCode) {
	
	return funcs.requestClientNamefromUID(serverConnectionHandlerID, clientUniqueIdentifier, returnCode);
}

// unsigned int requestClientNamefromDBID(uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const char*  returnCode)
unsigned int requestClientNamefromDBID(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const char*  returnCode) {
	
	return funcs.requestClientNamefromDBID(serverConnectionHandlerID, clientDatabaseID, returnCode);
}

// unsigned int requestClientEditDescription(uint64  serverConnectionHandlerID, anyID  clientID, const char*  clientDescription, const char*  returnCode)
unsigned int requestClientEditDescription(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, const char*  clientDescription, const char*  returnCode) {
	
	return funcs.requestClientEditDescription(serverConnectionHandlerID, clientID, clientDescription, returnCode);
}

// unsigned int requestClientSetIsTalker(uint64  serverConnectionHandlerID, anyID  clientID, int  isTalker, const char*  returnCode)
unsigned int requestClientSetIsTalker(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, int  isTalker, const char*  returnCode) {
	
	return funcs.requestClientSetIsTalker(serverConnectionHandlerID, clientID, isTalker, returnCode);
}

// unsigned int requestIsTalker(uint64  serverConnectionHandlerID, int  isTalkerRequest, const char*  isTalkerRequestMessage, const char*  returnCode)
unsigned int requestIsTalker(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, int  isTalkerRequest, const char*  isTalkerRequestMessage, const char*  returnCode) {
	
	return funcs.requestIsTalker(serverConnectionHandlerID, isTalkerRequest, isTalkerRequestMessage, returnCode);
}

// unsigned int requestSendClientQueryCommand(uint64  serverConnectionHandlerID, const char*  command, const char*  returnCode)
unsigned int requestSendClientQueryCommand(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  command, const char*  returnCode) {
	
	return funcs.requestSendClientQueryCommand(serverConnectionHandlerID, command, returnCode);
}

// unsigned int getTransferFileName(anyID  transferID, char**  result)
unsigned int getTransferFileName(const struct TS3Functions funcs, anyID  transferID, char**  result) {
	
	return funcs.getTransferFileName(transferID, result);
}

// unsigned int getTransferFilePath(anyID  transferID, char**  result)
unsigned int getTransferFilePath(const struct TS3Functions funcs, anyID  transferID, char**  result) {
	
	return funcs.getTransferFilePath(transferID, result);
}

// unsigned int getTransferFileSize(anyID  transferID, uint64*  result)
unsigned int getTransferFileSize(const struct TS3Functions funcs, anyID  transferID, uint64*  result) {
	
	return funcs.getTransferFileSize(transferID, result);
}

// unsigned int getTransferFileSizeDone(anyID  transferID, uint64*  result)
unsigned int getTransferFileSizeDone(const struct TS3Functions funcs, anyID  transferID, uint64*  result) {
	
	return funcs.getTransferFileSizeDone(transferID, result);
}

// unsigned int isTransferSender(anyID  transferID, int*  result)
unsigned int isTransferSender(const struct TS3Functions funcs, anyID  transferID, int*  result) {
	
	return funcs.isTransferSender(transferID, result);
}

// unsigned int getTransferStatus(anyID  transferID, int*  result)
unsigned int getTransferStatus(const struct TS3Functions funcs, anyID  transferID, int*  result) {
	
	return funcs.getTransferStatus(transferID, result);
}

// unsigned int getCurrentTransferSpeed(anyID  transferID, float*  result)
unsigned int getCurrentTransferSpeed(const struct TS3Functions funcs, anyID  transferID, float*  result) {
	
	return funcs.getCurrentTransferSpeed(transferID, result);
}

// unsigned int getAverageTransferSpeed(anyID  transferID, float*  result)
unsigned int getAverageTransferSpeed(const struct TS3Functions funcs, anyID  transferID, float*  result) {
	
	return funcs.getAverageTransferSpeed(transferID, result);
}

// unsigned int getTransferRunTime(anyID  transferID, uint64*  result)
unsigned int getTransferRunTime(const struct TS3Functions funcs, anyID  transferID, uint64*  result) {
	
	return funcs.getTransferRunTime(transferID, result);
}

// unsigned int sendFile(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  file, int  overwrite, int  resume, const char*  sourceDirectory, anyID*  result, const char*  returnCode)
unsigned int sendFile(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  file, int  overwrite, int  resume, const char*  sourceDirectory, anyID*  result, const char*  returnCode) {
	
	return funcs.sendFile(serverConnectionHandlerID, channelID, channelPW, file, overwrite, resume, sourceDirectory, result, returnCode);
}

// unsigned int requestFile(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  file, int  overwrite, int  resume, const char*  destinationDirectory, anyID*  result, const char*  returnCode)
unsigned int requestFile(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  file, int  overwrite, int  resume, const char*  destinationDirectory, anyID*  result, const char*  returnCode) {
	
	return funcs.requestFile(serverConnectionHandlerID, channelID, channelPW, file, overwrite, resume, destinationDirectory, result, returnCode);
}

// unsigned int haltTransfer(uint64  serverConnectionHandlerID, anyID  transferID, int  deleteUnfinishedFile, const char*  returnCode)
unsigned int haltTransfer(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  transferID, int  deleteUnfinishedFile, const char*  returnCode) {
	
	return funcs.haltTransfer(serverConnectionHandlerID, transferID, deleteUnfinishedFile, returnCode);
}

// unsigned int requestFileList(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  path, const char*  returnCode)
unsigned int requestFileList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  path, const char*  returnCode) {
	
	return funcs.requestFileList(serverConnectionHandlerID, channelID, channelPW, path, returnCode);
}

// unsigned int requestFileInfo(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  file, const char*  returnCode)
unsigned int requestFileInfo(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  file, const char*  returnCode) {
	
	return funcs.requestFileInfo(serverConnectionHandlerID, channelID, channelPW, file, returnCode);
}

// unsigned int requestDeleteFile(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char**  file, const char*  returnCode)
unsigned int requestDeleteFile(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char**  file, const char*  returnCode) {
	
	return funcs.requestDeleteFile(serverConnectionHandlerID, channelID, channelPW, file, returnCode);
}

// unsigned int requestCreateDirectory(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  directoryPath, const char*  returnCode)
unsigned int requestCreateDirectory(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPW, const char*  directoryPath, const char*  returnCode) {
	
	return funcs.requestCreateDirectory(serverConnectionHandlerID, channelID, channelPW, directoryPath, returnCode);
}

// unsigned int requestRenameFile(uint64  serverConnectionHandlerID, uint64  fromChannelID, const char*  channelPW, uint64  toChannelID, const char*  toChannelPW, const char*  oldFile, const char*  newFile, const char*  returnCode)
unsigned int requestRenameFile(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  fromChannelID, const char*  channelPW, uint64  toChannelID, const char*  toChannelPW, const char*  oldFile, const char*  newFile, const char*  returnCode) {
	
	return funcs.requestRenameFile(serverConnectionHandlerID, fromChannelID, channelPW, toChannelID, toChannelPW, oldFile, newFile, returnCode);
}

// unsigned int requestMessageAdd(uint64  serverConnectionHandlerID, const char*  toClientUID, const char*  subject, const char*  message, const char*  returnCode)
unsigned int requestMessageAdd(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  toClientUID, const char*  subject, const char*  message, const char*  returnCode) {
	
	return funcs.requestMessageAdd(serverConnectionHandlerID, toClientUID, subject, message, returnCode);
}

// unsigned int requestMessageDel(uint64  serverConnectionHandlerID, uint64  messageID, const char*  returnCode)
unsigned int requestMessageDel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  messageID, const char*  returnCode) {
	
	return funcs.requestMessageDel(serverConnectionHandlerID, messageID, returnCode);
}

// unsigned int requestMessageGet(uint64  serverConnectionHandlerID, uint64  messageID, const char*  returnCode)
unsigned int requestMessageGet(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  messageID, const char*  returnCode) {
	
	return funcs.requestMessageGet(serverConnectionHandlerID, messageID, returnCode);
}

// unsigned int requestMessageList(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestMessageList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestMessageList(serverConnectionHandlerID, returnCode);
}

// unsigned int requestMessageUpdateFlag(uint64  serverConnectionHandlerID, uint64  messageID, int  flag, const char*  returnCode)
unsigned int requestMessageUpdateFlag(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  messageID, int  flag, const char*  returnCode) {
	
	return funcs.requestMessageUpdateFlag(serverConnectionHandlerID, messageID, flag, returnCode);
}

// unsigned int verifyServerPassword(uint64  serverConnectionHandlerID, const char*  serverPassword, const char*  returnCode)
unsigned int verifyServerPassword(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  serverPassword, const char*  returnCode) {
	
	return funcs.verifyServerPassword(serverConnectionHandlerID, serverPassword, returnCode);
}

// unsigned int verifyChannelPassword(uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPassword, const char*  returnCode)
unsigned int verifyChannelPassword(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  channelPassword, const char*  returnCode) {
	
	return funcs.verifyChannelPassword(serverConnectionHandlerID, channelID, channelPassword, returnCode);
}

// unsigned int banclient(uint64  serverConnectionHandlerID, anyID  clientID, uint64  timeInSeconds, const char*  banReason, const char*  returnCode)
unsigned int banclient(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, anyID  clientID, uint64  timeInSeconds, const char*  banReason, const char*  returnCode) {
	
	return funcs.banclient(serverConnectionHandlerID, clientID, timeInSeconds, banReason, returnCode);
}

// unsigned int banadd(uint64  serverConnectionHandlerID, const char*  ipRegExp, const char*  nameRegexp, const char*  uniqueIdentity, uint64  timeInSeconds, const char*  banReason, const char*  returnCode)
unsigned int banadd(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  ipRegExp, const char*  nameRegexp, const char*  uniqueIdentity, uint64  timeInSeconds, const char*  banReason, const char*  returnCode) {
	
	return funcs.banadd(serverConnectionHandlerID, ipRegExp, nameRegexp, uniqueIdentity, timeInSeconds, banReason, returnCode);
}

// unsigned int banclientdbid(uint64  serverConnectionHandlerID, uint64  clientDBID, uint64  timeInSeconds, const char*  banReason, const char*  returnCode)
unsigned int banclientdbid(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  clientDBID, uint64  timeInSeconds, const char*  banReason, const char*  returnCode) {
	
	return funcs.banclientdbid(serverConnectionHandlerID, clientDBID, timeInSeconds, banReason, returnCode);
}

// unsigned int bandel(uint64  serverConnectionHandlerID, uint64  banID, const char*  returnCode)
unsigned int bandel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  banID, const char*  returnCode) {
	
	return funcs.bandel(serverConnectionHandlerID, banID, returnCode);
}

// unsigned int bandelall(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int bandelall(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.bandelall(serverConnectionHandlerID, returnCode);
}

// unsigned int requestBanList(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestBanList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestBanList(serverConnectionHandlerID, returnCode);
}

// unsigned int requestComplainAdd(uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, const char*  complainReason, const char*  returnCode)
unsigned int requestComplainAdd(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, const char*  complainReason, const char*  returnCode) {
	
	return funcs.requestComplainAdd(serverConnectionHandlerID, targetClientDatabaseID, complainReason, returnCode);
}

// unsigned int requestComplainDel(uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, uint64  fromClientDatabaseID, const char*  returnCode)
unsigned int requestComplainDel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, uint64  fromClientDatabaseID, const char*  returnCode) {
	
	return funcs.requestComplainDel(serverConnectionHandlerID, targetClientDatabaseID, fromClientDatabaseID, returnCode);
}

// unsigned int requestComplainDelAll(uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, const char*  returnCode)
unsigned int requestComplainDelAll(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, const char*  returnCode) {
	
	return funcs.requestComplainDelAll(serverConnectionHandlerID, targetClientDatabaseID, returnCode);
}

// unsigned int requestComplainList(uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, const char*  returnCode)
unsigned int requestComplainList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  targetClientDatabaseID, const char*  returnCode) {
	
	return funcs.requestComplainList(serverConnectionHandlerID, targetClientDatabaseID, returnCode);
}

// unsigned int requestServerGroupList(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestServerGroupList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestServerGroupList(serverConnectionHandlerID, returnCode);
}

// unsigned int requestServerGroupAdd(uint64  serverConnectionHandlerID, const char*  groupName, int  groupType, const char*  returnCode)
unsigned int requestServerGroupAdd(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  groupName, int  groupType, const char*  returnCode) {
	
	return funcs.requestServerGroupAdd(serverConnectionHandlerID, groupName, groupType, returnCode);
}

// unsigned int requestServerGroupDel(uint64  serverConnectionHandlerID, uint64  serverGroupID, int  force, const char*  returnCode)
unsigned int requestServerGroupDel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, int  force, const char*  returnCode) {
	
	return funcs.requestServerGroupDel(serverConnectionHandlerID, serverGroupID, force, returnCode);
}

// unsigned int requestServerGroupAddClient(uint64  serverConnectionHandlerID, uint64  serverGroupID, uint64  clientDatabaseID, const char*  returnCode)
unsigned int requestServerGroupAddClient(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, uint64  clientDatabaseID, const char*  returnCode) {
	
	return funcs.requestServerGroupAddClient(serverConnectionHandlerID, serverGroupID, clientDatabaseID, returnCode);
}

// unsigned int requestServerGroupDelClient(uint64  serverConnectionHandlerID, uint64  serverGroupID, uint64  clientDatabaseID, const char*  returnCode)
unsigned int requestServerGroupDelClient(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, uint64  clientDatabaseID, const char*  returnCode) {
	
	return funcs.requestServerGroupDelClient(serverConnectionHandlerID, serverGroupID, clientDatabaseID, returnCode);
}

// unsigned int requestServerGroupsByClientID(uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const char*  returnCode)
unsigned int requestServerGroupsByClientID(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const char*  returnCode) {
	
	return funcs.requestServerGroupsByClientID(serverConnectionHandlerID, clientDatabaseID, returnCode);
}

// unsigned int requestServerGroupAddPerm(uint64  serverConnectionHandlerID, uint64  serverGroupID, int  continueonerror, const unsigned int*  permissionIDArray, const int*  permissionValueArray, const int*  permissionNegatedArray, const int*  permissionSkipArray, int  arraySize, const char*  returnCode)
unsigned int requestServerGroupAddPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, int  continueonerror, const unsigned int*  permissionIDArray, const int*  permissionValueArray, const int*  permissionNegatedArray, const int*  permissionSkipArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestServerGroupAddPerm(serverConnectionHandlerID, serverGroupID, continueonerror, permissionIDArray, permissionValueArray, permissionNegatedArray, permissionSkipArray, arraySize, returnCode);
}

// unsigned int requestServerGroupDelPerm(uint64  serverConnectionHandlerID, uint64  serverGroupID, int  continueOnError, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode)
unsigned int requestServerGroupDelPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, int  continueOnError, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestServerGroupDelPerm(serverConnectionHandlerID, serverGroupID, continueOnError, permissionIDArray, arraySize, returnCode);
}

// unsigned int requestServerGroupPermList(uint64  serverConnectionHandlerID, uint64  serverGroupID, const char*  returnCode)
unsigned int requestServerGroupPermList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, const char*  returnCode) {
	
	return funcs.requestServerGroupPermList(serverConnectionHandlerID, serverGroupID, returnCode);
}

// unsigned int requestServerGroupClientList(uint64  serverConnectionHandlerID, uint64  serverGroupID, int  withNames, const char*  returnCode)
unsigned int requestServerGroupClientList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  serverGroupID, int  withNames, const char*  returnCode) {
	
	return funcs.requestServerGroupClientList(serverConnectionHandlerID, serverGroupID, withNames, returnCode);
}

// unsigned int requestChannelGroupList(uint64  serverConnectionHandlerID, const char*  returnCode)
unsigned int requestChannelGroupList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  returnCode) {
	
	return funcs.requestChannelGroupList(serverConnectionHandlerID, returnCode);
}

// unsigned int requestChannelGroupAdd(uint64  serverConnectionHandlerID, const char*  groupName, int  groupType, const char*  returnCode)
unsigned int requestChannelGroupAdd(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  groupName, int  groupType, const char*  returnCode) {
	
	return funcs.requestChannelGroupAdd(serverConnectionHandlerID, groupName, groupType, returnCode);
}

// unsigned int requestChannelGroupDel(uint64  serverConnectionHandlerID, uint64  channelGroupID, int  force, const char*  returnCode)
unsigned int requestChannelGroupDel(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelGroupID, int  force, const char*  returnCode) {
	
	return funcs.requestChannelGroupDel(serverConnectionHandlerID, channelGroupID, force, returnCode);
}

// unsigned int requestChannelGroupAddPerm(uint64  serverConnectionHandlerID, uint64  channelGroupID, int  continueonerror, const unsigned int*  permissionIDArray, const int*  permissionValueArray, int  arraySize, const char*  returnCode)
unsigned int requestChannelGroupAddPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelGroupID, int  continueonerror, const unsigned int*  permissionIDArray, const int*  permissionValueArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestChannelGroupAddPerm(serverConnectionHandlerID, channelGroupID, continueonerror, permissionIDArray, permissionValueArray, arraySize, returnCode);
}

// unsigned int requestChannelGroupDelPerm(uint64  serverConnectionHandlerID, uint64  channelGroupID, int  continueOnError, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode)
unsigned int requestChannelGroupDelPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelGroupID, int  continueOnError, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestChannelGroupDelPerm(serverConnectionHandlerID, channelGroupID, continueOnError, permissionIDArray, arraySize, returnCode);
}

// unsigned int requestChannelGroupPermList(uint64  serverConnectionHandlerID, uint64  channelGroupID, const char*  returnCode)
unsigned int requestChannelGroupPermList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelGroupID, const char*  returnCode) {
	
	return funcs.requestChannelGroupPermList(serverConnectionHandlerID, channelGroupID, returnCode);
}

// unsigned int requestSetClientChannelGroup(uint64  serverConnectionHandlerID, const uint64*  channelGroupIDArray, const uint64*  channelIDArray, const uint64*  clientDatabaseIDArray, int  arraySize, const char*  returnCode)
unsigned int requestSetClientChannelGroup(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const uint64*  channelGroupIDArray, const uint64*  channelIDArray, const uint64*  clientDatabaseIDArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestSetClientChannelGroup(serverConnectionHandlerID, channelGroupIDArray, channelIDArray, clientDatabaseIDArray, arraySize, returnCode);
}

// unsigned int requestChannelAddPerm(uint64  serverConnectionHandlerID, uint64  channelID, const unsigned int*  permissionIDArray, const int*  permissionValueArray, int  arraySize, const char*  returnCode)
unsigned int requestChannelAddPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const unsigned int*  permissionIDArray, const int*  permissionValueArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestChannelAddPerm(serverConnectionHandlerID, channelID, permissionIDArray, permissionValueArray, arraySize, returnCode);
}

// unsigned int requestChannelDelPerm(uint64  serverConnectionHandlerID, uint64  channelID, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode)
unsigned int requestChannelDelPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestChannelDelPerm(serverConnectionHandlerID, channelID, permissionIDArray, arraySize, returnCode);
}

// unsigned int requestChannelPermList(uint64  serverConnectionHandlerID, uint64  channelID, const char*  returnCode)
unsigned int requestChannelPermList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, const char*  returnCode) {
	
	return funcs.requestChannelPermList(serverConnectionHandlerID, channelID, returnCode);
}

// unsigned int requestClientAddPerm(uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, const int*  permissionValueArray, const int*  permissionSkipArray, int  arraySize, const char*  returnCode)
unsigned int requestClientAddPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, const int*  permissionValueArray, const int*  permissionSkipArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestClientAddPerm(serverConnectionHandlerID, clientDatabaseID, permissionIDArray, permissionValueArray, permissionSkipArray, arraySize, returnCode);
}

// unsigned int requestClientDelPerm(uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode)
unsigned int requestClientDelPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestClientDelPerm(serverConnectionHandlerID, clientDatabaseID, permissionIDArray, arraySize, returnCode);
}

// unsigned int requestClientPermList(uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const char*  returnCode)
unsigned int requestClientPermList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  clientDatabaseID, const char*  returnCode) {
	
	return funcs.requestClientPermList(serverConnectionHandlerID, clientDatabaseID, returnCode);
}

// unsigned int requestChannelClientAddPerm(uint64  serverConnectionHandlerID, uint64  channelID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, const int*  permissionValueArray, int  arraySize, const char*  returnCode)
unsigned int requestChannelClientAddPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, const int*  permissionValueArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestChannelClientAddPerm(serverConnectionHandlerID, channelID, clientDatabaseID, permissionIDArray, permissionValueArray, arraySize, returnCode);
}

// unsigned int requestChannelClientDelPerm(uint64  serverConnectionHandlerID, uint64  channelID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode)
unsigned int requestChannelClientDelPerm(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, uint64  clientDatabaseID, const unsigned int*  permissionIDArray, int  arraySize, const char*  returnCode) {
	
	return funcs.requestChannelClientDelPerm(serverConnectionHandlerID, channelID, clientDatabaseID, permissionIDArray, arraySize, returnCode);
}

// unsigned int requestChannelClientPermList(uint64  serverConnectionHandlerID, uint64  channelID, uint64  clientDatabaseID, const char*  returnCode)
unsigned int requestChannelClientPermList(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, uint64  channelID, uint64  clientDatabaseID, const char*  returnCode) {
	
	return funcs.requestChannelClientPermList(serverConnectionHandlerID, channelID, clientDatabaseID, returnCode);
}

// unsigned int privilegeKeyUse(uint64  serverConnectionHandler, const char*  tokenKey, const char*  returnCode)
unsigned int privilegeKeyUse(const struct TS3Functions funcs, uint64  serverConnectionHandler, const char*  tokenKey, const char*  returnCode) {
	
	return funcs.privilegeKeyUse(serverConnectionHandler, tokenKey, returnCode);
}

// unsigned int requestPermissionList(uint64  serverConnectionHandler, const char*  returnCode)
unsigned int requestPermissionList(const struct TS3Functions funcs, uint64  serverConnectionHandler, const char*  returnCode) {
	
	return funcs.requestPermissionList(serverConnectionHandler, returnCode);
}

// unsigned int requestPermissionOverview(uint64  serverConnectionHandler, uint64  clientDBID, uint64  channelID, const char*  returnCode)
unsigned int requestPermissionOverview(const struct TS3Functions funcs, uint64  serverConnectionHandler, uint64  clientDBID, uint64  channelID, const char*  returnCode) {
	
	return funcs.requestPermissionOverview(serverConnectionHandler, clientDBID, channelID, returnCode);
}

// unsigned int clientPropertyStringToFlag(const char*  clientPropertyString, size_t*  resultFlag)
unsigned int clientPropertyStringToFlag(const struct TS3Functions funcs, const char*  clientPropertyString, size_t*  resultFlag) {
	
	return funcs.clientPropertyStringToFlag(clientPropertyString, resultFlag);
}

// unsigned int channelPropertyStringToFlag(const char*  channelPropertyString, size_t*  resultFlag)
unsigned int channelPropertyStringToFlag(const struct TS3Functions funcs, const char*  channelPropertyString, size_t*  resultFlag) {
	
	return funcs.channelPropertyStringToFlag(channelPropertyString, resultFlag);
}

// unsigned int serverPropertyStringToFlag(const char*  serverPropertyString, size_t*  resultFlag)
unsigned int serverPropertyStringToFlag(const struct TS3Functions funcs, const char*  serverPropertyString, size_t*  resultFlag) {
	
	return funcs.serverPropertyStringToFlag(serverPropertyString, resultFlag);
}

// void getAppPath(char*  path, size_t  maxLen)
void getAppPath(const struct TS3Functions funcs, char*  path, size_t  maxLen) {
	
	return funcs.getAppPath(path, maxLen);
}

// void getResourcesPath(char*  path, size_t  maxLen)
void getResourcesPath(const struct TS3Functions funcs, char*  path, size_t  maxLen) {
	
	return funcs.getResourcesPath(path, maxLen);
}

// void getConfigPath(char*  path, size_t  maxLen)
void getConfigPath(const struct TS3Functions funcs, char*  path, size_t  maxLen) {
	
	return funcs.getConfigPath(path, maxLen);
}

// void getPluginPath(char*  path, size_t  maxLen, const char*  pluginID)
void getPluginPath(const struct TS3Functions funcs, char*  path, size_t  maxLen, const char*  pluginID) {
	
	return funcs.getPluginPath(path, maxLen, pluginID);
}

// void printMessage(uint64  serverConnectionHandlerID, const char*  message, enum PluginMessageTarget  messageTarget)
void printMessage(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  message, enum PluginMessageTarget  messageTarget) {
	
	return funcs.printMessage(serverConnectionHandlerID, message, messageTarget);
}

// void printMessageToCurrentTab(const char*  message)
void printMessageToCurrentTab(const struct TS3Functions funcs, const char*  message) {
	
	return funcs.printMessageToCurrentTab(message);
}

// void urlsToBB(const char*  text, char*  result, size_t  maxLen)
void urlsToBB(const struct TS3Functions funcs, const char*  text, char*  result, size_t  maxLen) {
	
	return funcs.urlsToBB(text, result, maxLen);
}

// void sendPluginCommand(uint64  serverConnectionHandlerID, const char*  pluginID, const char*  command, int  targetMode, const anyID*  targetIDs, const char*  returnCode)
void sendPluginCommand(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  pluginID, const char*  command, int  targetMode, const anyID*  targetIDs, const char*  returnCode) {
	
	return funcs.sendPluginCommand(serverConnectionHandlerID, pluginID, command, targetMode, targetIDs, returnCode);
}

// void getDirectories(const char*  path, char*  result, size_t  maxLen)
void getDirectories(const struct TS3Functions funcs, const char*  path, char*  result, size_t  maxLen) {
	
	return funcs.getDirectories(path, result, maxLen);
}

// unsigned int getServerConnectInfo(uint64  scHandlerID, char*  host, unsigned short*  port, char*  password, size_t  maxLen)
unsigned int getServerConnectInfo(const struct TS3Functions funcs, uint64  scHandlerID, char*  host, unsigned short*  port, char*  password, size_t  maxLen) {
	
	return funcs.getServerConnectInfo(scHandlerID, host, port, password, maxLen);
}

// unsigned int getChannelConnectInfo(uint64  scHandlerID, uint64  channelID, char*  path, char*  password, size_t  maxLen)
unsigned int getChannelConnectInfo(const struct TS3Functions funcs, uint64  scHandlerID, uint64  channelID, char*  path, char*  password, size_t  maxLen) {
	
	return funcs.getChannelConnectInfo(scHandlerID, channelID, path, password, maxLen);
}

// void createReturnCode(const char*  pluginID, char*  returnCode, size_t  maxLen)
void createReturnCode(const struct TS3Functions funcs, const char*  pluginID, char*  returnCode, size_t  maxLen) {
	
	return funcs.createReturnCode(pluginID, returnCode, maxLen);
}

// unsigned int requestInfoUpdate(uint64  scHandlerID, enum PluginItemType  itemType, uint64  itemID)
unsigned int requestInfoUpdate(const struct TS3Functions funcs, uint64  scHandlerID, enum PluginItemType  itemType, uint64  itemID) {
	
	return funcs.requestInfoUpdate(scHandlerID, itemType, itemID);
}

// uint64 getServerVersion(uint64  scHandlerID)
uint64 getServerVersion(const struct TS3Functions funcs, uint64  scHandlerID) {
	
	return funcs.getServerVersion(scHandlerID);
}

// unsigned int isWhispering(uint64  scHandlerID, anyID  clientID, int*  result)
unsigned int isWhispering(const struct TS3Functions funcs, uint64  scHandlerID, anyID  clientID, int*  result) {
	
	return funcs.isWhispering(scHandlerID, clientID, result);
}

// unsigned int isReceivingWhisper(uint64  scHandlerID, anyID  clientID, int*  result)
unsigned int isReceivingWhisper(const struct TS3Functions funcs, uint64  scHandlerID, anyID  clientID, int*  result) {
	
	return funcs.isReceivingWhisper(scHandlerID, clientID, result);
}

// unsigned int getAvatar(uint64  scHandlerID, anyID  clientID, char*  result, size_t  maxLen)
unsigned int getAvatar(const struct TS3Functions funcs, uint64  scHandlerID, anyID  clientID, char*  result, size_t  maxLen) {
	
	return funcs.getAvatar(scHandlerID, clientID, result, maxLen);
}

// void setPluginMenuEnabled(const char*  pluginID, int  menuID, int  enabled)
void setPluginMenuEnabled(const struct TS3Functions funcs, const char*  pluginID, int  menuID, int  enabled) {
	
	return funcs.setPluginMenuEnabled(pluginID, menuID, enabled);
}

// void requestHotkeyInputDialog(const char*  pluginID, const char*  keyword, int  isDown, void*  qParentWindow)
void requestHotkeyInputDialog(const struct TS3Functions funcs, const char*  pluginID, const char*  keyword, int  isDown, void*  qParentWindow) {
	
	return funcs.requestHotkeyInputDialog(pluginID, keyword, isDown, qParentWindow);
}

// unsigned int getHotkeyFromKeyword(const char*  pluginID, const char**  keywords, char**  hotkeys, size_t  arrayLen, size_t  hotkeyBufSize)
unsigned int getHotkeyFromKeyword(const struct TS3Functions funcs, const char*  pluginID, const char**  keywords, char**  hotkeys, size_t  arrayLen, size_t  hotkeyBufSize) {
	
	return funcs.getHotkeyFromKeyword(pluginID, keywords, hotkeys, arrayLen, hotkeyBufSize);
}

// unsigned int getClientDisplayName(uint64  scHandlerID, anyID  clientID, char*  result, size_t  maxLen)
unsigned int getClientDisplayName(const struct TS3Functions funcs, uint64  scHandlerID, anyID  clientID, char*  result, size_t  maxLen) {
	
	return funcs.getClientDisplayName(scHandlerID, clientID, result, maxLen);
}

// unsigned int getBookmarkList(struct PluginBookmarkList**  list)
unsigned int getBookmarkList(const struct TS3Functions funcs, struct PluginBookmarkList**  list) {
	
	return funcs.getBookmarkList(list);
}

// unsigned int getProfileList(enum PluginGuiProfile  profile, int*  defaultProfileIdx, char***  result)
unsigned int getProfileList(const struct TS3Functions funcs, enum PluginGuiProfile  profile, int*  defaultProfileIdx, char***  result) {
	
	return funcs.getProfileList(profile, defaultProfileIdx, result);
}

// unsigned int guiConnect(enum PluginConnectTab  connectTab, const char*  serverLabel, const char*  serverAddress, const char*  serverPassword, const char*  nickname, const char*  channel, const char*  channelPassword, const char*  captureProfile, const char*  playbackProfile, const char*  hotkeyProfile, const char*  soundProfile, const char*  userIdentity, const char*  oneTimeKey, const char*  phoneticName, uint64*  scHandlerID)
unsigned int guiConnect(const struct TS3Functions funcs, enum PluginConnectTab  connectTab, const char*  serverLabel, const char*  serverAddress, const char*  serverPassword, const char*  nickname, const char*  channel, const char*  channelPassword, const char*  captureProfile, const char*  playbackProfile, const char*  hotkeyProfile, const char*  soundProfile, const char*  userIdentity, const char*  oneTimeKey, const char*  phoneticName, uint64*  scHandlerID) {
	
	return funcs.guiConnect(connectTab, serverLabel, serverAddress, serverPassword, nickname, channel, channelPassword, captureProfile, playbackProfile, hotkeyProfile, soundProfile, userIdentity, oneTimeKey, phoneticName, scHandlerID);
}

// unsigned int guiConnectBookmark(enum PluginConnectTab  connectTab, const char*  bookmarkuuid, uint64*  scHandlerID)
unsigned int guiConnectBookmark(const struct TS3Functions funcs, enum PluginConnectTab  connectTab, const char*  bookmarkuuid, uint64*  scHandlerID) {
	
	return funcs.guiConnectBookmark(connectTab, bookmarkuuid, scHandlerID);
}

// unsigned int createBookmark(const char*  bookmarkuuid, const char*  serverLabel, const char*  serverAddress, const char*  serverPassword, const char*  nickname, const char*  channel, const char*  channelPassword, const char*  captureProfile, const char*  playbackProfile, const char*  hotkeyProfile, const char*  soundProfile, const char*  uniqueUserId, const char*  oneTimeKey, const char*  phoneticName)
unsigned int createBookmark(const struct TS3Functions funcs, const char*  bookmarkuuid, const char*  serverLabel, const char*  serverAddress, const char*  serverPassword, const char*  nickname, const char*  channel, const char*  channelPassword, const char*  captureProfile, const char*  playbackProfile, const char*  hotkeyProfile, const char*  soundProfile, const char*  uniqueUserId, const char*  oneTimeKey, const char*  phoneticName) {
	
	return funcs.createBookmark(bookmarkuuid, serverLabel, serverAddress, serverPassword, nickname, channel, channelPassword, captureProfile, playbackProfile, hotkeyProfile, soundProfile, uniqueUserId, oneTimeKey, phoneticName);
}

// unsigned int getPermissionIDByName(uint64  serverConnectionHandlerID, const char*  permissionName, unsigned int*  result)
unsigned int getPermissionIDByName(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  permissionName, unsigned int*  result) {
	
	return funcs.getPermissionIDByName(serverConnectionHandlerID, permissionName, result);
}

// unsigned int getClientNeededPermission(uint64  serverConnectionHandlerID, const char*  permissionName, int*  result)
unsigned int getClientNeededPermission(const struct TS3Functions funcs, uint64  serverConnectionHandlerID, const char*  permissionName, int*  result) {
	
	return funcs.getClientNeededPermission(serverConnectionHandlerID, permissionName, result);
}

// void notifyKeyEvent(const char * pluginID, const char * keyIdentifier, int  up_down)
void notifyKeyEvent(const struct TS3Functions funcs, const char * pluginID, const char * keyIdentifier, int  up_down) {
	
	return funcs.notifyKeyEvent(pluginID, keyIdentifier, up_down);
}

