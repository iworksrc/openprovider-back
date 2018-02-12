#!/usr/bin/env bash

NAME='openprovider-back'

IMAGE=$(sudo docker images -a --format="{{.Repository}}" | grep $NAME)

if [ -z $IMAGE ];then
    echo "run build.sh first, exiting..."
    exit 1
fi

#CONTAINER_ID=$(sudo docker ps -a --filter="name=$NAME" -q)
#
#echo "CONTAINER_ID="$CONTAINER_ID
#
#if [[ $CONTAINER_ID > "" ]];then
#    echo "container $NAME is present, try run..."
#    sudo docker run $NAME
#else
#    echo "conteiner with name $NAME not found, create&run..."
#    sudo docker run -h localhost -p 8080:8080 --name $NAME $NAME:latest
#fi

sudo docker run --rm -h localhost -p 8080:8080 --name $NAME $NAME:latest
