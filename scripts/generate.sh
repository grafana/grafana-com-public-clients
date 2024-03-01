#!/bin/bash
set -eo pipefail

# Download the openapi-generator-cli (if not already present)
if ! [ -f ./openapi-generator-cli.jar ]; then
  wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.2.0/openapi-generator-cli-7.2.0.jar -O ./openapi-generator-cli.jar
fi

# Cleanup old files
rm -rf ./go/gcom

# Workarounds for temporary schema issues
SCHEMA="$(cat openapi.json)"
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}
modify '.components.schemas.FormattedApiOrgPublic.properties.allowGCloudTrial = { "anyOf": [ { "type": "boolean" }, { "type": "number" } ] }'
# Remove all required fields on models, these objects aren't used in requests and validating the response doesn't work well
modify '.components = (.components | walk(if type == "object" then del(.required) else . end))'

TEMP_OPENAPI_FILE="openapi-temp.json"
echo "${SCHEMA}" > $TEMP_OPENAPI_FILE
trap "rm -f $TEMP_OPENAPI_FILE" EXIT


java -jar openapi-generator-cli.jar generate \
  -i ${TEMP_OPENAPI_FILE} \
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

