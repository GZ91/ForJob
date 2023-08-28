#!/bin/sh

docker_runner() {
    docker run \
        -e DATABASE_DSN=$DATABASE_DSN \
        -e SERVER_ADDRESS=$SERVER_ADDRESS \
        -e LOG_LEVEL=$LOG_LEVEL \
        -e Secret_Key=$Secret_Key \
        -e BASE_URL=$BASE_URL \
        -e LvlLogs=$LvlLogs \
        --detach -p 8081:8080 --name stamp-reductor-tst gitlab.sminex.com:5050/web-bim/sminex-stamp-reductor:tst
}

if [ "$(docker ps -aq -f name=stamp-reductor-tst)" ]; then
        docker rm -f stamp-reductor-tst
        docker_runner
else
        docker_runner
fi
