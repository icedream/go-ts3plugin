#!/bin/sh -e
ts3dir="$HOME/.ts3client"
mkdir -p "${ts3dir}/plugins"
go build -v -buildmode=c-shared -o "${ts3dir}/plugins/example.so"
