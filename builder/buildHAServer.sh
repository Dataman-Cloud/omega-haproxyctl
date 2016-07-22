#!/bin/bash

S6_OVERLAY_VERSION=v1.18.1.3

echo "" > /etc/apk/repositories
echo  http://mirrors.ustc.edu.cn/alpine/v3.4/main/  >>  /etc/apk/repositories
echo  http://mirrors.ustc.edu.cn/alpine/v3.4/community/  >>  /etc/apk/repositories


apk update && \
  apk add  --no-cache libnl3 libnl3-cli git bash go haproxy net-tools wget iptables iproute2 \
  && wget https://github.com/just-containers/s6-overlay/releases/download/${S6_OVERLAY_VERSION}/s6-overlay-amd64.tar.gz --no-check-certificate -O /tmp/s6-overlay.tar.gz \
  && tar xvfz /tmp/s6-overlay.tar.gz -C / \
  && rm -f /tmp/s6-overlay.tar.gz \
  && apk del wget

export GOROOT=/usr/lib/go
export GOPATH=/gopath
export GOBIN=/gopath/bin
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

cd /gopath/src/github.com/Dataman-Cloud/HAServer
go build -o HAServer && \
mkdir -p /var/haserver && \
cp /gopath/src/github.com/Dataman-Cloud/HAServer/HAServer /var/haserver/HAServer && \
mkdir -p /run/haproxy && \
mkdir -p /var/log/supervisor &&\
cd /

rm -rf /tmp/* /var/tmp/*
rm -f /etc/ssh/ssh_host_*
rm -rf /gopath
rm -rf /usr/lib/go
