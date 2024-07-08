#!/bin/bash

# This script generates a WatchYourPorts config from Docker containers.

# HOW TO USE
#    1. Run this script on a server, where Docker is installed:
#   ./docker-export.sh $ADDR
#    $ADDR is IP or domain name of the server, without http(s):// prefix
#    It will be used to ping ports
#    2. Paste the output to hosts.yaml file in WYP config dir
#    3. You can add ad many servers to hosts.yaml, as you want

docker ps -a --format "{{.Names}}">/tmp/wyp-docker.txt

echo $1':'
echo '    name:'
echo '    addr: '$1
echo '    portmap:'

while read NAME; do
    PORT=`docker inspect $NAME | grep HostPort | sed '1!d;s/"HostPort": //;s/,//;s/"//g'`

    if [ ${#PORT} -ge 1 ]; then
        echo '       '$PORT':'
        echo '            name: '$NAME
        echo '            port:'$PORT
        echo '            state: false'
        echo '            watch: true'
    fi
done </tmp/wyp-docker.txt