#!/usr/bin/env bash

echo "==> Start building docker container name:"
if [ -n "$1" ]
then
    echo $1.
else
    echo "all project's contaiers"
fi

docker-compose build $(if [ -n "$1" ]; then echo $1; fi;)

exit 0