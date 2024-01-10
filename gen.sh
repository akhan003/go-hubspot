#!/bin/bash

BINARY_PATH=$(which gofmt)

if [ -z "$BINARY_PATH" ]; then
    echo "Error: gofmt binary not found." >&2
    exit 1
fi

export GO_POST_PROCESS_FILE="$BINARY_PATH -w"

# Lists
openapi-generator-cli generate -c configs/lists.yaml