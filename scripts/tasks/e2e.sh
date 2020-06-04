#!/bin/bash
set -e
./build/ci/wait-for-it.sh localhost:8080 -- echo 'ready for e2e!'
go test -v  ./test/e2e/... -count 1
