package ts3plugin

var (
	Name, Version, Author, Description string
	ApiVersion                         int = 22
)

var (
	CommandKeyword string
	InfoTitle      string
)

var (
	RequestAutoload = false
)

var (
	// Will be set by TeamSpeak after plugin is loaded.
	pluginID string
)

var (
	Init                           func() int
	Shutdown                       func()
	OffersConfigure                func() PluginConfigureOffer
	Configure                      func(handle byte, qParentWidget byte)
	ProcessCommand                 func(serverConnectionHandlerID uint64, command string) bool
	CurrentServerConnectionChanged func(serverConnectionHandlerID uint64)
)

var (
	functions *TS3Functions
)

func Functions() *TS3Functions {
	return functions
}

func GetPluginID() string {
	return pluginID
}
