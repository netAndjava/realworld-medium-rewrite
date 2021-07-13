#!/usr/bin/env bash
set -ex

network="nqqcni"
pod="realworld-article"
s1="realworld-article-s1"
ns="nqqdc1ns"

ip_ns=$(podman inspect $ns|jq -r '.[].NetworkSettings.Networks."nqqdc1cni".IPAddress')

podman network create $network
podman pod create --name $pod --share net,cgroup,ipc

#1. create container and run program
podman run -dt --name $s1 --hostname $s1 --network $network --pod $pod --dns $ip_ns --dns-search "service.nqq"  hub.iohttps.com/article:dev bash

podman exec -dt $s1 article -db /var/etcd/mysql.toml -consul /var/etcd/consul.toml -config /var/etcd/dev.toml 
#2. check if article service start 

podman exec -it $s1 ss -lntp
