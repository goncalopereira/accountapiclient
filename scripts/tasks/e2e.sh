#!/bin/bash
set -e
./build/ci/wait-for-it.sh accountapi:8080 -- echo 'ready for e2e!'
go test ./test/e2e