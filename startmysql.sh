#!/bin/bash

port=$1
echo "bind to $port"
docker run -p $port:3306  --name local-mysql -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7.22
