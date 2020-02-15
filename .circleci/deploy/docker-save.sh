#!/bin/bash

set -e

mkdir -p /tmp/workspace/docker-cache

docker save -o /tmp/workspace/docker-cache/statisticodata_grpc.tar statisticodata_grpc:latest
docker save -o /tmp/workspace/docker-cache/statisticodata_rest.tar statisticodata_rest:latest
