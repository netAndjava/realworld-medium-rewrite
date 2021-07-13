#!/usr/bin/env bash
set -ex

# go on use set -x checking you step command, -e to stop at a error

dc="nqqdc1"
network=$dc"cni"
pod=$dc
s1=$dc"s1"  # nqqs1
# c1=$dc"c1"
# guest1=$dc"guest1"
# guest2=$dc"guest2"
ns=$dc"ns"
domain="nqq"

podman network create $network
# remove uts namespaces, we want different hostname
podman pod create --name $pod --share net,cgroup,ipc

# podman cp ./consul.d/consul.hcl $s1:/etc/consul.d/
podman run -dt --name $s1 --hostname $s1 --pod $pod --network $network consul:beta bash
ip_s1=$(podman inspect $s1|jq -r '.[].NetworkSettings.Networks."'$network'".IPAddress')

create_server(){
    podman exec -dt $1 consul agent \
        -server \
        -data-dir=/opt/consul \
        -config-dir=/etc/consul.d \
        -datacenter=$dc \
        -bootstrap-expect=$2
}

create_server $s1 1

podman run -dt --name $ns --hostname $ns --pod $pod --network $network consul:beta bash
ip_ns=$(podman inspect $ns|jq -r '.[].NetworkSettings.Networks."'$network'".IPAddress')
podman exec -dt -u root $ns consul agent -config-dir=/etc/consul.d \
    -datacenter=$dc \
    -dns-port 53 \
    -enable-local-script-checks=true \
    -enable-script-checks=true \
    -domain=$domain\
    -retry-join $ip_s1 \
    -recursor=223.5.5.5 -recursor=223.6.6.6
podman exec $ns consul services register -id "ns" -address $ip_ns -port 53 -name "ns"
podman exec $ns consul catalog services

# podman run -dt --name $c1 --dns $ip_ns --dns-search "service.${domain}" --hostname $c1 --pod $pod --network $network consul:beta bash
# podman exec -dt $c1 consul agent -config-dir=/etc/consul.d -data-dir=/opt/consul -datacenter=$dc
#
# podman exec $c1 consul join $ip_s1
#
# podman exec $c1 consul members list

create_guest(){

    name=$1
    port=$2
    id=$3
    podman run -dt --name $name --dns $ip_ns --dns-search "service${domain}" --hostname $name --pod $pod --network $network ubuntu:dev bash

    podman exec -dt $name python3 -m http.server $port

    ip_guest=$(podman inspect $name|jq -r '.[].NetworkSettings.Networks."'$network'".IPAddress')

    podman exec $ns consul services register -id $id -address $ip_guest -port $port -name web
}

# create_guest $guest1 8080 web1 
# create_guest $guest2 8030 web2
#
# podman exec $guest1 dig web.service.nqq
# podman exec $guest1 dig SRV web.service.nqq



