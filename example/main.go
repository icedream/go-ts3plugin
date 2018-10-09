package main

import (
	"fmt"
	"runtime/debug"

	"github.com/icedream/go-ts3plugin"
	"github.com/icedream/go-ts3plugin/teamlog"
	"github.com/icedream/go-ts3plugin/teamspeak"
)

const (
	Name    = "Test Go Plugin"
	Author  = "Carl Kittelberger"
	Version = "0.0.0"
)

func catchPanic() {
	if err := recover(); err != nil {
		MessageBox("PANIC!!!", fmt.Sprintf("%s\n%s", err, string(debug.Stack())), MB_ICONERROR|MB_OK)
	}
}

func init() {
	ts3plugin.Name = "Test Go Plugin"
	ts3plugin.Author = "Carl Kittelberger"
	ts3plugin.Version = "0.0.0"

	ts3plugin.Init = func() int {
		defer catchPanic()

		fmt.Println("############################################")
		fmt.Println(ts3plugin.Name)
		fmt.Println("by", ts3plugin.Name)
		fmt.Println(ts3plugin.Version)
		fmt.Println("############################################")

		version, errorCode := ts3plugin.Functions().GetClientLibVersion()
		if errorCode == teamspeak.ErrorOK {
			ts3plugin.Functions().LogMessage(
				fmt.Sprintf("TS3::Init - plugin ID %s running on %s!",
					ts3plugin.GetPluginID(),
					version),
				teamlog.LogLevelInfo, ts3plugin.Name, 0)
		} else {
			ts3plugin.Functions().LogMessage(
				"Could not get client lib version",
				teamlog.LogLevelError, ts3plugin.Name, 0)
		}

		return 0
	}

	ts3plugin.Shutdown = func() {
		defer catchPanic()
		defer freeMessageBoxLibraries()
	}

}

// This will never be run!
func main() {}
