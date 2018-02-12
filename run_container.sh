#!/usr/bin/env bash

NAME='openprovider-back'

IMAGE=$(sudo docker images -a --format="{{.Repository}}" | grep $NAME)

if [ -z $IMAGE ];then
    echo "run build.sh first, exiting..."
    exit 1
fi

sudo docker run --rm -h localhost -p 8080:8080 --name $NAME $NAME:latest
