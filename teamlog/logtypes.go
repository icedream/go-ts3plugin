package teamlog

//go:generate stringer -type=LogType,LogLevel -output generated_logtypes_string.go

/*
#cgo CFLAGS: -I../pluginsdk/include -I.
#include "teamlog/logtypes.h"
*/
import "C"

/******************************************************************************
From logtypes.h
*******************************************************************************/

type LogType uint

const (
	None         = LogType(uint(C.LogType_NONE))
	File         = LogType(uint(C.LogType_FILE))
	Console      = LogType(uint(C.LogType_CONSOLE))
	Userlogging  = LogType(uint(C.LogType_USERLOGGING))
	NoNetlogging = LogType(uint(C.LogType_NO_NETLOGGING))
	Database     = LogType(uint(C.LogType_DATABASE))
	Syslog       = LogType(uint(C.LogType_SYSLOG))
)

type LogLevel uint

const (
	Critical = LogLevel(uint(C.LogLevel_CRITICAL))
	Error    = LogLevel(uint(C.LogLevel_ERROR))
	Warning  = LogLevel(uint(C.LogLevel_WARNING))
	Debug    = LogLevel(uint(C.LogLevel_DEBUG))
	Info     = LogLevel(uint(C.LogLevel_INFO))
	Devel    = LogLevel(uint(C.LogLevel_DEVEL))
)
