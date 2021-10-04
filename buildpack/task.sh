#!/bin/bash

for path in "$PWD/cache" "$PWD/layers" "$PWD/source"; do
    echo "> Setting permissions on '$path'..."
    chown -R "1000:1000" "$path"
done

CACHE_DIR=$PWD/cache
CACHE_IMAGE=${APP_IMAGE}-cache

export CNB_REGISTRY_AUTH="{\"index.docker.io\": \"Basic $(echo -n "${DOCKER_USERNAME}:${DOCKER_PASSWORD}" | base64)\"}"
/cnb/lifecycle/creator -app=source \
    -cache-image=${CACHE_IMAGE} \
    -cache-dir=${CACHE_DIR} \
    -uid=1000 \
    -gid=1000 \
    -layers=layers \
    -platform=platform \
    -report=layers/report.toml \
    -process-type=web \
    -skip-restore=${SKIP_RESTORE} \
    -run-image=${RUN_IMAGE} \
    ${APP_IMAGE}