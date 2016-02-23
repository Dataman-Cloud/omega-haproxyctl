docker run -it \
    --net=host \
    -v /data/haproxy:/etc/haproxy \
    -v /data/run/haproxy:/run/haproxy \
    -e BIND=":5004" \
    -e CONFIG_PATH="config/production.gateway.json" \
    --name=haproxy \
    registry.shurenyun.com/haproxy-1.5.4:omega.v2.1
