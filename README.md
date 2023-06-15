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

## License

This library has been licensed under the BSD 3-Clause License in hopes to keep
it entirely compatible with the original TeamSpeak3 Plugin SDK.

There are known open-source licensing complications in regards to the
TeamSpeak3 Plugin SDK itself (see [ts3client-pluginsdk#3]) which I am trying
to comply with on best efforts and with best intentions in mind. If a
violation of copyright or license exists in this repository, do contact me
either via GitHub issues or via the e-mail used for my Git commits.

[ts3client-pluginsdk#3]: https://github.com/TeamSpeak-Systems/ts3client-pluginsdk/issues/3
