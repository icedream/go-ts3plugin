package teamlog

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
	LogTypeNone         = LogType(uint(C.LogType_NONE))
	LogTypeFile         = LogType(uint(C.LogType_FILE))
	LogTypeConsole      = LogType(uint(C.LogType_CONSOLE))
	LogTypeUserlogging  = LogType(uint(C.LogType_USERLOGGING))
	LogTypeNoNetlogging = LogType(uint(C.LogType_NO_NETLOGGING))
	LogTypeDatabase     = LogType(uint(C.LogType_DATABASE))
	LogTypeSyslog       = LogType(uint(C.LogType_SYSLOG))
)

type LogLevel uint

const (
	LogLevelCritical = LogLevel(uint(C.LogLevel_CRITICAL))
	LogLevelError    = LogLevel(uint(C.LogLevel_ERROR))
	LogLevelWarning  = LogLevel(uint(C.LogLevel_WARNING))
	LogLevelDebug    = LogLevel(uint(C.LogLevel_DEBUG))
	LogLevelInfo     = LogLevel(uint(C.LogLevel_INFO))
	LogLevelDevel    = LogLevel(uint(C.LogLevel_DEVEL))
)
