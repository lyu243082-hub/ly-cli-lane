#!/bin/bash
set -e
export PATH=/usr/local/go/bin:$PATH
export GOPROXY=https://goproxy.byted.org,https://goproxy.cn,direct
export GOSUMDB=sum.golang.google.cn
go mod download
go mod verify
go build -o bootstrap ./init.go
