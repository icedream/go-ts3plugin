@echo off
set ts3dir=%appdata%\TS3Client
set pluginname=devices-example-ts3plugin
set "dllname=%ts3dir%\plugins\%pluginname%.dll"

go build -v -buildmode=c-shared -o "%dllname%"
echo Installed to: %dllname%
