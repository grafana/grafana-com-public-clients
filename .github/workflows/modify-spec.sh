#! /usr/bin/env bash
set -euo pipefail

SCHEMA=`cat` # Read stdin
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}
modify '.components.schemas.FormattedApiApiKey.properties.id = { "anyOf": [ { "type": "string" }, { "type": "number" } ] }'
modify '.components.schemas.FormattedApiApiKeyListResponse.properties.items.items.properties.id = { "anyOf": [ { "type": "string" }, { "type": "number" } ] }'
modify '.components.schemas.FormattedOrgMembership.properties.allowGCloudTrial = { "anyOf": [ { "type": "boolean" }, { "type": "number" } ] }'
modify '.components.schemas.FormattedApiOrgPublic.properties.allowGCloudTrial = { "anyOf": [ { "type": "boolean" }, { "type": "number" } ] }'
modify '.paths["/v1/accesspolicies"].get.responses["200"].content["application/json"].schema = {
  "type": "object",
  "properties": {
    "items": {
      "type": "array",
      "items": {
        "$ref": "#/components/schemas/AuthAccessPolicy"
      }
    }
  }
}'

echo "${SCHEMA}"
