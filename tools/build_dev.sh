#!/usr/bin/env bash
# Docker hub https://cloud.docker.com/u/unknownholding/repository/docker/$DOCKER_REPO
img="$DOCKER_REPO:develop"

rm -r ~/$DIR/src/swarm-util/logs/*

docker build -t $img -f DockerfileDev .
docker push $img

echo "Finish build $img"
