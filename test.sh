#!/bin/sh -eux

cd `dirname $0`
cd ./src/

goimports -w .
gb generate app
go tool vet .
golint ./...

gb gae test ./... $@
