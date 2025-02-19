#!/usr/bin/env bash
set -euo pipefail

my_dir="$( cd "$(dirname "${0}")" && pwd )"
release_dir="$( cd "${my_dir}" && cd ../.. && pwd )"
workspace_dir="$( cd "${release_dir}" && cd .. && pwd )"

go_bin=$(go env GOPATH)
export PATH=${go_bin}/bin:${PATH}
export CGO_ENABLED=0

# inputs
semver_dir="${workspace_dir}/version-semver"

# outputs
output_dir=${workspace_dir}/out

semver="$(cat "${semver_dir}/number")"
timestamp=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

binname="azure-storage-cli-${semver}-${GOOS}-amd64"
if [ "${GOOS}" = "windows" ]; then
	binname="${binname}.exe"
fi

pushd "${release_dir}" > /dev/null
  git_rev=$(git rev-parse --short HEAD)
  version="${semver}-${git_rev}-${timestamp}"

  echo -e "\n building artifact with $(go version)..."
  go build -ldflags "-X main.version=${version}" \
    -o "out/${binname}"                          \
    github.com/cloudfoundry/bosh-azure-storage-cli

  echo -e "\n sha1 of artifact..."
  sha1sum "out/${binname}"

  mv "out/${binname}" "${output_dir}/"
popd > /dev/null
