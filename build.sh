#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o openprovider-back main.go

if [ -x ./build/openprovider-back ]; then
    echo "previous build exists, remove..."
    rm ./build/openprovider-back
fi

file openprovider-back
mv openprovider-back ./build/

docker build ./build -t openprovider-back