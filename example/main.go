package main

import (
	"fmt"
	"runtime/debug"

	teamspeak_plugin "github.com/icedream/go-ts3plugin"
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
	MessageBox(Name, "Hello, this is Test Go Plugin running init()!", MB_ICONINFORMATION|MB_OK)

	teamspeak_plugin.Name = "Test Go Plugin"
	teamspeak_plugin.Author = "Carl Kittelberger"
	teamspeak_plugin.Version = "0.0.0"

	teamspeak_plugin.Init = func() int {
		defer catchPanic()

		version, ok := teamspeak_plugin.Functions().GetClientLibVersion()
		if ok {
			MessageBox(Name,
				fmt.Sprintf("TS3::Init - plugin ID %s running on %s!",
					teamspeak_plugin.GetPluginID(),
					version),
				MB_ICONINFORMATION|MB_OK)
		} else {
			MessageBox(Name, "Could not get client lib version", MB_ICONERROR|MB_OK)
		}

		return 0
	}

	teamspeak_plugin.Shutdown = func() {
		defer catchPanic()
		defer freeMessageBoxLibraries()

		MessageBox(Name, "TS3::Shutdown", MB_ICONINFORMATION|MB_OK)
	}

}

func main() {
	defer catchPanic()
	MessageBox(Name, "Hello, this is Test Go Plugin running main()!", MB_ICONINFORMATION|MB_OK)
}
