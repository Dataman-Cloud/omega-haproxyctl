#! /bin/bash

set -e

# include haproxy binary
IMAGE="omega-haproxyctl"
VERSION="$(cat VERSION)"

docker build --no-cache -t ${IMAGE}:${VERSION} -f Dockerfile.builder .

docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest
