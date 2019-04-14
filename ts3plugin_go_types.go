package ts3plugin

import (
	"time"

	"github.com/icedream/go-ts3plugin/teamlog"
	"github.com/icedream/go-ts3plugin/teamspeak"
)

var (
	Name, Version, Author, Description string
	ApiVersion                         int = 24
)

var (
	CommandKeyword string
	InfoTitle      string
)

var (
	RequestAutoload = false
)

var (
	// Will be set by TeamSpeak after plugin is loaded.
	pluginID string
)

var (
	Init                           func() (ok bool)
	Shutdown                       func()
	OffersConfigure                func() PluginConfigureOffer
	Configure                      func(handle byte, qParentWidget byte)
	ProcessCommand                 func(serverConnectionHandlerID uint64, command string) (handled bool)
	CurrentServerConnectionChanged func(serverConnectionHandlerID uint64)

	OnConnectedStatusChangeEvent      func(serverConnectionHandlerID uint64, newStatus int, errorNumber uint)
	OnNewChannelEvent                 func(serverConnectionHandlerID uint64, channelID uint64, channelParentID uint64)
	OnUpdateChannelEvent              func(serverConnectionHandlerID uint64, channelID uint64)
	OnClientIDsFinishedEvent          func(serverConnectionHandlerID uint64)
	OnServerUpdatedEvent              func(serverConnectionHandlerID uint64)
	OnUserLoggingMessageEvent         func(logMessage string, logLevel teamlog.LogLevel, logChannel string, logID uint64, logTime time.Time, completeLogString string)
	OnUpdateChannelEditedEvent        func(serverConnectionHandlerID uint64, channelID uint64, invokerID teamspeak.AnyID, invokerName string, invokerUniqueIdentifier string)
	OnUpdateClientEvent               func(serverConnectionHandlerID uint64, channelID teamspeak.AnyID, invokerID teamspeak.AnyID, invokerName string, invokerUniqueIdentifier string)
	OnNewChannelCreatedEvent          func(serverConnectionHandlerID uint64, channelID uint64, channelParentID uint64, invokerID teamspeak.AnyID, invokerName string, invokerUniqueIdentifier string)
	OnDelChannelEvent                 func(serverConnectionHandlerID uint64, channelID uint64, invokerID teamspeak.AnyID, invokerName string, invokerUniqueIdentifier string)
	OnChannelMoveEvent                func(serverConnectionHandlerID uint64, channelID uint64, newChannelParentID uint64, invokerID teamspeak.AnyID, invokerName string, invokerUniqueIdentifier string)
	OnClientMoveEvent                 func(serverConnectionHandlerID uint64, clientID teamspeak.AnyID, oldChannelID uint64, newChannelID uint64, visibility teamspeak.Visibility, moveMessage string)
	OnClientMoveSubscriptionEvent     func(serverConnectionHandlerID uint64, clientID teamspeak.AnyID, oldChannelID uint64, newChannelID uint64, visibility teamspeak.Visibility)
	OnServerEditedEvent               func(serverConnectionHandlerID uint64, editerID teamspeak.AnyID, editerName string, editerUniqueIdentifier string)
	OnServerErrorEvent                func(serverConnectionHandlerID uint64, errorMessage string, errorCode uint, returnCode string, extraMessage string) int
	OnServerStopEvent                 func(serverConnectionHandlerID uint64, shutdownMessage string)
	OnTextMessageEvent                func(serverConnectionHandlerID uint64, targetMode teamspeak.AnyID, toID teamspeak.AnyID, fromID teamspeak.AnyID, fromName string, fromUniqueIdentifier string, message string, ffIgnored bool) int
	OnTalkStatusChangeEvent           func(serverConnectionHandlerID uint64, status int, isReceivedWhisper int, clientID teamspeak.AnyID)
	OnConnectionInfoEvent             func(serverConnectionHandlerID uint64, clientID teamspeak.AnyID)
	OnServerConnectionInfoEvent       func(serverConnectionHandlerID uint64)
	OnChannelSubscribeEvent           func(serverConnectionHandlerID uint64, channelID uint64)
	OnChannelSubscribeFinishedEvent   func(serverConnectionHandlerID uint64)
	OnChannelUnsubscribeEvent         func(serverConnectionHandlerID uint64, channelID uint64)
	OnChannelUnsubscribeFinishedEvent func(serverConnectionHandlerID uint64)
	OnChannelDescriptionUpdateEvent   func(serverConnectionHandlerID uint64, channelID uint64)
	OnChannelPasswordChangedEvent     func(serverConnectionHandlerID uint64, channelID uint64)
	OnPlaybackShutdownCompleteEvent   func(serverConnectionHandlerID uint64)
	OnEditCapturedVoiceDataEvent      func(serverConnectionHandlerID uint64, samples *Samples, isMuted bool) (shouldMute bool)
	// TODO - use a speaker enum type?
	OnEditMixedPlaybackVoiceDataEvent func(serverConnectionHandlerID uint64, samples *Samples, channelSpeakers []uint, channelFillMask *uint)
	OnEditPlaybackVoiceDataEvent      func(serverConnectionHandlerID uint64, samples *Samples)
	OnEditPostProcessVoiceDataEvent   func(serverConnectionHandlerID uint64, clientID teamspeak.AnyID, samples *Samples, channelSpeakers []uint, channelFillMask *uint)
)

var (
	functions *TS3Functions
)

func Functions() *TS3Functions {
	return functions
}

func GetPluginID() string {
	return pluginID
}
