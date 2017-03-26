#!/bin/sh

set -e

cd ./vmiauw

go get ./...
go test ./...
