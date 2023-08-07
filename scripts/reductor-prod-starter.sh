#!/bin/sh

docker_runner() {
    docker run \
        -e DATABASE_DSN=$DATABASE_DSN \
        -e SERVER_ADDRESS=$SERVER_ADDRESS \
        -e LOG_LEVEL=$LOG_LEVEL \
        -e Secret_Key=$Secret_Key \
        -e BASE_URL=$BASE_URL \
        -e LvlLogs=$LvlLogs \
        --detach -p 8080:8080 --name stamp-reductor gitlab.sminex.com:5050/web-bim/sminex-stamp-reductor
}

if [ "$(docker ps -aq -f name=stamp-reductor)" ]; then
        docker rm -f stamp-reductor
        docker_runner
else
    docker_runner
fi
