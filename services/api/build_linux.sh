#!/bin/bash
set -euo pipefail
export GOWORK=off
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -tags=jsoniter -o app ./cmd/app
