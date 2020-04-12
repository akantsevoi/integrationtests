#!/bin/sh

if [[ -z "$1" || -z "$2" ]]; then
    { echo "empty parameters $1 - MYSQL_HOST $2 - after run script"; exit 1; }
fi

MYSQL_HOST=$1
SECONDS_CHECK=40

echo "Check db: $MYSQL_HOST"

while ! mysqladmin ping -h"$MYSQL_HOST" --silent; do
    sleep 1
    SECONDS_CHECK="$(($SECONDS_CHECK-1))"
    if [ "$SECONDS_CHECK" -lt 0 ]; then
        { echo "DB doesn't answer $MYSQL_HOST"; exit 1; }
    fi

    echo "Wait $SECONDS_CHECK"
done

echo "DB responses"
sh $2