#!/bin/bash

set -e -u -x

ls
pwd

go test ./...
