#!/usr/bin/env bash

etcdcontainername=etcd1
etcdclientport=2379
etcdpodname=podetcd
etcdname=node1

podman pod rm -f ${etcdpodname} 
# podman pod create --name ${etcdpodname} --network cni-podman1 
podman run -dt --name ${etcdcontainername} --pod new:${etcdpodname} -p ${etcdclientport}:${etcdclientport} --network cni-podman1 hub.iohttps.com/etcd:test bash


etcdip=$(podman inspect etcd1 |jq -r '.[].NetworkSettings.Networks."cni-podman1".IPAddress')
host=http://${etcdip}:${etcdclientport}
peer_host=http://${etcdip}:2380
local_host=http://localhost:${etcdclientport}

podman exec -dt ${etcdcontainername} rm -rf /var/etcd/data/${etcdname}.etcd

podman exec -dt ${etcdcontainername} etcd --name ${etcdname} --data-dir /var/etcd/data/{etcdname}.etcd --initial-advertise-peer-urls ${peer_host} --listen-peer-urls ${peer_host} --listen-client-urls ${local_host},${host} --advertise-client-urls ${host} --initial-cluster="${etcdname}=${peer_host}" --initial-cluster-token etcd-cluster1

volume="realworlddb"
name=mysql
port=3306
password=nqq123
user=root
podname=podmariadb

# podman volume create ${volume}
podman pod rm -f ${podname} 

podman run -d --name ${name} --pod new:${podname} --network cni-podman1 -p ${port}:3306 -v ${volume}:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=${password} docker.io/library/mariadb  

mysqlip=$(ifconfig |grep eth0 -A 1 | awk 'NR==2{print $2}')

etcdctl --endpoints http://127.0.0.1:${etcdclientport} put mysql.ip ${mysqlip}
etcdctl put mysql.port ${port}
etcdctl put mysql.user ${user}
etcdctl put mysql.password ${password}
