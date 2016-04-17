#!/bin/sh -eux

go get -u github.com/constabulary/gb/...
go get -u code.palmstonegames.com/gb-gae

cd `dirname $0`

PROJECT=mstssk-jp
VERSION=`git rev-parse --abbrev-ref HEAD`

gb gae appcfg update \
  --application=$PROJECT \
  --version=$VERSION \
  src/
