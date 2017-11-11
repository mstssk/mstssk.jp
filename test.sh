#!/bin/sh -eux

cd `dirname $0`
cd ./src/

goreturns -w .
gb generate app
# go tool vet .
golint ./...

gb gae test ./... $@
