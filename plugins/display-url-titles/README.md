# Example TS3 plugin

This plugin is just implementing random message boxes and console output.

## Building the plugin

Make sure you use at least Go 1.11 with experimental module support.

```
# Windows:
go build -v -buildmode=c-shared -o test-ts3-plugin.dll

# Linux:
go build -v -buildmode=c-shared -o test-ts3-plugin.so
```
