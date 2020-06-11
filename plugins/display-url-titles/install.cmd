@echo off
set ts3dir=%appdata%\TS3Client
set pluginname=display-url-titles

go build -v -buildmode=c-shared -o "%ts3dir%\plugins\%pluginname%.dll"
