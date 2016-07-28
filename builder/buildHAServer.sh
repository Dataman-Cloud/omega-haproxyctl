#!/bin/bash

S6_OVERLAY_VERSION=v1.18.1.3

echo "" > /etc/apk/repositories
echo  http://mirrors.ustc.edu.cn/alpine/v3.4/main/  >>  /etc/apk/repositories
echo  http://mirrors.ustc.edu.cn/alpine/v3.4/community/  >>  /etc/apk/repositories


apk update && \
  apk add  --no-cache libnl3 libnl3-cli git bash go haproxy net-tools wget iptables iproute2 \
  && tar xvfz /tmp/s6-overlay-amd64.tar.gz -C / \
  && rm -f /tmp/s6-overlay-amd64.tar.gz  \
  && apk del wget

export GOROOT=/usr/lib/go
export GOPATH=/gopath
export GOBIN=/gopath/bin
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

cd /gopath/src/github.com/Dataman-Cloud/omega-haproxyctl
go build -o omega-haproxyctl && \
mkdir -p /var/omega-haproxyctl && \
cp /gopath/src/github.com/Dataman-Cloud/omega-haproxyctl/omega-haproxyctl /var/omega-haproxyctl/omega-haproxyctl && \
mkdir -p /run/haproxy && \
mkdir -p /var/log/supervisor &&\
cd /

rm -rf /tmp/* /var/tmp/*
rm -f /etc/ssh/ssh_host_*
rm -rf /gopath
rm -rf /usr/lib/go
