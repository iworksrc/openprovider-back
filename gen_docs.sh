#!/usr/bin/env bash

GODOCEXEC=$(which godoc)

if [ -z $GODOCEXEC ];then
    echo 'godoc is not installed, please install it first, exiting...'
    exit 1
else
    echo "godoc found in $GODOCEXEC, start geneate docs..."
fi


godoc -html ./go/ > ./docs/index.html

echo 'Done'