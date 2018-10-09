package teamspeak

/*
#cgo CFLAGS: -I../pluginsdk/include -I.

#include "teamspeak/public_errors.h"
*/
import "C"

/******************************************************************************
From public_errors.h
*******************************************************************************/

//The idea here is: the values are 2 bytes wide, the first byte identifies the group, the second the count within that group

//general
var ErrorOK = uint32(C.ERROR_ok)
var ErrorUndefined = uint32(C.ERROR_undefined)
var ErrorNotImplemented = uint32(C.ERROR_not_implemented)
var ErrorOkNoUpdate = uint32(C.ERROR_ok_no_update)
var ErrorDontNotify = uint32(C.ERROR_dont_notify)
var ErrorLibTimeLimitReached = uint32(C.ERROR_lib_time_limit_reached)

//dunno
var ErrorCommandNotFound = uint32(C.ERROR_command_not_found)
var ErrorUnableToBindNetworkPort = uint32(C.ERROR_unable_to_bind_network_port)
var ErrorNoNetworkPortAvailable = uint32(C.ERROR_no_network_port_available)
var ErrorPortAlreadyInUse = uint32(C.ERROR_port_already_in_use)

//client
var ErrorClientInvalidId = uint32(C.ERROR_client_invalid_id)
var ErrorClientNicknameInuse = uint32(C.ERROR_client_nickname_inuse)
var ErrorClientProtocolLimitReached = uint32(C.ERROR_client_protocol_limit_reached)
var ErrorClientInvalidType = uint32(C.ERROR_client_invalid_type)
var ErrorClientAlreadySubscribed = uint32(C.ERROR_client_already_subscribed)
var ErrorClientNotLoggedIn = uint32(C.ERROR_client_not_logged_in)
var ErrorClientCouldNotValidateIdentity = uint32(C.ERROR_client_could_not_validate_identity)
var ErrorClientVersionOutdated = uint32(C.ERROR_client_version_outdated)
var ErrorClientIsFlooding = uint32(C.ERROR_client_is_flooding)
var ErrorClientHacked = uint32(C.ERROR_client_hacked)
var ErrorClientCannotVerifyNow = uint32(C.ERROR_client_cannot_verify_now)
var ErrorClientLoginNotPermitted = uint32(C.ERROR_client_login_not_permitted)
var ErrorClientNotSubscribed = uint32(C.ERROR_client_not_subscribed)

//channel
var ErrorChannelInvalidId = uint32(C.ERROR_channel_invalid_id)
var ErrorChannelProtocolLimitReached = uint32(C.ERROR_channel_protocol_limit_reached)
var ErrorChannelAlreadyIn = uint32(C.ERROR_channel_already_in)
var ErrorChannelNameInuse = uint32(C.ERROR_channel_name_inuse)
var ErrorChannelNotEmpty = uint32(C.ERROR_channel_not_empty)
var ErrorChannelCanNotDeleteDefault = uint32(C.ERROR_channel_can_not_delete_default)
var ErrorChannelDefaultRequirePermanent = uint32(C.ERROR_channel_default_require_permanent)
var ErrorChannelInvalidFlags = uint32(C.ERROR_channel_invalid_flags)
var ErrorChannelParentNotPermanent = uint32(C.ERROR_channel_parent_not_permanent)
var ErrorChannelMaxclientsReached = uint32(C.ERROR_channel_maxclients_reached)
var ErrorChannelMaxfamilyReached = uint32(C.ERROR_channel_maxfamily_reached)
var ErrorChannelInvalidOrder = uint32(C.ERROR_channel_invalid_order)
var ErrorChannelNoFiletransferSupported = uint32(C.ERROR_channel_no_filetransfer_supported)
var ErrorChannelInvalidPassword = uint32(C.ERROR_channel_invalid_password)
var ErrorChannelInvalidSecurityHash = uint32(C.ERROR_channel_invalid_security_hash)

//note 0x030e is defined in public_rare_errors

//server
var ErrorServerInvalidId = uint32(C.ERROR_server_invalid_id)
var ErrorServerRunning = uint32(C.ERROR_server_running)
var ErrorServerIsShuttingDown = uint32(C.ERROR_server_is_shutting_down)
var ErrorServerMaxclientsReached = uint32(C.ERROR_server_maxclients_reached)
var ErrorServerInvalidPassword = uint32(C.ERROR_server_invalid_password)
var ErrorServerIsVirtual = uint32(C.ERROR_server_is_virtual)
var ErrorServerIsNotRunning = uint32(C.ERROR_server_is_not_running)
var ErrorServerIsBooting = uint32(C.ERROR_server_is_booting)
var ErrorServerStatusInvalid = uint32(C.ERROR_server_status_invalid)
var ErrorServerVersionOutdated = uint32(C.ERROR_server_version_outdated)
var ErrorServerDuplicateRunning = uint32(C.ERROR_server_duplicate_running)

//parameter
var ErrorParameterQuote = uint32(C.ERROR_parameter_quote)
var ErrorParameterInvalidCount = uint32(C.ERROR_parameter_invalid_count)
var ErrorParameterInvalid = uint32(C.ERROR_parameter_invalid)
var ErrorParameterNotFound = uint32(C.ERROR_parameter_not_found)
var ErrorParameterConvert = uint32(C.ERROR_parameter_convert)
var ErrorParameterInvalidSize = uint32(C.ERROR_parameter_invalid_size)
var ErrorParameterMissing = uint32(C.ERROR_parameter_missing)
var ErrorParameterChecksum = uint32(C.ERROR_parameter_checksum)

//unsorted, need further investigation
var ErrorVsCritical = uint32(C.ERROR_vs_critical)
var ErrorConnectionLost = uint32(C.ERROR_connection_lost)
var ErrorNotConnected = uint32(C.ERROR_not_connected)
var ErrorNoCachedConnectionInfo = uint32(C.ERROR_no_cached_connection_info)
var ErrorCurrentlyNotPossible = uint32(C.ERROR_currently_not_possible)
var ErrorFailedConnectionInitialisation = uint32(C.ERROR_failed_connection_initialisation)
var ErrorCouldNotResolveHostname = uint32(C.ERROR_could_not_resolve_hostname)
var ErrorInvalidServerConnectionHandlerId = uint32(C.ERROR_invalid_server_connection_handler_id)
var ErrorCouldNotInitialiseInputManager = uint32(C.ERROR_could_not_initialise_input_manager)
var ErrorClientlibraryNotInitialised = uint32(C.ERROR_clientlibrary_not_initialised)
var ErrorServerlibraryNotInitialised = uint32(C.ERROR_serverlibrary_not_initialised)
var ErrorWhisperTooManyTargets = uint32(C.ERROR_whisper_too_many_targets)
var ErrorWhisperNoTargets = uint32(C.ERROR_whisper_no_targets)
var ErrorConnectionIpProtocolMissing = uint32(C.ERROR_connection_ip_protocol_missing)

//file transfer
var ErrorFileInvalidName = uint32(C.ERROR_file_invalid_name)
var ErrorFileInvalidPermissions = uint32(C.ERROR_file_invalid_permissions)
var ErrorFileAlreadyExists = uint32(C.ERROR_file_already_exists)
var ErrorFileNotFound = uint32(C.ERROR_file_not_found)
var ErrorFileIoError = uint32(C.ERROR_file_io_error)
var ErrorFileInvalidTransferId = uint32(C.ERROR_file_invalid_transfer_id)
var ErrorFileInvalidPath = uint32(C.ERROR_file_invalid_path)
var ErrorFileNoFilesAvailable = uint32(C.ERROR_file_no_files_available)
var ErrorFileOverwriteExcludesResume = uint32(C.ERROR_file_overwrite_excludes_resume)
var ErrorFileInvalidSize = uint32(C.ERROR_file_invalid_size)
var ErrorFileAlreadyInUse = uint32(C.ERROR_file_already_in_use)
var ErrorFileCouldNotOpenConnection = uint32(C.ERROR_file_could_not_open_connection)
var ErrorFileNoSpaceLeftOnDevice = uint32(C.ERROR_file_no_space_left_on_device)
var ErrorFileExceedsFileSystemMaximumSize = uint32(C.ERROR_file_exceeds_file_system_maximum_size)
var ErrorFileTransferConnectionTimeout = uint32(C.ERROR_file_transfer_connection_timeout)
var ErrorFileConnectionLost = uint32(C.ERROR_file_connection_lost)
var ErrorFileExceedsSuppliedSize = uint32(C.ERROR_file_exceeds_supplied_size)
var ErrorFileTransferComplete = uint32(C.ERROR_file_transfer_complete)
var ErrorFileTransferCanceled = uint32(C.ERROR_file_transfer_canceled)
var ErrorFileTransferInterrupted = uint32(C.ERROR_file_transfer_interrupted)
var ErrorFileTransferServerQuotaExceeded = uint32(C.ERROR_file_transfer_server_quota_exceeded)
var ErrorFileTransferClientQuotaExceeded = uint32(C.ERROR_file_transfer_client_quota_exceeded)
var ErrorFileTransferReset = uint32(C.ERROR_file_transfer_reset)
var ErrorFileTransferLimitReached = uint32(C.ERROR_file_transfer_limit_reached)

//sound
var ErrorSoundPreprocessorDisabled = uint32(C.ERROR_sound_preprocessor_disabled)
var ErrorSoundInternalPreprocessor = uint32(C.ERROR_sound_internal_preprocessor)
var ErrorSoundInternalEncoder = uint32(C.ERROR_sound_internal_encoder)
var ErrorSoundInternalPlayback = uint32(C.ERROR_sound_internal_playback)
var ErrorSoundNoCaptureDeviceAvailable = uint32(C.ERROR_sound_no_capture_device_available)
var ErrorSoundNoPlaybackDeviceAvailable = uint32(C.ERROR_sound_no_playback_device_available)
var ErrorSoundCouldNotOpenCaptureDevice = uint32(C.ERROR_sound_could_not_open_capture_device)
var ErrorSoundCouldNotOpenPlaybackDevice = uint32(C.ERROR_sound_could_not_open_playback_device)
var ErrorSoundHandlerHasDevice = uint32(C.ERROR_sound_handler_has_device)
var ErrorSoundInvalidCaptureDevice = uint32(C.ERROR_sound_invalid_capture_device)
var ErrorSoundInvalidPlaybackDevice = uint32(C.ERROR_sound_invalid_playback_device)
var ErrorSoundInvalidWave = uint32(C.ERROR_sound_invalid_wave)
var ErrorSoundUnsupportedWave = uint32(C.ERROR_sound_unsupported_wave)
var ErrorSoundOpenWave = uint32(C.ERROR_sound_open_wave)
var ErrorSoundInternalCapture = uint32(C.ERROR_sound_internal_capture)
var ErrorSoundDeviceInUse = uint32(C.ERROR_sound_device_in_use)
var ErrorSoundDeviceAlreadyRegisterred = uint32(C.ERROR_sound_device_already_registerred)
var ErrorSoundUnknownDevice = uint32(C.ERROR_sound_unknown_device)
var ErrorSoundUnsupportedFrequency = uint32(C.ERROR_sound_unsupported_frequency)
var ErrorSoundInvalidChannelCount = uint32(C.ERROR_sound_invalid_channel_count)
var ErrorSoundReadWave = uint32(C.ERROR_sound_read_wave)

//for internal purposes only
var ErrorSoundNeedMoreData = uint32(C.ERROR_sound_need_more_data)

//for internal purposes only
var ErrorSoundDeviceBusy = uint32(C.ERROR_sound_device_busy)
var ErrorSoundNoData = uint32(C.ERROR_sound_no_data)
var ErrorSoundChannelMaskMismatch = uint32(C.ERROR_sound_channel_mask_mismatch)

//permissions
var ErrorPermissionsClientInsufficient = uint32(C.ERROR_permissions_client_insufficient)
var ErrorPermissions = uint32(C.ERROR_permissions)

//accounting
var ErrorAccountingVirtualserverLimitReached = uint32(C.ERROR_accounting_virtualserver_limit_reached)
var ErrorAccountingSlotLimitReached = uint32(C.ERROR_accounting_slot_limit_reached)
var ErrorAccountingLicenseFileNotFound = uint32(C.ERROR_accounting_license_file_not_found)
var ErrorAccountingLicenseDateNotOk = uint32(C.ERROR_accounting_license_date_not_ok)
var ErrorAccountingUnableToConnectToServer = uint32(C.ERROR_accounting_unable_to_connect_to_server)
var ErrorAccountingUnknownError = uint32(C.ERROR_accounting_unknown_error)
var ErrorAccountingServerError = uint32(C.ERROR_accounting_server_error)
var ErrorAccountingInstanceLimitReached = uint32(C.ERROR_accounting_instance_limit_reached)
var ErrorAccountingInstanceCheckError = uint32(C.ERROR_accounting_instance_check_error)
var ErrorAccountingLicenseFileInvalid = uint32(C.ERROR_accounting_license_file_invalid)
var ErrorAccountingRunningElsewhere = uint32(C.ERROR_accounting_running_elsewhere)
var ErrorAccountingInstanceDuplicated = uint32(C.ERROR_accounting_instance_duplicated)
var ErrorAccountingAlreadyStarted = uint32(C.ERROR_accounting_already_started)
var ErrorAccountingNotStarted = uint32(C.ERROR_accounting_not_started)
var ErrorAccountingToManyStarts = uint32(C.ERROR_accounting_to_many_starts)

//provisioning server
var ErrorProvisioningInvalidPassword = uint32(C.ERROR_provisioning_invalid_password)
var ErrorProvisioningInvalidRequest = uint32(C.ERROR_provisioning_invalid_request)
var ErrorProvisioningNoSlotsAvailable = uint32(C.ERROR_provisioning_no_slots_available)
var ErrorProvisioningPoolMissing = uint32(C.ERROR_provisioning_pool_missing)
var ErrorProvisioningPoolUnknown = uint32(C.ERROR_provisioning_pool_unknown)
var ErrorProvisioningUnknownIpLocation = uint32(C.ERROR_provisioning_unknown_ip_location)
var ErrorProvisioningInternalTriesExceeded = uint32(C.ERROR_provisioning_internal_tries_exceeded)
var ErrorProvisioningTooManySlotsRequested = uint32(C.ERROR_provisioning_too_many_slots_requested)
var ErrorProvisioningTooManyReserved = uint32(C.ERROR_provisioning_too_many_reserved)
var ErrorProvisioningCouldNotConnect = uint32(C.ERROR_provisioning_could_not_connect)
var ErrorProvisioningAuthServerNotConnected = uint32(C.ERROR_provisioning_auth_server_not_connected)
var ErrorProvisioningAuthDataTooLarge = uint32(C.ERROR_provisioning_auth_data_too_large)
var ErrorProvisioningAlreadyInitialized = uint32(C.ERROR_provisioning_already_initialized)
var ErrorProvisioningNotInitialized = uint32(C.ERROR_provisioning_not_initialized)
var ErrorProvisioningConnecting = uint32(C.ERROR_provisioning_connecting)
var ErrorProvisioningAlreadyConnected = uint32(C.ERROR_provisioning_already_connected)
var ErrorProvisioningNotConnected = uint32(C.ERROR_provisioning_not_connected)
var ErrorProvisioningIoError = uint32(C.ERROR_provisioning_io_error)
var ErrorProvisioningInvalidTimeout = uint32(C.ERROR_provisioning_invalid_timeout)
var ErrorProvisioningTs3serverNotFound = uint32(C.ERROR_provisioning_ts3server_not_found)
var ErrorProvisioningNoPermission = uint32(C.ERROR_provisioning_no_permission)
