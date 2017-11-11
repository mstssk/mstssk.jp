#!/bin/sh -eux

cd `dirname $0`

PROJECT=mstssk-jp
VERSION=`git rev-parse --abbrev-ref HEAD`

set +x
gb gae appcfg update \
  --application=$PROJECT \
  --version=$VERSION \
  --oauth2_access_token $(gcloud auth print-access-token 2> /dev/null) \
  src/
