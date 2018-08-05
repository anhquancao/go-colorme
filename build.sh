#!/bin/sh

DIR_OUTPUT=./build/
FILENAME=colorme

rm -rf ${DIR_OUTPUT}
mkdir ${DIR_OUTPUT}
go build -o ${DIR_OUTPUT}${FILENAME} main.go
cp -r ./public/. ${DIR_OUTPUT}public
cp env.example ${DIR_OUTPUT}