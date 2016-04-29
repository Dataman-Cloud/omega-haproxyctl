FROM alpine:3.2
MAINTAINER lhan <lhan@dataman-inc.com>

RUN mkdir -p /config /etc/default_haproxy

ADD config/production.json /config/production.json
ADD config/default_haproxy.cfg /etc/default_haproxy/default_haproxy.cfg

ADD . /gopath/src/github.com/Dataman-Cloud/HAServer
ADD files /usr/share/haproxy
ADD builder/supervisord.conf /etc/supervisord.conf
ADD builder/run.sh /run.sh
ADD builder/buildHAServer.sh /buildHAServer.sh
WORKDIR /

RUN sh /buildHAServer.sh

EXPOSE 5004 5091

CMD sh /run.sh   
