#!/bin/bash

set -e -u -x

export GOPATH=$(pwd)/gopath
cd gopath/src/github.com/vchrisr/vMiauw/

godep go test ./...
