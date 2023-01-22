#!/usr/bin/env bash

echo "==> Start building docker container name:"
if [ -n "$1" ]
then
    echo $1.
else
    echo "All project's contaiers"
fi

if [ ! -n "$1" ]
then
    echo "Drop old containers"
    docker-compose rm -vsf
    docker-compose down -v --remove-orphans
fi

docker-compose build $(if [ -n "$1" ]; then echo $1; fi;)

exit 0
