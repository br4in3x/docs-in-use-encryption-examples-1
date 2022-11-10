#!/bin/sh
# run this first to download mongodb enterprise dockerfile and entrypoint

export MONGODB_VERSION=6.0
curl \
  -O \
  --remote-name-all https://raw.githubusercontent.com/docker-library/mongo/master/$MONGODB_VERSION/{Dockerfile,docker-entrypoint.sh}

chmod 755 ./docker-entrypoint.sh
