package teamspeak

/*
#cgo CFLAGS: -I../pluginsdk/include -I.

#include <teamspeak/public_definitions.h>

// make #defines from public_definitions.h accessible to cgo
const int define_TS3_MAX_SIZE_CHANNEL_NAME                 = TS3_MAX_SIZE_CHANNEL_NAME;
const int define_TS3_MAX_SIZE_VIRTUALSERVER_NAME           = TS3_MAX_SIZE_VIRTUALSERVER_NAME;
const int define_TS3_MAX_SIZE_CLIENT_NICKNAME              = TS3_MAX_SIZE_CLIENT_NICKNAME;
const int define_TS3_MIN_SIZE_CLIENT_NICKNAME              = TS3_MIN_SIZE_CLIENT_NICKNAME;
const int define_TS3_MAX_SIZE_REASON_MESSAGE               = TS3_MAX_SIZE_REASON_MESSAGE;
const int define_TS3_MAX_SIZE_TEXTMESSAGE                  = TS3_MAX_SIZE_TEXTMESSAGE;
const int define_TS3_MAX_SIZE_CHANNEL_TOPIC                = TS3_MAX_SIZE_CHANNEL_TOPIC;
const int define_TS3_MAX_SIZE_CHANNEL_DESCRIPTION          = TS3_MAX_SIZE_CHANNEL_DESCRIPTION;
const int define_TS3_MAX_SIZE_VIRTUALSERVER_WELCOMEMESSAGE = TS3_MAX_SIZE_VIRTUALSERVER_WELCOMEMESSAGE;
const int define_TS3_MIN_SECONDS_CLIENTID_REUSE            = TS3_MIN_SECONDS_CLIENTID_REUSE;
const int define_MAX_VARIABLES_EXPORT_COUNT                = MAX_VARIABLES_EXPORT_COUNT;
const long long define_BANDWIDTH_LIMIT_UNLIMITED           = BANDWIDTH_LIMIT_UNLIMITED;
const unsigned long define_SPEAKER_FRONT_LEFT              = SPEAKER_FRONT_LEFT;
const unsigned long define_SPEAKER_FRONT_RIGHT             = SPEAKER_FRONT_RIGHT;
const unsigned long define_SPEAKER_FRONT_CENTER            = SPEAKER_FRONT_CENTER;
const unsigned long define_SPEAKER_LOW_FREQUENCY           = SPEAKER_LOW_FREQUENCY;
const unsigned long define_SPEAKER_BACK_LEFT               = SPEAKER_BACK_LEFT;
const unsigned long define_SPEAKER_BACK_RIGHT              = SPEAKER_BACK_RIGHT;
const unsigned long define_SPEAKER_FRONT_LEFT_OF_CENTER    = SPEAKER_FRONT_LEFT_OF_CENTER;
const unsigned long define_SPEAKER_FRONT_RIGHT_OF_CENTER   = SPEAKER_FRONT_RIGHT_OF_CENTER;
const unsigned long define_SPEAKER_BACK_CENTER             = SPEAKER_BACK_CENTER;
const unsigned long define_SPEAKER_SIDE_LEFT               = SPEAKER_SIDE_LEFT;
const unsigned long define_SPEAKER_SIDE_RIGHT              = SPEAKER_SIDE_RIGHT;
const unsigned long define_SPEAKER_TOP_CENTER              = SPEAKER_TOP_CENTER;
const unsigned long define_SPEAKER_TOP_FRONT_LEFT          = SPEAKER_TOP_FRONT_LEFT;
const unsigned long define_SPEAKER_TOP_FRONT_CENTER        = SPEAKER_TOP_FRONT_CENTER;
const unsigned long define_SPEAKER_TOP_FRONT_RIGHT         = SPEAKER_TOP_FRONT_RIGHT;
const unsigned long define_SPEAKER_TOP_BACK_LEFT           = SPEAKER_TOP_BACK_LEFT;
const unsigned long define_SPEAKER_TOP_BACK_CENTER         = SPEAKER_TOP_BACK_CENTER;
const unsigned long define_SPEAKER_TOP_BACK_RIGHT          = SPEAKER_TOP_BACK_RIGHT;
const unsigned long define_SPEAKER_HEADPHONES_LEFT         = SPEAKER_HEADPHONES_LEFT;
const unsigned long define_SPEAKER_HEADPHONES_RIGHT        = SPEAKER_HEADPHONES_RIGHT;
const unsigned long define_SPEAKER_MONO                    = SPEAKER_MONO;

*/
import "C"

/******************************************************************************
From public_definitions.h
*******************************************************************************/

var (
	//limited length, measured in characters
	MaxSizeChannelName = int(C.define_TS3_MAX_SIZE_CHANNEL_NAME)

	//limited length, measured in characters
	MaxSizeVirtualServerName = int(C.define_TS3_MAX_SIZE_VIRTUALSERVER_NAME)

	//limited length, measured in characters
	MaxSizeClientNickname = int(C.define_TS3_MAX_SIZE_CLIENT_NICKNAME)

	//limited length, measured in characters
	MinSizeClientNickname = int(C.define_TS3_MIN_SIZE_CLIENT_NICKNAME)

	//limited length, measured in characters
	MaxSizeReasonMessage = int(C.define_TS3_MAX_SIZE_REASON_MESSAGE)

	//limited length, measured in bytes (utf8 encoded)
	MaxSizeTextMessage = int(C.define_TS3_MAX_SIZE_TEXTMESSAGE)

	//limited length, measured in bytes (utf8 encoded)
	MaxSizeChannelTopic = int(C.define_TS3_MAX_SIZE_CHANNEL_TOPIC)

	//limited length, measured in bytes (utf8 encoded)
	MaxSizeChannelDescription = int(C.define_TS3_MAX_SIZE_CHANNEL_DESCRIPTION)

	//limited length, measured in bytes (utf8 encoded)
	MaxSizeVirtualServerWelcomeMessage = int(C.define_TS3_MAX_SIZE_VIRTUALSERVER_WELCOMEMESSAGE)

	//minimum amount of seconds before a clientID that was in use can be assigned to a new client
	MinSecondsClientIDReuse = int(C.define_TS3_MIN_SECONDS_CLIENTID_REUSE)
)

type AnyID uint16

func (value AnyID) toC() C.anyID {
	return C.anyID(value)
}

type TalkStatus uint8

func (value TalkStatus) toC() C.enum_TalkStatus {
	return C.enum_TalkStatus(value)
}

const (
	StatusNotTalking           = TalkStatus(C.STATUS_NOT_TALKING)
	StatusTalking              = TalkStatus(C.STATUS_TALKING)
	StatusTalkingWhileDisabled = TalkStatus(C.STATUS_TALKING_WHILE_DISABLED)
)

type CodecType uint8

func (value CodecType) toC() C.enum_CodecType {
	return C.enum_CodecType(value)
}

const (
	//mono,   16bit,  8kHz, bitrate dependent on the quality setting
	CodecSpeexNarrowBand = CodecType(C.CODEC_SPEEX_NARROWBAND)

	//mono,   16bit, 16kHz, bitrate dependent on the quality setting
	CodecSpeexWideBand = CodecType(C.CODEC_SPEEX_WIDEBAND)

	//mono,   16bit, 32kHz, bitrate dependent on the quality setting
	CodecSpeexUltraWideBand = CodecType(C.CODEC_SPEEX_ULTRAWIDEBAND)

	//mono,   16bit, 48kHz, bitrate dependent on the quality setting
	CodecCeltMono = CodecType(C.CODEC_CELT_MONO)

	//mono,   16bit, 48khz, bitrate dependent on the quality setting, optimized for voice
	CodecOpusVoice = CodecType(C.CODEC_OPUS_VOICE)

	//stereo, 16bit, 48khz, bitrate dependent on the quality setting, optimized for music
	CodecOpusMusic = CodecType(C.CODEC_OPUS_MUSIC)
)

type CodecEncryptionMode uint8

func (value CodecEncryptionMode) toC() C.enum_CodecEncryptionMode {
	return C.enum_CodecEncryptionMode(value)
}

const (
	CodecEncryptionPerChannel = CodecEncryptionMode(C.CODEC_ENCRYPTION_PER_CHANNEL)
	CodecEncryptionForcedOff  = CodecEncryptionMode(C.CODEC_ENCRYPTION_FORCED_OFF)
	CodecEncryptionForcedOn   = CodecEncryptionMode(C.CODEC_ENCRYPTION_FORCED_ON)
)

type TextMessageTargetMode uint8

func (value TextMessageTargetMode) toC() C.enum_TextMessageTargetMode {
	return C.enum_TextMessageTargetMode(value)
}

const (
	TextMessageTargetClient  = TextMessageTargetMode(C.TextMessageTarget_CLIENT)
	TextMessageTargetChannel = TextMessageTargetMode(C.TextMessageTarget_CHANNEL)
	TextMessageTargetServer  = TextMessageTargetMode(C.TextMessageTarget_SERVER)
	TextMessageTargetMax     = TextMessageTargetMode(C.TextMessageTarget_MAX)
)

type MuteInputStatus uint8

func (value MuteInputStatus) toC() C.enum_MuteInputStatus {
	return C.enum_MuteInputStatus(value)
}

const (
	MuteInputNone  = MuteInputStatus(C.MUTEINPUT_NONE)
	MuteInputMuted = MuteInputStatus(C.MUTEINPUT_MUTED)
)

type MuteOutputStatus uint8

func (value MuteOutputStatus) toC() C.enum_MuteOutputStatus {
	return C.enum_MuteOutputStatus(value)
}

const (
	MuteOutputNone  = MuteOutputStatus(C.MUTEOUTPUT_NONE)
	MuteOutputMuted = MuteOutputStatus(C.MUTEOUTPUT_MUTED)
)

type HardwareInputStatus uint8

func (value HardwareInputStatus) toC() C.enum_HardwareInputStatus {
	return C.enum_HardwareInputStatus(value)
}

const (
	HardwareInputDisabled = HardwareInputStatus(C.HARDWAREINPUT_DISABLED)
	HardwareInputEnabled  = HardwareInputStatus(C.HARDWAREINPUT_ENABLED)
)

type HardwareOutputStatus uint8

func (value HardwareOutputStatus) toC() C.enum_HardwareOutputStatus {
	return C.enum_HardwareOutputStatus(value)
}

const (
	HardwareOutputDisabled = HardwareOutputStatus(C.HARDWAREOUTPUT_DISABLED)
	HardwareOutputEnabled  = HardwareOutputStatus(C.HARDWAREOUTPUT_ENABLED)
)

type InputDeactivationStatus uint8

func (value InputDeactivationStatus) toC() C.enum_InputDeactivationStatus {
	return C.enum_InputDeactivationStatus(value)
}

const (
	InputActive      = InputDeactivationStatus(C.INPUT_ACTIVE)
	InputDeactivated = InputDeactivationStatus(C.INPUT_DEACTIVATED)
)

type ReasonIdentifier uint8

func (value ReasonIdentifier) toC() C.enum_ReasonIdentifier {
	return C.enum_ReasonIdentifier(value)
}

const (
	//no reason data
	ReasonNone = ReasonIdentifier(C.REASON_NONE)

	//{SectionInvoker}
	ReasonMoved = ReasonIdentifier(C.REASON_MOVED)

	//no reason data
	ReasonSubscription = ReasonIdentifier(C.REASON_SUBSCRIPTION)

	//reasonmsg=reason
	ReasonLostConnection = ReasonIdentifier(C.REASON_LOST_CONNECTION)

	//{SectionInvoker} reasonmsg=reason               //{SectionInvoker} is only added server->client
	ReasonKickChannel = ReasonIdentifier(C.REASON_KICK_CHANNEL)

	//{SectionInvoker} reasonmsg=reason               //{SectionInvoker} is only added server->client
	ReasonKickServer = ReasonIdentifier(C.REASON_KICK_SERVER)

	//{SectionInvoker} reasonmsg=reason bantime=time  //{SectionInvoker} is only added server->client
	ReasonKickServerBan = ReasonIdentifier(C.REASON_KICK_SERVER_BAN)

	//reasonmsg=reason
	ReasonServerStop = ReasonIdentifier(C.REASON_SERVERSTOP)

	//reasonmsg=reason
	ReasonClientDisconnect = ReasonIdentifier(C.REASON_CLIENTDISCONNECT)

	//no reason data
	ReasonChannelUpdate = ReasonIdentifier(C.REASON_CHANNELUPDATE)

	//{SectionInvoker}
	ReasonChannelEdit = ReasonIdentifier(C.REASON_CHANNELEDIT)

	//reasonmsg=reason
	ReasonClientdisconnectServerShutdown = ReasonIdentifier(C.REASON_CLIENTDISCONNECT_SERVER_SHUTDOWN)
)

type ChannelProperty uint8

func (value ChannelProperty) toC() C.enum_ChannelProperties {
	return C.enum_ChannelProperties(value)
}

const (
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyName = ChannelProperty(C.CHANNEL_NAME)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyTopic = ChannelProperty(C.CHANNEL_TOPIC)

	//Must be requested (=> requestChannelDescription)
	ChannelPropertyDescription = ChannelProperty(C.CHANNEL_DESCRIPTION)

	//not available client side
	ChannelPropertyPassword = ChannelProperty(C.CHANNEL_PASSWORD)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyCodec = ChannelProperty(C.CHANNEL_CODEC)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyCodecQuality = ChannelProperty(C.CHANNEL_CODEC_QUALITY)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyMaxClients = ChannelProperty(C.CHANNEL_MAXCLIENTS)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyMaxFamilyClients = ChannelProperty(C.CHANNEL_MAXFAMILYCLIENTS)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyOrder = ChannelProperty(C.CHANNEL_ORDER)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagPermanent = ChannelProperty(C.CHANNEL_FLAG_PERMANENT)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagSemiPermanent = ChannelProperty(C.CHANNEL_FLAG_SEMI_PERMANENT)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagDefault = ChannelProperty(C.CHANNEL_FLAG_DEFAULT)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagPassword = ChannelProperty(C.CHANNEL_FLAG_PASSWORD)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyCodecLatencyFactor = ChannelProperty(C.CHANNEL_CODEC_LATENCY_FACTOR)

	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyCodecIsUnencrypted = ChannelProperty(C.CHANNEL_CODEC_IS_UNENCRYPTED)

	//Not available client side, not used in teamspeak, only SDK. Sets the options+salt for security hash.
	ChannelPropertySecuritySalt = ChannelProperty(C.CHANNEL_SECURITY_SALT)

	//How many seconds to wait before deleting this channel
	ChannelPropertyDeleteDelay = ChannelProperty(C.CHANNEL_DELETE_DELAY)

	ChannelPropertyEndMarker = ChannelProperty(C.CHANNEL_ENDMARKER)
)

type ClientProperty uint8

func (value ClientProperty) toC() C.enum_ClientProperties {
	return C.enum_ClientProperties(value)
}

const (
	//automatically up-to-date for any client "in view", can be used to identify this particular client installation
	ClientPropertyUniqueIdentifier = ClientProperty(C.CLIENT_UNIQUE_IDENTIFIER)

	//automatically up-to-date for any client "in view"
	ClientPropertyNickname = ClientProperty(C.CLIENT_NICKNAME)

	//for other clients than ourself, this needs to be requested (=> requestClientVariables)
	ClientPropertyVersion = ClientProperty(C.CLIENT_VERSION)

	//for other clients than ourself, this needs to be requested (=> requestClientVariables)
	ClientPropertyPlatform = ClientProperty(C.CLIENT_PLATFORM)

	//automatically up-to-date for any client that can be heard (in room / whisper)
	ClientPropertyFlagTalking = ClientProperty(C.CLIENT_FLAG_TALKING)

	//automatically up-to-date for any client "in view", this clients microphone mute status
	ClientPropertyInputMuted = ClientProperty(C.CLIENT_INPUT_MUTED)

	//automatically up-to-date for any client "in view", this clients headphones/speakers/mic combined mute status
	ClientPropertyOutputMuted = ClientProperty(C.CLIENT_OUTPUT_MUTED)

	//automatically up-to-date for any client "in view", this clients headphones/speakers only mute status
	ClientPropertyOutputOnlyMuted = ClientProperty(C.CLIENT_OUTPUTONLY_MUTED)

	//automatically up-to-date for any client "in view", this clients microphone hardware status (is the capture device opened?)
	ClientPropertyInputHardware = ClientProperty(C.CLIENT_INPUT_HARDWARE)

	//automatically up-to-date for any client "in view", this clients headphone/speakers hardware status (is the playback device opened?)
	ClientPropertyOutputHardware = ClientProperty(C.CLIENT_OUTPUT_HARDWARE)

	//only usable for ourself, not propagated to the network
	ClientPropertyInputDeactivated = ClientProperty(C.CLIENT_INPUT_DEACTIVATED)

	//internal use
	ClientPropertyIdleTime = ClientProperty(C.CLIENT_IDLE_TIME)

	//only usable for ourself, the default channel we used to connect on our last connection attempt
	ClientPropertyDefaultChannel = ClientProperty(C.CLIENT_DEFAULT_CHANNEL)

	//internal use
	ClientPropertyDefaultChannelPassword = ClientProperty(C.CLIENT_DEFAULT_CHANNEL_PASSWORD)

	//internal use
	ClientPropertyServerPassword = ClientProperty(C.CLIENT_SERVER_PASSWORD)

	//automatically up-to-date for any client "in view", not used by TeamSpeak, free storage for sdk users
	ClientPropertyMetaData = ClientProperty(C.CLIENT_META_DATA)

	//only make sense on the client side locally, "1" if this client is currently muted by us, "0" if he is not
	ClientPropertyIsMuted = ClientProperty(C.CLIENT_IS_MUTED)

	//automatically up-to-date for any client "in view"
	ClientPropertyIsRecording = ClientProperty(C.CLIENT_IS_RECORDING)

	//internal use
	ClientPropertyVolumeModificator = ClientProperty(C.CLIENT_VOLUME_MODIFICATOR)

	//sign
	ClientPropertyVersionSign = ClientProperty(C.CLIENT_VERSION_SIGN)

	//SDK use, not used by teamspeak. Hash is provided by an outside source. A channel will use the security salt + other client data to calculate a hash, which must be the same as the one provided here.
	ClientPropertySecurityHash = ClientProperty(C.CLIENT_SECURITY_HASH)

	ClientPropertyEndMarker = ClientProperty(C.CLIENT_ENDMARKER)
)

type VirtualServerProperty uint8

func (value VirtualServerProperty) toC() C.enum_VirtualServerProperties {
	return C.enum_VirtualServerProperties(value)
}

const (
	//available when connected, can be used to identify this particular server installation
	VirtualServerPropertyUniqueIdentifier = VirtualServerProperty(C.VIRTUALSERVER_UNIQUE_IDENTIFIER)

	//available and always up-to-date when connected
	VirtualServerPropertyName = VirtualServerProperty(C.VIRTUALSERVER_NAME)

	//available when connected,  (=> requestServerVariables)
	VirtualServerPropertyWelcomeMessage = VirtualServerProperty(C.VIRTUALSERVER_WELCOMEMESSAGE)

	//available when connected
	VirtualServerPropertyPlatform = VirtualServerProperty(C.VIRTUALSERVER_PLATFORM)

	//available when connected
	VirtualServerPropertyVersion = VirtualServerProperty(C.VIRTUALSERVER_VERSION)

	//only available on request (=> requestServerVariables), stores the maximum number of clients that may currently join the server
	VirtualServerPropertyMaxClients = VirtualServerProperty(C.VIRTUALSERVER_MAXCLIENTS)

	//not available to clients, the server password
	VirtualServerPropertyPassword = VirtualServerProperty(C.VIRTUALSERVER_PASSWORD)

	//only available on request (=> requestServerVariables),
	VirtualServerPropertyClientsOnline = VirtualServerProperty(C.VIRTUALSERVER_CLIENTS_ONLINE)

	//only available on request (=> requestServerVariables),
	VirtualServerPropertyChannelsOnline = VirtualServerProperty(C.VIRTUALSERVER_CHANNELS_ONLINE)

	//available when connected, stores the time when the server was created
	VirtualServerPropertyCreated = VirtualServerProperty(C.VIRTUALSERVER_CREATED)

	//only available on request (=> requestServerVariables), the time since the server was started
	VirtualServerPropertyUptime = VirtualServerProperty(C.VIRTUALSERVER_UPTIME)

	//available and always up-to-date when connected
	VirtualServerPropertyCodecEncryptionMode = VirtualServerProperty(C.VIRTUALSERVER_CODEC_ENCRYPTION_MODE)

	VirtualServerPropertyEndMarker = VirtualServerProperty(C.VIRTUALSERVER_ENDMARKER)
)

type ConnectionProperty uint8

func (value ConnectionProperty) toC() C.enum_ConnectionProperties {
	return C.enum_ConnectionProperties(value)
}

const (
	//average latency for a round trip through and back this connection
	ConnectionPropertyPing = ConnectionProperty(C.CONNECTION_PING)

	//standard deviation of the above average latency
	ConnectionPropertyPingDeviation = ConnectionProperty(C.CONNECTION_PING_DEVIATION)

	//how long the connection exists already
	ConnectionPropertyConnectedTime = ConnectionProperty(C.CONNECTION_CONNECTED_TIME)

	//how long since the last action of this client
	ConnectionPropertyIdleTime = ConnectionProperty(C.CONNECTION_IDLE_TIME)

	//IP of this client (as seen from the server side)
	ConnectionPropertyClientIP = ConnectionProperty(C.CONNECTION_CLIENT_IP)

	//Port of this client (as seen from the server side)
	ConnectionPropertyClientPort = ConnectionProperty(C.CONNECTION_CLIENT_PORT)

	//IP of the server (seen from the client side) - only available on yourself, not for remote clients, not available server side
	ConnectionPropertyServerIP = ConnectionProperty(C.CONNECTION_SERVER_IP)

	//Port of the server (seen from the client side) - only available on yourself, not for remote clients, not available server side
	ConnectionPropertyServerPort = ConnectionProperty(C.CONNECTION_SERVER_PORT)

	//how many Speech packets were sent through this connection
	ConnectionPropertyPacketsSentSpeech = ConnectionProperty(C.CONNECTION_PACKETS_SENT_SPEECH)

	ConnectionPropertyPacketsSentKeepAlive = ConnectionProperty(C.CONNECTION_PACKETS_SENT_KEEPALIVE)

	ConnectionPropertyPacketsSentControl = ConnectionProperty(C.CONNECTION_PACKETS_SENT_CONTROL)

	//how many packets were sent totally (this is PACKETS_SENT_SPEECH + PACKETS_SENT_KEEPALIVE + PACKETS_SENT_CONTROL)
	ConnectionPropertyPacketsSentTotal = ConnectionProperty(C.CONNECTION_PACKETS_SENT_TOTAL)

	ConnectionPropertyBytesSentSpeech = ConnectionProperty(C.CONNECTION_BYTES_SENT_SPEECH)

	ConnectionPropertyBytesSentKeepAlive = ConnectionProperty(C.CONNECTION_BYTES_SENT_KEEPALIVE)

	ConnectionPropertyBytesSentControl = ConnectionProperty(C.CONNECTION_BYTES_SENT_CONTROL)

	ConnectionPropertyBytesSentTotal = ConnectionProperty(C.CONNECTION_BYTES_SENT_TOTAL)

	ConnectionPropertyPacketsReceivedSpeech = ConnectionProperty(C.CONNECTION_PACKETS_RECEIVED_SPEECH)

	ConnectionPropertyPacketsReceivedKeepAlive = ConnectionProperty(C.CONNECTION_PACKETS_RECEIVED_KEEPALIVE)

	ConnectionPropertyPacketsReceivedControl = ConnectionProperty(C.CONNECTION_PACKETS_RECEIVED_CONTROL)

	ConnectionPropertyPacketsReceivedTotal = ConnectionProperty(C.CONNECTION_PACKETS_RECEIVED_TOTAL)

	ConnectionPropertyBytesReceivedSpeech = ConnectionProperty(C.CONNECTION_BYTES_RECEIVED_SPEECH)

	ConnectionPropertyBytesReceivedKeepAlive = ConnectionProperty(C.CONNECTION_BYTES_RECEIVED_KEEPALIVE)

	ConnectionPropertyBytesReceivedControl = ConnectionProperty(C.CONNECTION_BYTES_RECEIVED_CONTROL)

	ConnectionPropertyBytesReceivedTotal = ConnectionProperty(C.CONNECTION_BYTES_RECEIVED_TOTAL)

	ConnectionPropertyPacketLossSpeech = ConnectionProperty(C.CONNECTION_PACKETLOSS_SPEECH)

	ConnectionPropertyPacketLossKeepAlive = ConnectionProperty(C.CONNECTION_PACKETLOSS_KEEPALIVE)

	ConnectionPropertyPacketLossControl = ConnectionProperty(C.CONNECTION_PACKETLOSS_CONTROL)

	//the probability with which a packet round trip failed because a packet was lost
	ConnectionPropertyPacketLossTotal = ConnectionProperty(C.CONNECTION_PACKETLOSS_TOTAL)

	//the probability with which a speech packet failed from the server to the client
	ConnectionPropertyServer2ClientPacketLossSpeech = ConnectionProperty(C.CONNECTION_SERVER2CLIENT_PACKETLOSS_SPEECH)

	ConnectionPropertyServer2ClientPacketLossKeepAlive = ConnectionProperty(C.CONNECTION_SERVER2CLIENT_PACKETLOSS_KEEPALIVE)

	ConnectionPropertyServer2ClientPacketLossControl = ConnectionProperty(C.CONNECTION_SERVER2CLIENT_PACKETLOSS_CONTROL)

	ConnectionPropertyServer2ClientPacketLossTotal = ConnectionProperty(C.CONNECTION_SERVER2CLIENT_PACKETLOSS_TOTAL)

	ConnectionPropertyClient2ServerPacketLossSpeech = ConnectionProperty(C.CONNECTION_CLIENT2SERVER_PACKETLOSS_SPEECH)

	ConnectionPropertyClient2ServerPacketLossKeepAlive = ConnectionProperty(C.CONNECTION_CLIENT2SERVER_PACKETLOSS_KEEPALIVE)

	ConnectionPropertyClient2ServerPacketLossControl = ConnectionProperty(C.CONNECTION_CLIENT2SERVER_PACKETLOSS_CONTROL)

	ConnectionPropertyClient2ServerPacketLossTotal = ConnectionProperty(C.CONNECTION_CLIENT2SERVER_PACKETLOSS_TOTAL)

	//howmany bytes of speech packets we sent during the last second
	ConnectionPropertyBandwidthSentLastSecondSpeech = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_SECOND_SPEECH)

	ConnectionPropertyBandwidthSentLastSecondKeepAlive = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_SECOND_KEEPALIVE)

	ConnectionPropertyBandwidthSentLastSecondControl = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_SECOND_CONTROL)

	ConnectionPropertyBandwidthSentLastSecondTotal = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_SECOND_TOTAL)

	//howmany bytes/s of speech packets we sent in average during the last minute
	ConnectionPropertyBandwidthSentLastMinuteSpeech = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_MINUTE_SPEECH)

	ConnectionPropertyBandwidthSentLastMinuteKeepAlive = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_MINUTE_KEEPALIVE)

	ConnectionPropertyBandwidthSentLastMinuteControl = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_MINUTE_CONTROL)

	ConnectionPropertyBandwidthSentLastMinuteTotal = ConnectionProperty(C.CONNECTION_BANDWIDTH_SENT_LAST_MINUTE_TOTAL)

	ConnectionPropertyBandwidthReceivedLastSecondSpeech = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_SECOND_SPEECH)

	ConnectionPropertyBandwidthReceivedLastSecondKeepAlive = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_SECOND_KEEPALIVE)

	ConnectionPropertyBandwidthReceivedLastSecondControl = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_SECOND_CONTROL)

	ConnectionPropertyBandwidthReceivedLastSecondTotal = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_SECOND_TOTAL)

	ConnectionPropertyBandwidthReceivedLastMinuteSpeech = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_MINUTE_SPEECH)

	ConnectionPropertyBandwidthReceivedLastMinuteKeepAlive = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_MINUTE_KEEPALIVE)

	ConnectionPropertyBandwidthReceivedLastMinuteControl = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_MINUTE_CONTROL)

	ConnectionPropertyBandwidthReceivedLastMinuteTotal = ConnectionProperty(C.CONNECTION_BANDWIDTH_RECEIVED_LAST_MINUTE_TOTAL)

	ConnectionPropertyEndMarker = ConnectionProperty(C.CONNECTION_ENDMARKER)
)

type Vector struct {
	X, Y, Z float32
}

// TODO
func (v Vector) toC() C.TS3_VECTOR {
	return C.TS3_VECTOR{
		x: C.float(v.X),
		y: C.float(v.Y),
		z: C.float(v.Z),
	}
}

type GroupWhisperType uint8

func (value GroupWhisperType) toC() C.enum_GroupWhisperType {
	return C.enum_GroupWhisperType(value)
}

const (
	GroupWhisperTypeServerGroup      = GroupWhisperType(C.GROUPWHISPERTYPE_SERVERGROUP)
	GroupWhisperTypeChannelGroup     = GroupWhisperType(C.GROUPWHISPERTYPE_CHANNELGROUP)
	GroupWhisperTypeChannelCommander = GroupWhisperType(C.GROUPWHISPERTYPE_CHANNELCOMMANDER)
	GroupWhisperTypeAllClients       = GroupWhisperType(C.GROUPWHISPERTYPE_ALLCLIENTS)
	GroupWhisperTypeEndMarker        = GroupWhisperType(C.GROUPWHISPERTYPE_ENDMARKER)
)

type GroupWhisperTargetMode uint8

func (value GroupWhisperTargetMode) toC() C.enum_GroupWhisperTargetMode {
	return C.enum_GroupWhisperTargetMode(value)
}

const (
	GroupWhisperTargetModeAll                   = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_ALL)
	GroupWhisperTargetModeCurrentChannel        = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_CURRENTCHANNEL)
	GroupWhisperTargetModeParentChannel         = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_PARENTCHANNEL)
	GroupWhisperTargetModeAllParentChannels     = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_ALLPARENTCHANNELS)
	GroupWhisperTargetModeChannelFamily         = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_CHANNELFAMILY)
	GroupWhisperTargetModeAncestorChannelFamily = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_ANCESTORCHANNELFAMILY)
	GroupWhisperTargetModeSubchannels           = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_SUBCHANNELS)
	GroupWhisperTargetModeEndMarker             = GroupWhisperTargetMode(C.GROUPWHISPERTARGETMODE_ENDMARKER)
)

type MonoSoundDestination uint8

func (value MonoSoundDestination) toC() C.enum_MonoSoundDestination {
	return C.enum_MonoSoundDestination(value)
}

const (
	/* Send mono sound to all available speakers */
	MonoSoundDestinationAll = MonoSoundDestination(C.MONO_SOUND_DESTINATION_ALL)
	/* Send mono sound to front center speaker if available */
	MonoSoundDestinationFrontCenter = MonoSoundDestination(C.MONO_SOUND_DESTINATION_FRONT_CENTER)
	/* Send mono sound to front left/right speakers if available */
	MonoSoundDestinationFrontLeftAndRight = MonoSoundDestination(C.MONO_SOUND_DESTINATION_FRONT_LEFT_AND_RIGHT)
)

type SecuritySaltOption uint8

func (value SecuritySaltOption) toC() C.enum_SecuritySaltOptions {
	return C.enum_SecuritySaltOptions(value)
}

const (
	/* put nickname into security hash */
	SecuritySaltCheckNickname = SecuritySaltOption(C.SECURITY_SALT_CHECK_NICKNAME)
	/* put (game)meta data into security hash */
	SecuritySaltCheckMetaData = SecuritySaltOption(C.SECURITY_SALT_CHECK_META_DATA)
)

/*this enum is used to disable client commands on the server*/
type ClientCommand uint8

func (value ClientCommand) toC() C.enum_ClientCommand {
	return C.enum_ClientCommand(value)
}

const (
	ClientCommandRequestConnectionInfo        = ClientCommand(C.CLIENT_COMMAND_requestConnectionInfo)
	ClientCommandRequestClientMove            = ClientCommand(C.CLIENT_COMMAND_requestClientMove)
	ClientCommandRequestXXMuteClients         = ClientCommand(C.CLIENT_COMMAND_requestXXMuteClients)
	ClientCommandRequestClientKickFromXXX     = ClientCommand(C.CLIENT_COMMAND_requestClientKickFromXXX)
	ClientCommandFlushChannelCreation         = ClientCommand(C.CLIENT_COMMAND_flushChannelCreation)
	ClientCommandFlushChannelUpdates          = ClientCommand(C.CLIENT_COMMAND_flushChannelUpdates)
	ClientCommandRequestChannelMove           = ClientCommand(C.CLIENT_COMMAND_requestChannelMove)
	ClientCommandRequestChannelDelete         = ClientCommand(C.CLIENT_COMMAND_requestChannelDelete)
	ClientCommandRequestChannelDescription    = ClientCommand(C.CLIENT_COMMAND_requestChannelDescription)
	ClientCommandRequestChannelXXSubscribeXXX = ClientCommand(C.CLIENT_COMMAND_requestChannelXXSubscribeXXX)
	ClientCommandRequestServerConnectionInfo  = ClientCommand(C.CLIENT_COMMAND_requestServerConnectionInfo)
	ClientCommandRequestSendXXXTextMsg        = ClientCommand(C.CLIENT_COMMAND_requestSendXXXTextMsg)
	ClientCommandFileTransfers                = ClientCommand(C.CLIENT_COMMAND_filetransfers)
	ClientCommandEndMarker                    = ClientCommand(C.CLIENT_COMMAND_ENDMARKER)
)

/* Access Control List*/
type ACLType uint8

func (value ACLType) toC() C.enum_ACLType {
	return C.enum_ACLType(value)
}

const (
	ACLNone      = ACLType(C.ACL_NONE)
	ACLWhiteList = ACLType(C.ACL_WHITE_LIST)
	ACLBlackList = ACLType(C.ACL_BLACK_LIST)
)

/* file transfer actions*/
type FTAction uint8

func (value FTAction) toC() C.enum_FTAction {
	return C.enum_FTAction(value)
}

const (
	FTInitServer  = FTAction(C.FT_INIT_SERVER)
	FTInitChannel = FTAction(C.FT_INIT_CHANNEL)
	FTUpload      = FTAction(C.FT_UPLOAD)
	FTDownload    = FTAction(C.FT_DOWNLOAD)
	FTDelete      = FTAction(C.FT_DELETE)
	FTCreateDir   = FTAction(C.FT_CREATEDIR)
	FTRename      = FTAction(C.FT_RENAME)
	FTFileList    = FTAction(C.FT_FILELIST)
	FTFileInfo    = FTAction(C.FT_FILEINFO)
)

/* file transfer status */
type FileTransferState uint8

func (value FileTransferState) toC() C.enum_FileTransferState {
	return C.enum_FileTransferState(value)
}

const (
	FileTransferInitialising = FileTransferState(C.FILETRANSFER_INITIALISING)
	FileTransferActive       = FileTransferState(C.FILETRANSFER_ACTIVE)
	FileTransferFinished     = FileTransferState(C.FILETRANSFER_FINISHED)
)

/* file transfer types */
type FileListType uint8

func (value FileListType) toC() C.int {
	return C.int(value)
}

const (
	FileListTypeDirectory = FileListType(C.FileListType_Directory)
	FileListTypeFile      = FileListType(C.FileListType_File)
)

var (
	MaxVariablesExportCount = int(C.define_MAX_VARIABLES_EXPORT_COUNT)
)

// TODO - struct VariablesExportItem
// type VariablesExportItem struct {
// 	itemIsValid uint8
// 	proposedIsSet unsafe.Pointer
// 	current unsafe.Pointer
// 	Proposed unsafe.Pointer
// }
//
// func (value VariablesExportItem) toC() C.struct_VariablesExportItem {
// 	var itemIsValid uint8 = 0
// 	if (value.ItemIsValid) { itemIsValid = 1}
//
// 	var proposedIsSet uint8 = 0
// 	if (value.ProposedIsSet) { proposedIsSet = 1}
//
// 	return C.struct_VariablesExportItem{
// 		itemIsValid: C.uchar(itemIsValid),
// 		proposedIsSet: C.uchar(proposedIsSet),
// 		current: C.
// 	}
// }

// TODO - struct VariablesExport
// TODO - struct ClientMiniExport
// TODO - struct TransformFilePathExport
// TODO - struct TransformFilePathExportReturns
// TODO - struct FileTransferCallbackExport

/*define for file transfer bandwith limits*/
var (
	BandwidthLimitUnlimited = int64(C.define_BANDWIDTH_LIMIT_UNLIMITED)
)

/*defines for speaker locations used by some sound callbacks*/
var (
	SpeakerFrontLeft          = uint32(C.define_SPEAKER_FRONT_LEFT)
	SpeakerFrontRight         = uint32(C.define_SPEAKER_FRONT_RIGHT)
	SpeakerFrontCenter        = uint32(C.define_SPEAKER_FRONT_CENTER)
	SpeakerLowFrequency       = uint32(C.define_SPEAKER_LOW_FREQUENCY)
	SpeakerBackLeft           = uint32(C.define_SPEAKER_BACK_LEFT)
	SpeakerBackRight          = uint32(C.define_SPEAKER_BACK_RIGHT)
	SpeakerFrontLeftOfCenter  = uint32(C.define_SPEAKER_FRONT_LEFT_OF_CENTER)
	SpeakerFrontRightOfCenter = uint32(C.define_SPEAKER_FRONT_RIGHT_OF_CENTER)
	SpeakerBackCenter         = uint32(C.define_SPEAKER_BACK_CENTER)
	SpeakerSideLeft           = uint32(C.define_SPEAKER_SIDE_LEFT)
	SpeakerSideRight          = uint32(C.define_SPEAKER_SIDE_RIGHT)
	SpeakerTopCenter          = uint32(C.define_SPEAKER_TOP_CENTER)
	SpeakerTopFrontLeft       = uint32(C.define_SPEAKER_TOP_FRONT_LEFT)
	SpeakerTopFrontCenter     = uint32(C.define_SPEAKER_TOP_FRONT_CENTER)
	SpeakerTopFrontRight      = uint32(C.define_SPEAKER_TOP_FRONT_RIGHT)
	SpeakerTopBackLeft        = uint32(C.define_SPEAKER_TOP_BACK_LEFT)
	SpeakerTopBackCenter      = uint32(C.define_SPEAKER_TOP_BACK_CENTER)
	SpeakerTopBackRight       = uint32(C.define_SPEAKER_TOP_BACK_RIGHT)
	SpeakerHeadphonesLeft     = uint32(C.define_SPEAKER_HEADPHONES_LEFT)
	SpeakerHeadphonesRight    = uint32(C.define_SPEAKER_HEADPHONES_RIGHT)
	SpeakerMono               = uint32(C.define_SPEAKER_MONO)
)
