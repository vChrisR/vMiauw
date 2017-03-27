#!/bin/bash

set -e -u -x

export GOPATH=$PWD/vMiauw
go test ./...
