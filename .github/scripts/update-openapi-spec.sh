#!/usr/bin/env bash
set -euo pipefail

# This script is called by Renovate's postUpgradeTasks.
# It downloads the latest OpenAPI spec and checks whether the API surface
# actually changed (ignoring volatile metadata like x-generated-at and
# x-commit-sha). If nothing meaningful changed, it restores the original
# openapi.json so Renovate sees no diff and skips PR creation.

SPEC_URL="https://grafana.com/api/openapi.json"
SPEC_FILE="openapi.json"

# Download the latest spec
curl -sf "${SPEC_URL}" -o /tmp/new-spec.json

# Get the current spec from git (before Renovate's modification)
git show "HEAD:${SPEC_FILE}" > /tmp/old-spec.json

# Strip volatile metadata fields from both and compare.
# These fields change on every upstream rebuild even when the API surface
# is identical, causing unnecessary PRs.
strip_metadata() {
    sed -E '/"x-generated-at"|"x-commit-sha"/d'
}

old_hash=$(strip_metadata < /tmp/old-spec.json | sha256sum | cut -d' ' -f1)
new_hash=$(strip_metadata < /tmp/new-spec.json | sha256sum | cut -d' ' -f1)

if [ "${old_hash}" = "${new_hash}" ]; then
    echo "No meaningful API changes detected, restoring original spec."
    git checkout -- "${SPEC_FILE}"
else
    echo "API changes detected, updating spec."
    cp /tmp/new-spec.json "${SPEC_FILE}"
fi
