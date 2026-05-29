# Grafana.com API clients

This repository contains auto generated GCOM API clients, created using [GCOM API OpenAPI spec](https://grafana.com/api/openapi.json)
and [OpenAPI Generator](https://openapi-generator.tech/).

Renovate takes care of periodically generating PRs to update with the latest client, which causes the openapi action to regenerate the golang client. It uses https://grafana.com/api/openapi.json to download it and update it.

## Available clients

- [Golang](./go/gcom/README.md)
