#!/usr/bin/env bash
# Docker hub https://cloud.docker.com/u/unknownholding/repository/docker/$DOCKER_REPO
img="$DOCKER_REPO"

export GOROOT=/usr/lib/go-1.10
export GOPATH=/home/igor/$DIR
go build -i -o ./apps/Back ./main.go #gosetup


docker build -t $img -f Dockerfile .
docker push $img
echo "Finish build image $img"
