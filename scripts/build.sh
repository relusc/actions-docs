#!/bin/bash
set -eu

# Unset flags
export GOFLAGS=""

GOARCH=amd64
OSLIST="darwin windows linux"
APPNAME=actions-docs
SHA=${COMMIT:-local}

version=$(git tag --points-at "$SHA" || echo "nil")
if [ "$version" == "nil" ] || [ "$version" == "" ]; then
  # No tags found for current commit, commit ID or dummy value 'local' will be used instead
  version="$SHA"
fi

# Create directory for built binaries
mkdir -p build

for GOOS in $OSLIST; do
  ext=""
  if [ "$GOOS" = "windows" ]; then
    ext=".exe"
  fi

  go_ldflags="-X main.version=$version -X main.osarch=$GOOS/$GOARCH -X 'main.buildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'"

  output_file="build/${APPNAME}_${version}_${GOOS}_${GOARCH}"
  GOOS="${GOOS}" GOARCH="${GOARCH}" go build -o "${output_file}${ext}" -ldflags "${go_ldflags}" ./cmd
done
