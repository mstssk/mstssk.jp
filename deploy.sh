#!/bin/sh -eux

go get -u github.com/constabulary/gb/...
go get -u github.com/PalmStoneGames/gb-gae

cd `dirname $0`

PROJECT=mstssk-jp
VERSION=`git rev-parse --abbrev-ref HEAD`

gb gae appcfg update \
  --application=$PROJECT \
  --version=$VERSION \
  --oauth2_access_token $(gcloud auth print-access-token 2> /dev/null) \
  src/
