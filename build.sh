#! /bin/bash

set -e

# include haproxy binary
IMAGE="omega-haproxyctl"
VERSION="$(cat VERSION)"
HOST="catalog.shurenyun.com"
NAMESPACE="library"

docker build --no-cache -t ${IMAGE}:${VERSION} -f Dockerfile.builder .

docker tag ${IMAGE}:${VERSION} $(HOST)/$(NAMESPACE)/${IMAGE}:latest
