#! /bin/bash

set -e

IMAGE="omega-haproxyctl"
VERSION="$(cat VERSION)"


docker run \
  -e MARTINI_ENV=production \
  --privileged --rm \
  --net=host ${IMAGE}:${VERSION}
