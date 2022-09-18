#!/bin/bash

export CGO_ENABLED=1
export GO111MODULE=on
export PKG="{{ .ProjectName }}/..."

if [ -n "$CI" ]; then
    echo -n "Checking go vet in ${PWD}: "
    ERRS=$(go vet ${PKG} 2>&1 || true)
    if [ -n "${ERRS}" ]; then
        echo "FAIL"
        echo "${ERRS}"
        echo
        exit 1
    fi
    echo "PASS"
    echo
fi

echo -n "Running tests: "
export CONFIG_FILE="${PWD}/config.test.yaml"
if [ $(uname -s) = "Darwin" ]; then
    export CONFIG_FILE="${PWD}/config.macos.yaml"
fi
echo "CONFIG_FILE=${CONFIG_FILE}"

export SQLITE_SCHEMA="${PWD}/migrations/sqlite/schema.sql"
export POSTGRES_SCHEMA="${PWD}/migrations/postgres/schema.sql"

if [ -n "$CI" ]; then
    echo "Running go test in CI mode..."
	gocov test ${PKG} | gocov-xml > coverage.xml
else
    go test -v -failfast -p 1 ${PKG} | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
fi
echo "PASS"
echo
