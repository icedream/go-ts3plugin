package ts3plugin

/*
#cgo CFLAGS: -I./pluginsdk/include -I.

#include <stdio.h> // uint types

#include "plugin_definitions.h"
*/
import "C"

/* Return values for ts3plugin_offersConfigure */
type PluginConfigureOffer int

const (
	PLUGIN_OFFERS_NO_CONFIGURE         = 0 /* Plugin does not implement ts3plugin_configure */
	PLUGIN_OFFERS_CONFIGURE_NEW_THREAD     /* Plugin does implement ts3plugin_configure and requests to run this function in an own thread */
	PLUGIN_OFFERS_CONFIGURE_QT_THREAD      /* Plugin does implement ts3plugin_configure and requests to run this function in the Qt GUI thread */
)

type PluginMessageTarget int

const (
	PLUGIN_MESSAGE_TARGET_SERVER = 0
	PLUGIN_MESSAGE_TARGET_CHANNEL
)

type PluginItemType int

const (
	PLUGIN_SERVER = 0
	PLUGIN_CHANNEL
	PLUGIN_CLIENT
)

type PluginMenuType int

const (
	PLUGIN_MENU_TYPE_GLOBAL = 0
	PLUGIN_MENU_TYPE_CHANNEL
	PLUGIN_MENU_TYPE_CLIENT
)

type PluginMenuItem struct {
	MenuType PluginMenuType
	Id       int
	Text     [C.PLUGIN_MENU_BUFSZ]rune
	Icon     [C.PLUGIN_MENU_BUFSZ]rune
}

type PluginHotkey struct {
	Keyword     [C.PLUGIN_HOTKEY_BUFSZ]rune
	Description [C.PLUGIN_HOTKEY_BUFSZ]rune
}

type PluginBookmarkItem struct {
	Name     string
	IsFolder bool
	reserved [3]byte
	Uuid     string
	Folder   PluginBookmarkList
}

type PluginBookmarkList struct {
	Items []PluginBookmarkItem
}

type PluginGuiProfile int

const (
	PLUGIN_GUI_SOUND_CAPTURE PluginGuiProfile = 0
	PLUGIN_GUI_SOUND_PLAYBACK
	PLUGIN_GUI_HOTKEY
	PLUGIN_GUI_SOUNDPACK
	PLUGIN_GUI_IDENTITY
)

type PluginConnectTab int

const (
	PLUGIN_CONNECT_TAB_NEW PluginConnectTab = 0
	PLUGIN_CONNECT_TAB_CURRENT
	PLUGIN_CONNECT_TAB_NEW_IF_CURRENT_CONNECTED
)
