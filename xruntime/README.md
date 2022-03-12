# :zap: xruntime
xruntime is a easy way to get runtime info

## Installation
``` sh
go get -u github.com/hkloudou/xlib/xruntime
```

## use build script
``` Makefile
build:
    sh ${shell go env GOMODCACHE}/github.com/hkloudou/xlib@v1.0.49/scripts/gobuild.sh appNameDemo $(shell go env GOPATH)/bin/xruntimeDemo
```