# TeamSpeak3 Plugin SDK for Golang

This library allows developers to write plugins for the TeamSpeak3 client in
Go.

Note that its current functionality is limited. A huge part has not been 
implemented yet but this should be enough to interact with the chat or set up
automations.

## Requirements

- Go 1.18 or newer (older versions have not been tested against).
- A working C compiler suite.
- Your own unpacked copy of [Version 24 the TeamSpeak3 Plugin SDK](https://github.com/TeamSpeak-Systems/ts3client-pluginsdk/tree/a39d50383b4a023941c31c08fe0c9766b149ed01).
    To download it just click "Code" and then "Download ZIP" there.
- Having the `include/` directory exported as part of environment variable
    `CGO_CPPFLAGS` (e.g. `-I/path/to/your/sdk/copy/include/`).
