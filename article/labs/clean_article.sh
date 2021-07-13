#!/usr/bin/env bash

set -x

network="nqqcni"
pod="realworld-article"

podman network rm $network -f
podman pod rm $pod -f
