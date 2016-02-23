#!/bin/bash
haproxy -f /etc/default_haproxy/default_haproxy.cfg -p /var/run/haproxy.pid
if [ -e /etc/haproxy/haproxy.cfg ]; then
  haproxy -c -f /etc/haproxy/haproxy.cfg
  if [ $? -eq 0 ]; then
    haproxy -f /etc/haproxy/haproxy.cfg -p /var/run/haproxy.pid -D -sf $(cat /var/run/haproxy.pid)
  fi
fi

/usr/bin/supervisord
