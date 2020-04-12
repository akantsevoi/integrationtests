#!/bin/sh

RESULT=$(CGO_ENABLED=0 go test ./cmd/integrationtests -count=1 2> /dev/null)
if [ $? -ne 0 ]; then
    echo "Test errors: $RESULT" >&2
    exit 1
else 
    echo "Test success: $RESULT"
fi