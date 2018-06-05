#!/bin/bash
export PROJ="github.com/networld-to/homecontrol"
go build -o build/homecontrol_server \
  -ldflags "-s -w -X '${PROJ}/version.Version=$(git describe --dirty --always)'\
  -X '${PROJ}/version.Build=$(date -u '+%Y-%m-%dT%H:%M:%SZ')'" server/main.go

go build -o build/homecontrol_client \
  -ldflags "-s -w -X '${PROJ}/version.Version=$(git describe --dirty --always)'\
  -X '${PROJ}/version.Build=$(date -u '+%Y-%m-%dT%H:%M:%SZ')'" client/main.go
