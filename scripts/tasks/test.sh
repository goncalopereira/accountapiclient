#!/bin/bash
set -e
go clean -testcache
go test -v ./internal/... ./pkg/...