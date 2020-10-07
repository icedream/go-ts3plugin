#!/bin/sh -e
ts3dir="$HOME/.ts3client"
pluginname="devices-example-ts3plugin"
mkdir -p "${ts3dir}/plugins"
go build \
    -buildmode=c-shared \
    -o "${ts3dir}/plugins/${pluginname}.so" \
    "$@"
