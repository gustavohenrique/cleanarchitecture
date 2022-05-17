#!/usr/bin/env bash

OUTPUT_DIR="${PWD}"

base_dir=$(dirname $PWD)
spa_dir="${base_dir}/spa-quasar/dist"
if [ -d "${spa_dir}" ]; then
    echo "==> Copying spa-quasar to src/app/httpserver/ui/spa..."
    rm -rf src/app/httpserver/ui/spa 2>/dev/null
    cp -r ${spa_dir} src/app/httpserver/ui
else
    echo "No UI found in ${spa_dir}"
    exit 1
fi

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64 arm"}
XC_OS=${XC_OS:-linux darwin windows freebsd}
XC_EXCLUDE_OSARCH="!darwin/arm !darwin/386 !freebsd/arm"

# Delete the old dir
echo "==> Removing old directory..."
rm -rf ${OUTPUT_DIR}/{bin,releases}
mkdir -p ${OUTPUT_DIR}/bin/

# If its dev mode, only build for ourself
if [[ -n "${DEV_MODE}" ]]; then
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)
fi

if ! which gox > /dev/null; then
    echo "==> Installing gox..."
    go get -u github.com/mitchellh/gox
fi

# Instruct gox to build statically linked binaries
export CGO_ENABLED=0

# Set module download mode to readonly to not implicitly update go.mod
export GOFLAGS="-mod=readonly"

# In release mode we don't want debug information in the binary
if [[ -n "${RELEASE_MODE}" ]]; then
    LD_FLAGS="-s -w"
fi

# Ensure all remote modules are downloaded and cached before build so that
# the concurrent builds launched by gox won't race to redundantly download them.
go mod download

# Build!
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="${XC_EXCLUDE_OSARCH}" \
    -ldflags "${LD_FLAGS}" \
    -output "${OUTPUT_DIR}/releases/{{.OS}}_{{.Arch}}/myproject" \
    ${OUTPUT_DIR}/cmd/myproject

# Copy our OS/Arch to the bin/ directory
DEV_PLATFORM="${OUTPUT_DIR}/releases/$(go env GOOS)_$(go env GOARCH)"
if [[ -d "${DEV_PLATFORM}" ]]; then
    for F in $(find ${DEV_PLATFORM} -mindepth 1 -maxdepth 1 -type f); do
        cp ${F} ${OUTPUT_DIR}/bin/
    done
fi

if [ -n "${RELEASE_MODE}" ]; then
    # Zip and copy to the dist dir
    echo "==> Packaging..."
    for PLATFORM in $(find $OUTPUT_DIR/releases -mindepth 1 -maxdepth 1 -type d); do
        OSARCH=$(basename ${PLATFORM})
        echo "--> ${OSARCH}"

        pushd $PLATFORM >/dev/null 2>&1
        zip ../${OSARCH}.zip ./*
        popd >/dev/null 2>&1
    done
fi

# Done!
echo
echo "==> Results:"
ls -hl $OUTPUT_DIR/bin/
