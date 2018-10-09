@echo off
set ts3dir=%appdata%\TS3Client
go build -v -buildmode=c-shared -o "%ts3dir%\plugins\example.dll"
