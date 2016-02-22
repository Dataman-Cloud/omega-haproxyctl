FROM alpine:3.2
MAINTAINER lhan <lhan@dataman-inc.com>

RUN mkdir -p /config

ADD config/haproxy_template.gateway.centos.cfg /config/haproxy_template.gateway.cfg
ADD config/production.gateway.json /config/production.gateway.json
ADD config/production.example.json /config/production.example.json

ADD . /gopath/src/github.com/Dataman-Cloud/HAServer
#ADD haproxy /usr/share/haproxy
ADD builder/supervisord.conf /etc/supervisord.conf
ADD builder/run.sh /run.sh
ADD builder/buildHAServer.sh /buildHAServer.sh
WORKDIR /

RUN sh /buildHAServer.sh

VOLUME /var/log/supervisor
VOLUME /config
VOLUME /etc/haproxy

EXPOSE 80

CMD sh /run.sh   
