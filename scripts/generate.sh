#!/bin/bash
set -eo pipefail


echo "Generating GO Client"
# Cleanup old files
rm -rf ./go/gcom
docker run --rm \
  -v "${PWD}":/local openapitools/openapi-generator-cli generate \
  -i /local/openapi.json \
  -g go \
  -o /local/go/gcom \
  --git-user-id grafana \
  --git-repo-id grafana-com-public-clients/go \
  -p packageName=gcom \
  -p isGoSubmodule=true

