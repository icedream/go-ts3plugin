package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

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

var wg sync.WaitGroup

var isFeedRunning bool
var customDeviceID string
var customDeviceShutdownC = make(chan interface{})
var sampleRate = 44100
var channels = 2

func startFeed() {
	logDebug("Checking if we need to start feed…")
	if isFeedRunning {
		logDebug("Feed already running, skipping.")
		return
	}
	isFeedRunning = true

	logDebug("Starting feed…")

	// capture feed
	wg.Add(1)
	go func() {
		defer wg.Done()

		logDebug("Now running capture feed loop")
		sampleBufferInterval := 100 * time.Millisecond
		sampleBuffers := int(float64(sampleRate) * (float64(sampleBufferInterval) / float64(time.Second)))
		totalSamplesGenerate := 0
		nextSampleBufferAt := time.Now()
		samples := make([]int16, sampleBuffers*channels)

		logDebug("Will generate %d samples every %s", sampleBuffers, sampleBufferInterval)

		for {
			for i := range samples {
				samples[i] = int16(rand.Intn(math.MaxInt16-math.MinInt16)+math.MinInt16) / 2
			}

			select {
			case <-customDeviceShutdownC:
				logDebug("Exiting capture feed loop")
				return
			default:
				<-time.After(nextSampleBufferAt.Sub(time.Now()))
				nextSampleBufferAt = nextSampleBufferAt.Add(sampleBufferInterval)

				totalSamplesGenerate += sampleBuffers
				logDebug("Generated samples: %d s", totalSamplesGenerate/sampleRate)

				if errCode := ts3plugin.Functions().ProcessCustomCaptureData(customDeviceID, samples); errCode != teamspeak.ErrorOK {
					errCodeStr, _ := ts3plugin.Functions().GetErrorMessage(errCode)
					logWarn("Failed to process custom capture data: %s", errCodeStr)
				}

				samples := make([]int16, 4096)
				if errCode := ts3plugin.Functions().AcquireCustomPlaybackData(customDeviceID, samples); errCode == teamspeak.ErrorSoundNoData {
					continue // we got no data, just continue
				} else if errCode != teamspeak.ErrorOK {
					errCodeStr, _ := ts3plugin.Functions().GetErrorMessage(errCode)
					logWarn("Failed to acquire playback data: %s", errCodeStr)
				}
			}
		}
	}()
}

func stopFeed() {
	logDebug("Checking if we need to stop feed…")
	if !isFeedRunning {
		logDebug("Feed already stopped, skipping.")
		return
	}

	logDebug("Stopping feed…")

	customDeviceShutdownC <- nil
	wg.Wait()

	isFeedRunning = false
}

func init() {
	ts3plugin.Name = "DevicesExamplePlugin"
	ts3plugin.Author = "Carl Kittelberger"
	ts3plugin.Version = "v0.0.0"

	customDeviceID = fmt.Sprintf("%sCustomCapture", ts3plugin.Name)

	ts3plugin.OnConnectedStatusChangeEvent = func(serverConnectionHandlerID uint64, newStatus int, errorNumber uint) {
		if teamspeak.ConnectStatus(newStatus) == teamspeak.StatusConnected {
			// startFeed()

			// // open playback device
			// if errCode := ts3plugin.Functions().OpenPlaybackDevice(serverConnectionHandlerID, "custom", customDeviceID); errCode != teamspeak.ErrorOK {
			// 	errCodeStr, _ := ts3plugin.Functions().GetErrorMessage(errCode)
			// 	logWarn("Failed to open playback device: %s", errCodeStr)
			// }

			// // open capture device
			// if errCode := ts3plugin.Functions().OpenCaptureDevice(serverConnectionHandlerID, "custom", customDeviceID); errCode != teamspeak.ErrorOK {
			// 	errCodeStr, _ := ts3plugin.Functions().GetErrorMessage(errCode)
			// 	logWarn("Failed to open capture device: %s", errCodeStr)
			// }
		}
	}

	ts3plugin.Init = func() (ok bool) {
		logDebug("Plugin starting up...")
		ok = true

		defer func() {
			if err := recover(); err != nil {
				logCritical("runtime panic: %s", err)
				ok = false
			}
		}()

		// logDebug("Fetching playback modes...")
		// playbackModes, errCode := ts3plugin.Functions().GetPlaybackModeList()
		// errMsg, _ := ts3plugin.Functions().GetErrorMessage(errCode)
		// logDebug("Fetch playback modes result: %s", errMsg)
		// if errCode != teamspeak.ErrorOK {
		// 	logWarn("Failed to retrieve playback mode list")
		// } else {
		// 	logDebug("Got %d results", len(playbackModes))
		// 	for _, modeID := range playbackModes {
		// 		logDebug("Fetching playback modes for %s...", modeID)
		// 		playbackDevices, errCode := ts3plugin.Functions().GetPlaybackDeviceList(modeID)
		// 		errMsg, _ = ts3plugin.Functions().GetErrorMessage(errCode)
		// 		logDebug("Fetch playback devices for %s result: %s", modeID, errMsg)
		// 		if errCode != teamspeak.ErrorOK {
		// 			logWarn("Failed to retrieve playback device list for mode %s", modeID)
		// 		} else {
		// 			for _, playbackDevice := range playbackDevices {
		// 				logWarn("%q", playbackDevice)
		// 			}
		// 		}
		// 	}
		// }

		// defaultCaptureMode, errCode := ts3plugin.Functions().GetDefaultCaptureMode()
		// if errCode != teamspeak.ErrorOK {
		// 	logWarn("Failed to retrieve default capture mode")
		// } else {
		// 	logDebug("Default capture mode: %s", defaultCaptureMode)
		// }

		// defaultCaptureDevice, errCode := ts3plugin.Functions().GetDefaultCaptureDevice(defaultCaptureMode)
		// if errCode != teamspeak.ErrorOK {
		// 	logWarn("Failed to retrieve default capture device")
		// } else {
		// 	logDebug("Default capture device: %q", defaultCaptureDevice)
		// }

		// defaultPlaybackMode, errCode := ts3plugin.Functions().GetDefaultPlayBackMode()
		// if errCode != teamspeak.ErrorOK {
		// 	logWarn("Failed to retrieve default playback mode")
		// } else {
		// 	logDebug("Default playback mode: %s", defaultPlaybackMode)
		// }

		// defaultPlaybackDevice, errCode := ts3plugin.Functions().GetDefaultPlaybackDevice(defaultPlaybackMode)
		// if errCode != teamspeak.ErrorOK {
		// 	logWarn("Failed to retrieve default playback device")
		// } else {
		// 	logDebug("Default playback device: %q", defaultPlaybackDevice)
		// }

		// Set up custom device
		if errCode := ts3plugin.Functions().RegisterCustomDevice(
			customDeviceID,
			fmt.Sprintf("%s Custom Capture", ts3plugin.Name),
			sampleRate,
			channels,
			sampleRate,
			channels); errCode != teamspeak.ErrorOK {
			ok = false
			logError("Failed to initialize custom audio device")
			return
		}

		startFeed()

		logDebug("Plugin started up.")

		return
	}

	ts3plugin.Shutdown = func() {
		logDebug("Plugin shutting down…")

		stopFeed()

		ts3plugin.Functions().UnregisterCustomDevice(customDeviceID)

		logDebug("Plugin shut down.")
	}
}

// This will never be run!
func main() {
	fmt.Println("=======================================")
	fmt.Println("This is a TeamSpeak3 plugin, do not run this as a CLI application!")
	fmt.Println("Args were: ", os.Args)
	fmt.Println("=======================================")
}
