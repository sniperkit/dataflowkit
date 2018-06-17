#!/bin/sh
set -x
set -e

env GOOS=linux go build -v -o docker/fetch.d
docker build --force-rm -t sniperkit/dfk-fetch-d:3.7-alpine --no-cache -f ./docker/dockerfile-alpine3.7