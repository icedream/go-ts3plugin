package main

import (
	"fmt"
	"os"

	ts3plugin "github.com/icedream/go-ts3plugin"
	"github.com/icedream/go-ts3plugin/teamlog"
	"github.com/icedream/go-ts3plugin/teamspeak"
)

func log(msg string, severity teamlog.LogLevel) {
	if ts3plugin.Functions() != nil {
		ts3plugin.Functions().LogMessage(
			msg,
			severity, ts3plugin.Name, 0)
	}
}

func logDebug(msg string, args ...interface{}) {
	log(fmt.Sprintf(msg, args...), teamlog.Debug)
}

func logInfo(msg string, args ...interface{}) {
	log(fmt.Sprintf(msg, args...), teamlog.Info)
}

func logWarn(msg string, args ...interface{}) {
	log(fmt.Sprintf(msg, args...), teamlog.Warning)
}

func logError(msg string, args ...interface{}) {
	log(fmt.Sprintf(msg, args...), teamlog.Error)
}

func logCritical(msg string, args ...interface{}) {
	log(fmt.Sprintf(msg, args...), teamlog.Critical)
}

func init() {
	ts3plugin.Name = "DevicesExamplePlugin"
	ts3plugin.Author = "Carl Kittelberger"
	ts3plugin.Version = "v0.0.0"

	ts3plugin.Init = func() (ok bool) {
		logDebug("Plugin starting up...")
		ok = true

		defer func() {
			if err := recover(); err != nil {
				logCritical("runtime panic: %s", err)
				ok = false
			}
		}()

		logDebug("Fetching playback modes...")
		playbackModes, errCode := ts3plugin.Functions().GetPlaybackModeList()
		errMsg, _ := ts3plugin.Functions().GetErrorMessage(errCode)
		logDebug("Fetch playback modes result: %s", errMsg)
		if errCode != teamspeak.ErrorOK {
			logWarn("Failed to retrieve playback mode list")
		} else {
			logDebug("Got %d results", len(playbackModes))
			for _, modeID := range playbackModes {
				logDebug("Fetching playback modes for %s...", modeID)
				playbackDevices, errCode := ts3plugin.Functions().GetPlaybackDeviceList(modeID)
				errMsg, _ = ts3plugin.Functions().GetErrorMessage(errCode)
				logDebug("Fetch playback devices for %s result: %s", modeID, errMsg)
				if errCode != teamspeak.ErrorOK {
					logWarn("Failed to retrieve playback device list for mode %s", modeID)
				} else {
					for _, playbackDevice := range playbackDevices {
						logWarn("%q", playbackDevice)
					}
				}
			}
		}

		defaultCaptureMode, errCode := ts3plugin.Functions().GetDefaultCaptureMode()
		if errCode != teamspeak.ErrorOK {
			logWarn("Failed to retrieve default capture mode")
		} else {
			logDebug("Default capture mode: %s", defaultCaptureMode)
		}

		defaultCaptureDevice, errCode := ts3plugin.Functions().GetDefaultCaptureDevice(defaultCaptureMode)
		if errCode != teamspeak.ErrorOK {
			logWarn("Failed to retrieve default capture device")
		} else {
			logDebug("Default capture device: %q", defaultCaptureDevice)
		}

		defaultPlaybackMode, errCode := ts3plugin.Functions().GetDefaultPlayBackMode()
		if errCode != teamspeak.ErrorOK {
			logWarn("Failed to retrieve default playback mode")
		} else {
			logDebug("Default playback mode: %s", defaultPlaybackMode)
		}

		defaultPlaybackDevice, errCode := ts3plugin.Functions().GetDefaultPlaybackDevice(defaultPlaybackMode)
		if errCode != teamspeak.ErrorOK {
			logWarn("Failed to retrieve default playback device")
		} else {
			logDebug("Default playback device: %q", defaultPlaybackDevice)
		}

		logDebug("Plugin started up.")

		return
	}
}

// This will never be run!
func main() {
	fmt.Println("=======================================")
	fmt.Println("This is a TeamSpeak3 plugin, do not run this as a CLI application!")
	fmt.Println("Args were: ", os.Args)
	fmt.Println("=======================================")
}
