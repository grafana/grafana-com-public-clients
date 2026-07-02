#!/usr/bin/env bash
set -euo pipefail

OPENAPI_URL="https://grafana.com/api/openapi.json"
SPEC_PATH="openapi.json"
FRESH_SPEC="/tmp/openapi-fresh.json"

curl -fsSL "${OPENAPI_URL}" -o "${FRESH_SPEC}"

normalize() {
  jq -c 'del(.info."x-commit-sha", .info."x-generated-at")' "$1"
}

if [ "$(normalize "${SPEC_PATH}")" = "$(normalize "${FRESH_SPEC}")" ]; then
  echo "No OpenAPI drift detected (ignoring info.x-commit-sha and info.x-generated-at)."
  echo "drift=false" >> "${GITHUB_OUTPUT}"
  exit 0
fi

echo "OpenAPI drift detected."
echo "drift=true" >> "${GITHUB_OUTPUT}"
