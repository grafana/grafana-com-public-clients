#!/bin/bash
set -eo pipefail

wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.2.0/openapi-generator-cli-7.2.0.jar -O ./openapi-generator-cli.jar

# Cleanup old files
rm -rf ./go/gcom


java -jar openapi-generator-cli.jar generate \
  -i openapi.json \
  -g go \
  -o ./go/gcom \
  --git-user-id grafana \
  --git-repo-id grafana-com-public-clients/go \
  -p packageName=gcom \
  -p isGoSubmodule=true \
  -p disallowAdditionalPropertiesIfNotPresent=false

if ! command -v goimports &> /dev/null
then
    go install golang.org/x/tools/cmd/goimports@latest
fi
find ./go/gcom -name \*.go -exec goimports -w {} \;

