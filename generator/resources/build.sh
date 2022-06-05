#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

BINARY=${1}
OS="linux"
ARCH="amd64"
VERSION=$(git describe --tags --always)
NOW=$(date +"%Y%m%d_%H%M")

echo "Building ${NOW} ${VERSION}"

GOARCH="${ARCH}"
GOOS="${OS}"
GO111MODULE=on
APP_NAME=$(go list -m)

CGO_ENABLED=0 go build \
    -installsuffix "static" \
    -ldflags "-X ${APP_NAME}/pkg.VERSION=${VERSION} -X ${APP_NAME}/pkg.DATETIME=${NOW} -s -w" \
    -o "${BINARY}-${VERSION}"

ls -shFx ${BINARY}-${VERSION}
upx=`which upx`
if [ "$?" -eq 0 ]; then
    upx ${BINARY}-${VERSION}
fi
