FROM haproxy:1.5.18-alpine
MAINTAINER Xiao Deshi <dsxiao@dataman-inc.com>

RUN mkdir -p /config /etc/default_haproxy

COPY config/production.json /config/production.json
COPY config/default_haproxy.cfg /etc/default_haproxy/default_haproxy.cfg

COPY . /gopath/src/github.com/Dataman-Cloud/HAServer
COPY files /usr/share/haproxy
COPY builder/buildHAServer.sh /buildHAServer.sh
WORKDIR /

RUN sh /buildHAServer.sh

EXPOSE 5004 5091

# Add in base configuration
COPY root /

ENTRYPOINT ["/init"]
CMD []
