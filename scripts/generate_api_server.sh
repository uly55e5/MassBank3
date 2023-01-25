#!/bin/bash


OPENAPI_GENERATOR=~/bin/openapi-generator-cli.sh
export GO_POST_PROCESS_FILE="gofmt -w"

MB3_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"/..


${OPENAPI_GENERATOR}  generate -c ${MB3_DIR}/scripts/config-openapi.yaml -o ${MB3_DIR}/cmd/mb3server/