#!/bin/bash
set -euo pipefail
export GOWORK=off
go build -tags=jsoniter -o app ./cmd/app
