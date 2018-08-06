#!/bin/sh

DIR_OUTPUT=./build/
FILENAME=colorme

rm -rf ${DIR_OUTPUT}public ${DIR_OUTPUT}${FILENAME} ${DIR_OUTPUT}env.example
GOOS=linux GOARCH=amd64 go build -o ${DIR_OUTPUT}${FILENAME} -v main.go
cp env.example ${DIR_OUTPUT}
cp -r ./public/. ${DIR_OUTPUT}public
cd ${DIR_OUTPUT}
git add .
git commit -a -m "ADD build at `date`"
git push origin master
cd ..
