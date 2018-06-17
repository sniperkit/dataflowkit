#!/bin/sh
set -x
set -e

env GOOS=linux go build -v -o docker/parse.d
docker build --force-rm -t sniperkit/dfk-parse:3.7-alpine -f dockerfile-alpine3.7 --no-cache .