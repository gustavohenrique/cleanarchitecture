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
echo "CONFIG_FILE=${CONTINUUM_CONFIG_FILE}"

export DGRAPH_SCHEMA_FILE="${PWD}/migrations/dgraph/schema.dql"
export SQLITE_SCHEMA_FILE="${PWD}/migrations/sqlite/schema.sql"
export POSTGRES_SCHEMA_FILE="${PWD}/migrations/postgres/schema.sql"

if [ -n "$CI" ]; then
    echo "Running go test in CI mode..."
    go test -v -p 1 -coverprofile=coverage.txt -failfast ${PKG}
    go tool cover -html ./coverage.txt -o coverage.html
else
    go test -v -failfast -p 1 ${PKG} | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
fi
echo "PASS"
echo
