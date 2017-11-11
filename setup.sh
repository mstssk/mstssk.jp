#!/bin/sh -eux

cd `dirname $0`

go get -u github.com/constabulary/gb/...
go get -u github.com/PalmStoneGames/gb-gae

go get -u sourcegraph.com/sqs/goreturns
go get -u github.com/golang/lint/golint
