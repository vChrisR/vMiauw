#!/bin/bash

set -e -u -x

cd ./vmiauw

go get ./...
go test ./...
