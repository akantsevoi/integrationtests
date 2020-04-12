#!/bin/bash

docker-compose -f docker-compose-tests.yml up --build --abort-on-container-exit --exit-code-from app #2> /dev/null

if [ $? -eq 0 ]
then
    docker-compose -f docker-compose-tests.yml down --volumes
    echo "Application tests success"
else
    docker-compose -f docker-compose-tests.yml down --volumes
    echo "Application tests errors" >&2
    exit 1
fi