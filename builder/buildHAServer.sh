#!/bin/bash
apk update && apk add curl git bash go haproxy supervisor net-tools && rm -rf /var/cache/apk/*
export GOROOT=/usr/lib/go
export GOPATH=/gopath
export GOBIN=/gopath/bin
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

cd /gopath/src/github.com/Dataman-Cloud/HAServer
go get github.com/tools/godep && \
go build && \
mkdir -p /var/haserver && \
cp /gopath/src/github.com/Dataman-Cloud/HAServer/HAServer /var/haserver/HAServer && \
mkdir -p /run/haproxy && \
mkdir -p /var/log/supervisor &&\
cd /

rm -rf /tmp/* /var/tmp/*
rm -f /etc/ssh/ssh_host_*
rm -rf /gopath
rm -rf /usr/lib/go
