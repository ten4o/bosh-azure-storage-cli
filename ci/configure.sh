#!/usr/bin/env bash
set -euo pipefail

fly -t bosh set-pipeline -p "bosh-azure-storage-cli" \
    -c "$(dirname "${0}")/pipeline.yml"
