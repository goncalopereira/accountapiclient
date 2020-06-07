#!/bin/bash
set -e
./build/ci/wait-for-it.sh "${API_HOST}":"${API_PORT}" -- echo 'ready for e2e!'
go test -v  ./test/e2e/... -count 1
