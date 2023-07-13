#!/usr/bin/env bash
set -euo pipefail

export CGO_ENABLED=0

# log current go version
go version

# run unit tests
go run github.com/onsi/ginkgo/v2/ginkgo --skip-file=integration ./...
