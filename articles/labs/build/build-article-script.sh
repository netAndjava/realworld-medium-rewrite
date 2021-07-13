#!/usr/bin/env bash

set -ex

#1. create container
mycontainer=$(buildah from ubuntu:dev)

#2. create mount
mymount=$(buildah mount $mycontainer)

#4. create directory to save config file
buildah run --isolation=chroot $mycontainer -- sh -c "mkdir -p /var/etc/"

#5  copy program and config file
cp ./article $mymount/usr/local/bin
cp ~/go/src/iohttps.com/live/realworld-medium-rewrite/cmd/article/dev.toml $mymount/var/etc/
cp ~/go/src/iohttps.com/live/realworld-medium-rewrite/cmd/configs/mysql.toml $mymount/var/etc/
cp ~/go/src/iohttps.com/live/realworld-medium-rewrite/cmd/configs/consul.toml $mymount/var/etc/

#6 config author,user,enviroment information
buildah config --cmd '["article","-db","/var/etc/mysql.toml","-consul","/var/etc/consul.toml","-config","/var/etc/dev.toml"]' $mycontainer

buildah config --author "nqq@aozsky.com" $mycontainer

buildah config --user "root" $mycontainer

buildah config --env "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin" $mycontainer
buildah config --env "TERM=xterm" $mycontainer
buildah config --env "LANG=en_US.UTF-8" $mycontainer

buildah config --port 5001 $mycontainer

#7 commit container

buildah commit $mycontainer hub.iohttps.com/article:dev

#8 remove container
buildah rm $mycontainer
