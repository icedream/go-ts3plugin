package teamspeak

/*
#cgo CFLAGS: -I../pluginsdk/include -I.

#include "teamspeak/public_errors_rare.h"
*/
import "C"

/******************************************************************************
From public_errors_rare.h
*******************************************************************************/

//client
var (
	ErrorClientInvalidPassword        = uint32(C.ERROR_client_invalid_password)
	ErrorClientTooManyClonesConnected = uint32(C.ERROR_client_too_many_clones_connected)
	ErrorClientIsOnline               = uint32(C.ERROR_client_is_online)
)

//channel
var (
	ErrorChannelIsPrivateChannel = uint32(C.ERROR_channel_is_private_channel)

//note 0x030f is defined in public_errors;
)

//database
var (
	ErrorDatabase                = uint32(C.ERROR_database)
	ErrorDatabaseEmptyResult     = uint32(C.ERROR_database_empty_result)
	ErrorDatabaseDuplicateEntry  = uint32(C.ERROR_database_duplicate_entry)
	ErrorDatabaseNoModifications = uint32(C.ERROR_database_no_modifications)
	ErrorDatabaseConstraint      = uint32(C.ERROR_database_constraint)
	ErrorDatabaseReinvoke        = uint32(C.ERROR_database_reinvoke)
)

//permissions
var (
	ErrorPermissionInvalidGroupId               = uint32(C.ERROR_permission_invalid_group_id)
	ErrorPermissionDuplicateEntry               = uint32(C.ERROR_permission_duplicate_entry)
	ErrorPermissionInvalidPermId                = uint32(C.ERROR_permission_invalid_perm_id)
	ErrorPermissionEmptyResult                  = uint32(C.ERROR_permission_empty_result)
	ErrorPermissionDefaultGroupForbidden        = uint32(C.ERROR_permission_default_group_forbidden)
	ErrorPermissionInvalidSize                  = uint32(C.ERROR_permission_invalid_size)
	ErrorPermissionInvalidValue                 = uint32(C.ERROR_permission_invalid_value)
	ErrorPermissionsGroupNotEmpty               = uint32(C.ERROR_permissions_group_not_empty)
	ErrorPermissionsInsufficientGroupPower      = uint32(C.ERROR_permissions_insufficient_group_power)
	ErrorPermissionsInsufficientPermissionPower = uint32(C.ERROR_permissions_insufficient_permission_power)
	ErrorPermissionTemplateGroupIsUsed          = uint32(C.ERROR_permission_template_group_is_used)

//0x0a0c is in public_errors.h
)

//server
var (
	ErrorServerDeploymentActive      = uint32(C.ERROR_server_deployment_active)
	ErrorServerUnableToStopOwnServer = uint32(C.ERROR_server_unable_to_stop_own_server)
	ErrorServerWrongMachineid        = uint32(C.ERROR_server_wrong_machineid)
	ErrorServerModalQuit             = uint32(C.ERROR_server_modal_quit)
)

//messages
var (
	ErrorMessageInvalidId = uint32(C.ERROR_message_invalid_id)
)

//ban
var (
	ErrorBanInvalidId        = uint32(C.ERROR_ban_invalid_id)
	ErrorConnectFailedBanned = uint32(C.ERROR_connect_failed_banned)
	ErrorRenameFailedBanned  = uint32(C.ERROR_rename_failed_banned)
	ErrorBanFlooding         = uint32(C.ERROR_ban_flooding)
)

//tts
var (
	ErrorTtsUnableToInitialize = uint32(C.ERROR_tts_unable_to_initialize)
)

//privilege key
var (
	ErrorPrivilegeKeyInvalid = uint32(C.ERROR_privilege_key_invalid)
)

//voip
var (
	ErrorVoipPjsua                     = uint32(C.ERROR_voip_pjsua)
	ErrorVoipAlreadyInitialized        = uint32(C.ERROR_voip_already_initialized)
	ErrorVoipTooManyAccounts           = uint32(C.ERROR_voip_too_many_accounts)
	ErrorVoipInvalidAccount            = uint32(C.ERROR_voip_invalid_account)
	ErrorVoipInternalError             = uint32(C.ERROR_voip_internal_error)
	ErrorVoipInvalidConnectionId       = uint32(C.ERROR_voip_invalid_connectionId)
	ErrorVoipCannotAnswerInitiatedCall = uint32(C.ERROR_voip_cannot_answer_initiated_call)
	ErrorVoipNotInitialized            = uint32(C.ERROR_voip_not_initialized)
)
