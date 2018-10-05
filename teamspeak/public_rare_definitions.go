package teamspeak

/*
#cgo CFLAGS: -I../pluginsdk/include -I.

#include "teamspeak/public_rare_definitions.h"

// expose defines

const unsigned int define_TS3_MAX_SIZE_CLIENT_NICKNAME_NONSDK =  TS3_MAX_SIZE_CLIENT_NICKNAME_NONSDK;
const unsigned int define_TS3_MIN_SIZE_CLIENT_NICKNAME_NONSDK =  TS3_MIN_SIZE_CLIENT_NICKNAME_NONSDK;
const unsigned int define_TS3_MAX_SIZE_AWAY_MESSAGE =  TS3_MAX_SIZE_AWAY_MESSAGE;
const unsigned int define_TS3_MAX_SIZE_GROUP_NAME =  TS3_MAX_SIZE_GROUP_NAME;
const unsigned int define_TS3_MAX_SIZE_TALK_REQUEST_MESSAGE =  TS3_MAX_SIZE_TALK_REQUEST_MESSAGE;
const unsigned int define_TS3_MAX_SIZE_COMPLAIN_MESSAGE =  TS3_MAX_SIZE_COMPLAIN_MESSAGE;
const unsigned int define_TS3_MAX_SIZE_CLIENT_DESCRIPTION =  TS3_MAX_SIZE_CLIENT_DESCRIPTION;
const unsigned int define_TS3_MAX_SIZE_HOST_MESSAGE =  TS3_MAX_SIZE_HOST_MESSAGE;
const unsigned int define_TS3_MAX_SIZE_HOSTBUTTON_TOOLTIP =  TS3_MAX_SIZE_HOSTBUTTON_TOOLTIP;
const unsigned int define_TS3_MAX_SIZE_POKE_MESSAGE =  TS3_MAX_SIZE_POKE_MESSAGE;
const unsigned int define_TS3_MAX_SIZE_OFFLINE_MESSAGE =  TS3_MAX_SIZE_OFFLINE_MESSAGE;
const unsigned int define_TS3_MAX_SIZE_OFFLINE_MESSAGE_SUBJECT =  TS3_MAX_SIZE_OFFLINE_MESSAGE_SUBJECT;
const unsigned int define_TS3_MAX_SIZE_PLUGIN_COMMAND =  TS3_MAX_SIZE_PLUGIN_COMMAND;
const unsigned int define_TS3_MAX_SIZE_VIRTUALSERVER_HOSTBANNER_GFX_URL =  TS3_MAX_SIZE_VIRTUALSERVER_HOSTBANNER_GFX_URL;

*/
import "C"

var (
	//limited length, measured in characters
	TS3_MAX_SIZE_CLIENT_NICKNAME_NONSDK = uint(C.define_TS3_MAX_SIZE_CLIENT_NICKNAME_NONSDK)

	//limited length, measured in characters
	TS3_MIN_SIZE_CLIENT_NICKNAME_NONSDK = uint(C.define_TS3_MIN_SIZE_CLIENT_NICKNAME_NONSDK)

	//limited length, measured in characters
	TS3_MAX_SIZE_AWAY_MESSAGE = uint(C.define_TS3_MAX_SIZE_AWAY_MESSAGE)

	//limited length, measured in characters
	TS3_MAX_SIZE_GROUP_NAME = uint(C.define_TS3_MAX_SIZE_GROUP_NAME)

	//limited length, measured in characters
	TS3_MAX_SIZE_TALK_REQUEST_MESSAGE = uint(C.define_TS3_MAX_SIZE_TALK_REQUEST_MESSAGE)

	//limited length, measured in characters
	TS3_MAX_SIZE_COMPLAIN_MESSAGE = uint(C.define_TS3_MAX_SIZE_COMPLAIN_MESSAGE)

	//limited length, measured in characters
	TS3_MAX_SIZE_CLIENT_DESCRIPTION = uint(C.define_TS3_MAX_SIZE_CLIENT_DESCRIPTION)

	//limited length, measured in characters
	TS3_MAX_SIZE_HOSTBUTTON_TOOLTIP = uint(C.define_TS3_MAX_SIZE_HOSTBUTTON_TOOLTIP)

	//limited length, measured in characters
	TS3_MAX_SIZE_HOST_MESSAGE = uint(C.define_TS3_MAX_SIZE_HOST_MESSAGE)

	//limited length, measured in characters
	TS3_MAX_SIZE_POKE_MESSAGE = uint(C.define_TS3_MAX_SIZE_POKE_MESSAGE)

	//limited length, measured in characters
	TS3_MAX_SIZE_OFFLINE_MESSAGE = uint(C.define_TS3_MAX_SIZE_OFFLINE_MESSAGE)

	//limited length, measured in characters
	TS3_MAX_SIZE_OFFLINE_MESSAGE_SUBJECT = uint(C.define_TS3_MAX_SIZE_OFFLINE_MESSAGE_SUBJECT)

	//limited length, measured in bytes (utf8 encoded)
	TS3_MAX_SIZE_PLUGIN_COMMAND = uint(C.define_TS3_MAX_SIZE_PLUGIN_COMMAND)

	//limited length, measured in bytes (utf8 encoded)
	TS3_MAX_SIZE_VIRTUALSERVER_HOSTBANNER_GFX_URL = uint(C.define_TS3_MAX_SIZE_VIRTUALSERVER_HOSTBANNER_GFX_URL)
)

type GroupShowNameTreeMode uint8

func (value GroupShowNameTreeMode) toC() C.enum_GroupShowNameTreeMode {
	return C.enum_GroupShowNameTreeMode(value)
}

var (
	//dont group show name
	GroupShowNameTreeModeNone = GroupShowNameTreeMode(C.GroupShowNameTreeMode_NONE)
	//show group name before client name
	GroupShowNameTreeModeBefore = GroupShowNameTreeMode(C.GroupShowNameTreeMode_BEFORE)
	//show group name behind client name
	GroupShowNameTreeModeBehind = GroupShowNameTreeMode(C.GroupShowNameTreeMode_BEHIND)
)

type PluginTargetMode uint8

func (value PluginTargetMode) toC() C.enum_PluginTargetMode {
	return C.enum_PluginTargetMode(value)
}

var (
	//send plugincmd to all clients in current channel
	PluginCommandTargetCurrentChannel = PluginTargetMode(C.PluginCommandTarget_CURRENT_CHANNEL)
	//send plugincmd to all clients on server
	PluginCommandTargetServer = PluginTargetMode(C.PluginCommandTarget_SERVER)
	//send plugincmd to all given client ids
	PluginCommandTargetClient = PluginTargetMode(C.PluginCommandTarget_CLIENT)
	//send plugincmd to all subscribed clients in current channel
	PluginCommandTargetCurrentChannelSubscribedClients = PluginTargetMode(C.PluginCommandTarget_CURRENT_CHANNEL_SUBSCRIBED_CLIENTS)

	PluginCommandTargetMax = PluginTargetMode(C.PluginCommandTarget_MAX)
)

type ServerBinding uint8

func (value ServerBinding) toC() C.int {
	return C.int(value)
}

var (
	ServerBindingVirtualServer = ServerBinding(C.SERVER_BINDING_VIRTUALSERVER)
	ServerBindingServerQuery   = ServerBinding(C.SERVER_BINDING_SERVERQUERY)
	ServerBindingFileTransfer  = ServerBinding(C.SERVER_BINDING_FILETRANSFER)
)

type HostMessageMode uint8

func (value HostMessageMode) toC() C.enum_HostMessageMode {
	return C.enum_HostMessageMode(value)
}

var (
	//dont display anything
	HostMessageModeNone = HostMessageMode(C.HostMessageMode_NONE)
	//display message inside log
	HostMessageModeLog = HostMessageMode(C.HostMessageMode_LOG)
	//display message inside a modal dialog
	HostMessageModeModal = HostMessageMode(C.HostMessageMode_MODAL)
	//display message inside a modal dialog and quit/close server/connection
	HostMessageModeModalquit = HostMessageMode(C.HostMessageMode_MODALQUIT)
)

type HostBannerMode uint8

func (value HostBannerMode) toC() C.enum_HostBannerMode {
	return C.enum_HostBannerMode(value)
}

var (
	//Do not adjust
	HostBannerModeNoAdjust = HostBannerMode(C.HostBannerMode_NO_ADJUST)
	//Adjust but ignore aspect ratio
	HostBannerModeAdjustIgnoreAspect = HostBannerMode(C.HostBannerMode_ADJUST_IGNORE_ASPECT)
	//Adjust and keep aspect ratio
	HostBannerModeAdjustKeepAspect = HostBannerMode(C.HostBannerMode_ADJUST_KEEP_ASPECT)
)

type ClientType uint8

func (value ClientType) toC() C.enum_ClientType {
	return C.enum_ClientType(value)
}

var (
	ClientTypeNormal      = ClientType(C.ClientType_NORMAL)
	ClientTypeServerQuery = ClientType(C.ClientType_SERVERQUERY)
)

type AwayStatus uint8

func (value AwayStatus) toC() C.enum_AwayStatus {
	return C.enum_AwayStatus(value)
}

var (
	AwayNone = AwayStatus(C.AWAY_NONE)
	AwayZzz  = AwayStatus(C.AWAY_ZZZ)
)

type CommandLinePropertiesRare uint8

func (value CommandLinePropertiesRare) toC() C.enum_CommandLinePropertiesRare {
	return C.enum_CommandLinePropertiesRare(value)
}

var (
	CommandLinePropertyNothing       = CommandLinePropertiesRare(C.COMMANDLINE_NOTHING)
	CommandLinePropertyEndMarkerRare = CommandLinePropertiesRare(C.COMMANDLINE_ENDMARKER_RARE)
)

type ServerInstancePropertiesRare uint8

func (value ServerInstancePropertiesRare) toC() C.enum_ServerInstancePropertiesRare {
	return C.enum_ServerInstancePropertiesRare(value)
}

var (
	ServerInstancePropertyDatabaseVersion           = ServerInstancePropertiesRare(C.SERVERINSTANCE_DATABASE_VERSION)
	ServerInstancePropertyFileTransferPort          = ServerInstancePropertiesRare(C.SERVERINSTANCE_FILETRANSFER_PORT)
	ServerInstancePropertyServerEntropy             = ServerInstancePropertiesRare(C.SERVERINSTANCE_SERVER_ENTROPY)
	ServerInstancePropertyMonthlyTimestamp          = ServerInstancePropertiesRare(C.SERVERINSTANCE_MONTHLY_TIMESTAMP)
	ServerInstancePropertyMaxDownloadTotalBandwidth = ServerInstancePropertiesRare(C.SERVERINSTANCE_MAX_DOWNLOAD_TOTAL_BANDWIDTH)
	ServerInstancePropertyMaxUploadTotalBandwidth   = ServerInstancePropertiesRare(C.SERVERINSTANCE_MAX_UPLOAD_TOTAL_BANDWIDTH)
	ServerInstancePropertyGuestServerqueryGroup     = ServerInstancePropertiesRare(C.SERVERINSTANCE_GUEST_SERVERQUERY_GROUP)
	//how many commands we can issue while in the SERVERINSTANCE_SERVERQUERY_FLOOD_TIME window
	ServerInstancePropertyServerQueryFloodCommands = ServerInstancePropertiesRare(C.SERVERINSTANCE_SERVERQUERY_FLOOD_COMMANDS)
	//time window in seconds for max command execution check
	ServerInstancePropertyServerQueryFloodTime = ServerInstancePropertiesRare(C.SERVERINSTANCE_SERVERQUERY_FLOOD_TIME)
	//how many seconds someone get banned if he floods
	ServerInstancePropertyServerQueryBanTime          = ServerInstancePropertiesRare(C.SERVERINSTANCE_SERVERQUERY_BAN_TIME)
	ServerInstancePropertyTemplateServerAdminGroup    = ServerInstancePropertiesRare(C.SERVERINSTANCE_TEMPLATE_SERVERADMIN_GROUP)
	ServerInstancePropertyTemplateServerDefaultGroup  = ServerInstancePropertiesRare(C.SERVERINSTANCE_TEMPLATE_SERVERDEFAULT_GROUP)
	ServerInstancePropertyTemplateChannelAdminGroup   = ServerInstancePropertiesRare(C.SERVERINSTANCE_TEMPLATE_CHANNELADMIN_GROUP)
	ServerInstancePropertyTemplateChannelDefaultGroup = ServerInstancePropertiesRare(C.SERVERINSTANCE_TEMPLATE_CHANNELDEFAULT_GROUP)
	ServerInstancePropertyPermissionsVersion          = ServerInstancePropertiesRare(C.SERVERINSTANCE_PERMISSIONS_VERSION)
	ServerInstancePropertyPendingConnectionsPerIp     = ServerInstancePropertiesRare(C.SERVERINSTANCE_PENDING_CONNECTIONS_PER_IP)
	ServerInstancePropertyEndMarkerRare               = ServerInstancePropertiesRare(C.SERVERINSTANCE_ENDMARKER_RARE)
)

type VirtualServerPropertiesRare uint8

func (value VirtualServerPropertiesRare) toC() C.enum_VirtualServerPropertiesRare {
	return C.enum_VirtualServerPropertiesRare(value)
}

var (
	VirtualServerPropertyDummy_0 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_0)

	VirtualServerPropertyDummy_1 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_1)

	VirtualServerPropertyDummy_2 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_2)

	VirtualServerPropertyDummy_3 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_3)

	VirtualServerPropertyDummy_4 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_4)

	VirtualServerPropertyDummy_5 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_5)

	VirtualServerPropertyDummy_6 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_6)

	VirtualServerPropertyDummy_7 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_7)

	VirtualServerPropertyDummy_8 = VirtualServerPropertiesRare(C.VIRTUALSERVER_DUMMY_8)
	//internal use
	VirtualServerPropertyKeypair = VirtualServerPropertiesRare(C.VIRTUALSERVER_KEYPAIR)
	//available when connected, not updated while connected
	VirtualServerPropertyHostmessage = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTMESSAGE)
	//available when connected, not updated while connected
	VirtualServerPropertyHostmessageMode = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTMESSAGE_MODE)
	//not available to clients, stores the folder used for file transfers
	VirtualServerPropertyFilebase = VirtualServerPropertiesRare(C.VIRTUALSERVER_FILEBASE)
	//the client permissions server group that a new client gets assigned
	VirtualServerPropertyDefaultServerGroup = VirtualServerPropertiesRare(C.VIRTUALSERVER_DEFAULT_SERVER_GROUP)
	//the channel permissions group that a new client gets assigned when joining a channel
	VirtualServerPropertyDefaultChannelGroup = VirtualServerPropertiesRare(C.VIRTUALSERVER_DEFAULT_CHANNEL_GROUP)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyFlagPassword = VirtualServerPropertiesRare(C.VIRTUALSERVER_FLAG_PASSWORD)
	//the channel permissions group that a client gets assigned when creating a channel
	VirtualServerPropertyDefaultChannelAdminGroup = VirtualServerPropertiesRare(C.VIRTUALSERVER_DEFAULT_CHANNEL_ADMIN_GROUP)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMaxDownloadTotalBandwidth = VirtualServerPropertiesRare(C.VIRTUALSERVER_MAX_DOWNLOAD_TOTAL_BANDWIDTH)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMaxUploadTotalBandwidth = VirtualServerPropertiesRare(C.VIRTUALSERVER_MAX_UPLOAD_TOTAL_BANDWIDTH)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbannerUrl = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBANNER_URL)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbannerGfxUrl = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBANNER_GFX_URL)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbannerGfxInterval = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBANNER_GFX_INTERVAL)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyComplainAutobanCount = VirtualServerPropertiesRare(C.VIRTUALSERVER_COMPLAIN_AUTOBAN_COUNT)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyComplainAutobanTime = VirtualServerPropertiesRare(C.VIRTUALSERVER_COMPLAIN_AUTOBAN_TIME)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyComplainRemoveTime = VirtualServerPropertiesRare(C.VIRTUALSERVER_COMPLAIN_REMOVE_TIME)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMinClientsInChannelBeforeForcedSilence = VirtualServerPropertiesRare(C.VIRTUALSERVER_MIN_CLIENTS_IN_CHANNEL_BEFORE_FORCED_SILENCE)
	//available when connected, always up-to-date
	VirtualServerPropertyPrioritySpeakerDimmModificator = VirtualServerPropertiesRare(C.VIRTUALSERVER_PRIORITY_SPEAKER_DIMM_MODIFICATOR)
	//available when connected
	VirtualServerPropertyId = VirtualServerPropertiesRare(C.VIRTUALSERVER_ID)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyAntifloodPointsTickReduce = VirtualServerPropertiesRare(C.VIRTUALSERVER_ANTIFLOOD_POINTS_TICK_REDUCE)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyAntifloodPointsNeededCommandBlock = VirtualServerPropertiesRare(C.VIRTUALSERVER_ANTIFLOOD_POINTS_NEEDED_COMMAND_BLOCK)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyAntifloodPointsNeededIpBlock = VirtualServerPropertiesRare(C.VIRTUALSERVER_ANTIFLOOD_POINTS_NEEDED_IP_BLOCK)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyClientConnections = VirtualServerPropertiesRare(C.VIRTUALSERVER_CLIENT_CONNECTIONS)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyQueryClientConnections = VirtualServerPropertiesRare(C.VIRTUALSERVER_QUERY_CLIENT_CONNECTIONS)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbuttonTooltip = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBUTTON_TOOLTIP)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbuttonUrl = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBUTTON_URL)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbuttonGfxUrl = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBUTTON_GFX_URL)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyQueryClientsOnline = VirtualServerPropertiesRare(C.VIRTUALSERVER_QUERYCLIENTS_ONLINE)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyDownloadQuota = VirtualServerPropertiesRare(C.VIRTUALSERVER_DOWNLOAD_QUOTA)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyUploadQuota = VirtualServerPropertiesRare(C.VIRTUALSERVER_UPLOAD_QUOTA)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMonthBytesDownloaded = VirtualServerPropertiesRare(C.VIRTUALSERVER_MONTH_BYTES_DOWNLOADED)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMonthBytesUploaded = VirtualServerPropertiesRare(C.VIRTUALSERVER_MONTH_BYTES_UPLOADED)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalBytesDownloaded = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_BYTES_DOWNLOADED)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalBytesUploaded = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_BYTES_UPLOADED)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyPort = VirtualServerPropertiesRare(C.VIRTUALSERVER_PORT)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyAutostart = VirtualServerPropertiesRare(C.VIRTUALSERVER_AUTOSTART)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMachineId = VirtualServerPropertiesRare(C.VIRTUALSERVER_MACHINE_ID)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyNeededIdentitySecurityLevel = VirtualServerPropertiesRare(C.VIRTUALSERVER_NEEDED_IDENTITY_SECURITY_LEVEL)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyLogClient = VirtualServerPropertiesRare(C.VIRTUALSERVER_LOG_CLIENT)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyLogQuery = VirtualServerPropertiesRare(C.VIRTUALSERVER_LOG_QUERY)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyLogChannel = VirtualServerPropertiesRare(C.VIRTUALSERVER_LOG_CHANNEL)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyLogPermissions = VirtualServerPropertiesRare(C.VIRTUALSERVER_LOG_PERMISSIONS)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyLogServer = VirtualServerPropertiesRare(C.VIRTUALSERVER_LOG_SERVER)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyLogFileTransfer = VirtualServerPropertiesRare(C.VIRTUALSERVER_LOG_FILETRANSFER)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMinClientVersion = VirtualServerPropertiesRare(C.VIRTUALSERVER_MIN_CLIENT_VERSION)
	//available when connected, always up-to-date
	VirtualServerPropertyNamePhonetic = VirtualServerPropertiesRare(C.VIRTUALSERVER_NAME_PHONETIC)
	//available when connected, always up-to-date
	VirtualServerPropertyIconId = VirtualServerPropertiesRare(C.VIRTUALSERVER_ICON_ID)
	//available when connected, always up-to-date
	VirtualServerPropertyReservedSlots = VirtualServerPropertiesRare(C.VIRTUALSERVER_RESERVED_SLOTS)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalPacketLossSpeech = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_PACKETLOSS_SPEECH)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalPacketLossKeepalive = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_PACKETLOSS_KEEPALIVE)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalPacketLossControl = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_PACKETLOSS_CONTROL)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalPacketLossTotal = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_PACKETLOSS_TOTAL)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyTotalPing = VirtualServerPropertiesRare(C.VIRTUALSERVER_TOTAL_PING)
	//internal use | contains comma separated ip list
	VirtualServerPropertyIP = VirtualServerPropertiesRare(C.VIRTUALSERVER_IP)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyWeblistEnabled = VirtualServerPropertiesRare(C.VIRTUALSERVER_WEBLIST_ENABLED)
	//internal use
	VirtualServerPropertyAutogeneratedPrivilegeKey = VirtualServerPropertiesRare(C.VIRTUALSERVER_AUTOGENERATED_PRIVILEGEKEY)
	//available when connected
	VirtualServerPropertyAskForPrivilegeKey = VirtualServerPropertiesRare(C.VIRTUALSERVER_ASK_FOR_PRIVILEGEKEY)
	//available when connected, always up-to-date
	VirtualServerPropertyHostbannerMode = VirtualServerPropertiesRare(C.VIRTUALSERVER_HOSTBANNER_MODE)
	//available when connected, always up-to-date
	VirtualServerPropertyChannelTempDeleteDelayDefault = VirtualServerPropertiesRare(C.VIRTUALSERVER_CHANNEL_TEMP_DELETE_DELAY_DEFAULT)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMinAndroidVersion = VirtualServerPropertiesRare(C.VIRTUALSERVER_MIN_ANDROID_VERSION)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMinIosVersion = VirtualServerPropertiesRare(C.VIRTUALSERVER_MIN_IOS_VERSION)
	//only available on request (=> requestServerVariables)
	VirtualServerPropertyMinWinPhoneVersion = VirtualServerPropertiesRare(C.VIRTUALSERVER_MIN_WINPHONE_VERSION)

	VirtualServerPropertyEndMarkerRare = VirtualServerPropertiesRare(C.VIRTUALSERVER_ENDMARKER_RARE)
)

type ChannelPropertiesRare uint8

func (value ChannelPropertiesRare) toC() C.enum_ChannelPropertiesRare {
	return C.enum_ChannelPropertiesRare(value)
}

var (
	ChannelPropertyDummy_2 = ChannelPropertiesRare(C.CHANNEL_DUMMY_2)

	ChannelPropertyDummy_3 = ChannelPropertiesRare(C.CHANNEL_DUMMY_3)

	ChannelPropertyDummy_4 = ChannelPropertiesRare(C.CHANNEL_DUMMY_4)

	ChannelPropertyDummy_5 = ChannelPropertiesRare(C.CHANNEL_DUMMY_5)

	ChannelPropertyDummy_6 = ChannelPropertiesRare(C.CHANNEL_DUMMY_6)

	ChannelPropertyDummy_7 = ChannelPropertiesRare(C.CHANNEL_DUMMY_7)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagMaxClientsUnlimited = ChannelPropertiesRare(C.CHANNEL_FLAG_MAXCLIENTS_UNLIMITED)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagMaxFamilyClientsUnlimited = ChannelPropertiesRare(C.CHANNEL_FLAG_MAXFAMILYCLIENTS_UNLIMITED)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagMaxFamilyClientsInherited = ChannelPropertiesRare(C.CHANNEL_FLAG_MAXFAMILYCLIENTS_INHERITED)
	//Only available client side, stores whether we are subscribed to this channel
	ChannelPropertyFlagAreSubscribed = ChannelPropertiesRare(C.CHANNEL_FLAG_ARE_SUBSCRIBED)
	//not available client side, the folder used for file-transfers for this channel
	ChannelPropertyFilePath = ChannelPropertiesRare(C.CHANNEL_FILEPATH)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyNeededTalkPower = ChannelPropertiesRare(C.CHANNEL_NEEDED_TALK_POWER)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyForcedSilence = ChannelPropertiesRare(C.CHANNEL_FORCED_SILENCE)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyNamePhonetic = ChannelPropertiesRare(C.CHANNEL_NAME_PHONETIC)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyIconId = ChannelPropertiesRare(C.CHANNEL_ICON_ID)
	//Available for all channels that are "in view", always up-to-date
	ChannelPropertyFlagPrivate = ChannelPropertiesRare(C.CHANNEL_FLAG_PRIVATE)

	ChannelPropertyEndMarkerRare = ChannelPropertiesRare(C.CHANNEL_ENDMARKER_RARE)
)

type ClientPropertiesRare uint8

func (value ClientPropertiesRare) toC() C.enum_ClientPropertiesRare {
	return C.enum_ClientPropertiesRare(value)
}

var (
	ClientPropertyDummy_3 = ClientPropertiesRare(C.CLIENT_DUMMY_3)

	ClientPropertyDummy_4 = ClientPropertiesRare(C.CLIENT_DUMMY_4)

	ClientPropertyDummy_5 = ClientPropertiesRare(C.CLIENT_DUMMY_5)

	ClientPropertyDummy_6 = ClientPropertiesRare(C.CLIENT_DUMMY_6)

	ClientPropertyDummy_7 = ClientPropertiesRare(C.CLIENT_DUMMY_7)

	ClientPropertyDummy_8 = ClientPropertiesRare(C.CLIENT_DUMMY_8)

	ClientPropertyDummy_9 = ClientPropertiesRare(C.CLIENT_DUMMY_9)
	//internal use
	ClientPropertyKeyOffset = ClientPropertiesRare(C.CLIENT_KEY_OFFSET)
	//internal use
	ClientPropertyLastVarRequest = ClientPropertiesRare(C.CLIENT_LAST_VAR_REQUEST)
	//used for serverquery clients, makes no sense on normal clients currently
	ClientPropertyLoginName = ClientPropertiesRare(C.CLIENT_LOGIN_NAME)
	//used for serverquery clients, makes no sense on normal clients currently
	ClientPropertyLoginPassword = ClientPropertiesRare(C.CLIENT_LOGIN_PASSWORD)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, holds database client id
	ClientPropertyDatabaseID = ClientPropertiesRare(C.CLIENT_DATABASE_ID)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, holds database client id
	ClientPropertyChannelGroupId = ClientPropertiesRare(C.CLIENT_CHANNEL_GROUP_ID)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, holds all servergroups client belongs too
	ClientPropertyServerGroups = ClientPropertiesRare(C.CLIENT_SERVERGROUPS)
	//this needs to be requested (=> requestClientVariables), first time this client connected to this server
	ClientPropertyCreated = ClientPropertiesRare(C.CLIENT_CREATED)
	//this needs to be requested (=> requestClientVariables), last time this client connected to this server
	ClientPropertyLastConnected = ClientPropertiesRare(C.CLIENT_LASTCONNECTED)
	//this needs to be requested (=> requestClientVariables), how many times this client connected to this server
	ClientPropertyTotalConnections = ClientPropertiesRare(C.CLIENT_TOTALCONNECTIONS)
	//automatically up-to-date for any client "in view", this clients away status
	ClientPropertyAway = ClientPropertiesRare(C.CLIENT_AWAY)
	//automatically up-to-date for any client "in view", this clients away message
	ClientPropertyAwayMessage = ClientPropertiesRare(C.CLIENT_AWAY_MESSAGE)
	//automatically up-to-date for any client "in view", determines if this is a real client or a server-query connection
	ClientPropertyType = ClientPropertiesRare(C.CLIENT_TYPE)
	//automatically up-to-date for any client "in view", this client got an avatar
	ClientPropertyFlagAvatar = ClientPropertiesRare(C.CLIENT_FLAG_AVATAR)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, holds database client id
	ClientPropertyTalkPower = ClientPropertiesRare(C.CLIENT_TALK_POWER)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, holds timestamp where client requested to talk
	ClientPropertyTalkRequest = ClientPropertiesRare(C.CLIENT_TALK_REQUEST)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, holds matter for the request
	ClientPropertyTalkRequestMsg = ClientPropertiesRare(C.CLIENT_TALK_REQUEST_MSG)
	//automatically up-to-date for any client "in view"
	ClientPropertyDescription = ClientPropertiesRare(C.CLIENT_DESCRIPTION)
	//automatically up-to-date for any client "in view"
	ClientPropertyIsTalker = ClientPropertiesRare(C.CLIENT_IS_TALKER)
	//this needs to be requested (=> requestClientVariables)
	ClientPropertyMonthBytesUploaded = ClientPropertiesRare(C.CLIENT_MONTH_BYTES_UPLOADED)
	//this needs to be requested (=> requestClientVariables)
	ClientPropertyMonthBytesDownloaded = ClientPropertiesRare(C.CLIENT_MONTH_BYTES_DOWNLOADED)
	//this needs to be requested (=> requestClientVariables)
	ClientPropertyTotalBytesUploaded = ClientPropertiesRare(C.CLIENT_TOTAL_BYTES_UPLOADED)
	//this needs to be requested (=> requestClientVariables)
	ClientPropertyTotalBytesDownloaded = ClientPropertiesRare(C.CLIENT_TOTAL_BYTES_DOWNLOADED)
	//automatically up-to-date for any client "in view"
	ClientPropertyIsPrioritySpeaker = ClientPropertiesRare(C.CLIENT_IS_PRIORITY_SPEAKER)
	//automatically up-to-date for any client "in view"
	ClientPropertyUnreadMessages = ClientPropertiesRare(C.CLIENT_UNREAD_MESSAGES)
	//automatically up-to-date for any client "in view"
	ClientPropertyNicknamePhonetic = ClientPropertiesRare(C.CLIENT_NICKNAME_PHONETIC)
	//automatically up-to-date for any client "in view"
	ClientPropertyNeededServerQueryViewPower = ClientPropertiesRare(C.CLIENT_NEEDED_SERVERQUERY_VIEW_POWER)
	//only usable for ourself, the default token we used to connect on our last connection attempt
	ClientPropertyDefaultToken = ClientPropertiesRare(C.CLIENT_DEFAULT_TOKEN)
	//automatically up-to-date for any client "in view"
	ClientPropertyIconId = ClientPropertiesRare(C.CLIENT_ICON_ID)
	//automatically up-to-date for any client "in view"
	ClientPropertyIsChannelCommander = ClientPropertiesRare(C.CLIENT_IS_CHANNEL_COMMANDER)
	//automatically up-to-date for any client "in view"
	ClientPropertyCountry = ClientPropertiesRare(C.CLIENT_COUNTRY)
	//automatically up-to-date for any client "in view", only valid with PERMISSION feature, contains channel_id where the channel_group_id is set from
	ClientPropertyChannelGroupInheritedChannelId = ClientPropertiesRare(C.CLIENT_CHANNEL_GROUP_INHERITED_CHANNEL_ID)
	//automatically up-to-date for any client "in view", stores icons for partner badges
	ClientPropertyBadges = ClientPropertiesRare(C.CLIENT_BADGES)

	ClientPropertyEndMarkerRare = ClientPropertiesRare(C.CLIENT_ENDMARKER_RARE)
)

type ConnectionPropertiesRare uint8

func (value ConnectionPropertiesRare) toC() C.enum_ConnectionPropertiesRare {
	return C.enum_ConnectionPropertiesRare(value)
}

var (
	ConnectionPropertyDummy_0 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_0)

	ConnectionPropertyDummy_1 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_1)

	ConnectionPropertyDummy_2 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_2)

	ConnectionPropertyDummy_3 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_3)

	ConnectionPropertyDummy_4 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_4)

	ConnectionPropertyDummy_5 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_5)

	ConnectionPropertyDummy_6 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_6)

	ConnectionPropertyDummy_7 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_7)

	ConnectionPropertyDummy_8 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_8)

	ConnectionPropertyDummy_9 = ConnectionPropertiesRare(C.CONNECTION_DUMMY_9)
	//how many bytes per second are currently being sent by file transfers
	ConnectionPropertyFileTransferBandwidthSent = ConnectionPropertiesRare(C.CONNECTION_FILETRANSFER_BANDWIDTH_SENT)
	//how many bytes per second are currently being received by file transfers
	ConnectionPropertyFileTransferBandwidthReceived = ConnectionPropertiesRare(C.CONNECTION_FILETRANSFER_BANDWIDTH_RECEIVED)
	//how many bytes we received in total through file transfers
	ConnectionPropertyFileTransferBytesReceivedTotal = ConnectionPropertiesRare(C.CONNECTION_FILETRANSFER_BYTES_RECEIVED_TOTAL)
	//how many bytes we sent in total through file transfers
	ConnectionPropertyFileTransferBytesSentTotal = ConnectionPropertiesRare(C.CONNECTION_FILETRANSFER_BYTES_SENT_TOTAL)

	ConnectionPropertyEndMarkerRare = ConnectionPropertiesRare(C.CONNECTION_ENDMARKER_RARE)
)

type LicenseViolationType uint8

func (value LicenseViolationType) toC() C.enum_LicenseViolationType {
	return C.enum_LicenseViolationType(value)
}

var (
	NoViolation   = LicenseViolationType(C.NO_VIOLATION)
	SlotViolation = LicenseViolationType(C.SLOT_VIOLATION)
	SlotSuspicion = LicenseViolationType(C.SLOT_SUSPICION)
)

type BBCodeTag uint32

func (value BBCodeTag) toC() C.enum_BBCodeTags {
	return C.enum_BBCodeTags(value)
}

var (
	BBCodeTagB            = BBCodeTag(C.BBCodeTag_B)
	BBCodeTagI            = BBCodeTag(C.BBCodeTag_I)
	BBCodeTagU            = BBCodeTag(C.BBCodeTag_U)
	BBCodeTagS            = BBCodeTag(C.BBCodeTag_S)
	BBCodeTagSup          = BBCodeTag(C.BBCodeTag_SUP)
	BBCodeTagSub          = BBCodeTag(C.BBCodeTag_SUB)
	BBCodeTagColor        = BBCodeTag(C.BBCodeTag_COLOR)
	BBCodeTagSize         = BBCodeTag(C.BBCodeTag_SIZE)
	BBCodeTagGroupText    = BBCodeTag(C.BBCodeTag_group_text)
	BBCodeTagLeft         = BBCodeTag(C.BBCodeTag_LEFT)
	BBCodeTagRight        = BBCodeTag(C.BBCodeTag_RIGHT)
	BBCodeTagCenter       = BBCodeTag(C.BBCodeTag_CENTER)
	BBCodeTagGroupAlign   = BBCodeTag(C.BBCodeTag_group_align)
	BBCodeTagUrl          = BBCodeTag(C.BBCodeTag_URL)
	BBCodeTagImage        = BBCodeTag(C.BBCodeTag_IMAGE)
	BBCodeTagHr           = BBCodeTag(C.BBCodeTag_HR)
	BBCodeTagList         = BBCodeTag(C.BBCodeTag_LIST)
	BBCodeTagListitem     = BBCodeTag(C.BBCodeTag_LISTITEM)
	BBCodeTagGroupList    = BBCodeTag(C.BBCodeTag_group_list)
	BBCodeTagTable        = BBCodeTag(C.BBCodeTag_TABLE)
	BBCodeTagTr           = BBCodeTag(C.BBCodeTag_TR)
	BBCodeTagTh           = BBCodeTag(C.BBCodeTag_TH)
	BBCodeTagTd           = BBCodeTag(C.BBCodeTag_TD)
	BBCodeTagGroupTable   = BBCodeTag(C.BBCodeTag_group_table)
	BBCodeTagDefSimple    = BBCodeTag(C.BBCodeTag_def_simple)
	BBCodeTagDefSimpleImg = BBCodeTag(C.BBCodeTag_def_simple_Img)
	BBCodeTagDefExtended  = BBCodeTag(C.BBCodeTag_def_extended)
)

// typedef int(*ExtraBBCodeValidator)(void* userparam, const char* tag, const char* paramValue, int paramValueSize, const char* childValue, int childValueSize); // TODO
// typedef const char* (*ExtraBBCodeParamTransform)(void* userparam, const char* tag, const char* paramValue); // TODO
