#!/bin/sh -e
ts3dir="$HOME/.ts3client"
mkdir -p "${ts3dir}/plugins"
go build -v -buildmode=c-shared -o "${ts3dir}/plugins/example.so"
# go build -v -buildmode=c-shared -ldflags '-X "main.Name=First Go Test plugin"'  -o "${ts3dir}/plugins/example.so"
# go build -v -buildmode=c-shared -ldflags '-X "main.Name=Second Go Test plugin"' -o "${ts3dir}/plugins/example2.so"
