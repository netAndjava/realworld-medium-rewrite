#!/usr/bin/env bash

# 1. build
go build -o article ~/go/src/iohttps.com/live/realworld-medium-rewrite/cmd/article/main.go
# 2. build image
sh ~/go/src/iohttps.com/live/realworld-medium-rewrite/articles/labs/build/build.sh
# 3. clean
sh ~/go/src/iohttps.com/live/realworld-medium-rewrite/articles/labs/clean_article.sh
# 4. run container and start service
sh ~/go/src/iohttps.com/live/realworld-medium-rewrite/articles/labs/article.sh

rm article
