#!/usr/bin/env bash
set -ex

network="nqqdbcni"
pod="nqqmysql"
dbs1="nqqdbs1"
ns="nqqdc1ns"

podman network create $network
podman pod create --name $pod --share net,cgroup,ipc
podman volume create mysqldata

# ip_ns=$(podman inspect $ns|jq -r '.[].NetworkSettings.Networks."nqqdc1ns".IPAddress')

podman run -dt --name $dbs1 --hostname $dbs1 --pod $pod --network $network -v mysqldata:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=nqq123 \
    mariadb:latest 

podman exec $dbs1 ss -lntp

#register mysql service

ip_s1=$(podman inspect $dbs1|jq -r '.[].NetworkSettings.Networks."'$network'".IPAddress')

podman exec $ns consul services register -id "db1" -address $ip_s1 -port 3306 -name "mysql"

