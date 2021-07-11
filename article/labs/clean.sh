#!/usr/bin/env bash

network="nqqdbcni"
pod="nqqmysql"

podman network rm $network -f
podman pod rm $pod -f
podman volume rm mysqldata

