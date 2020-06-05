#!/bin/bash
set -e
go test -v ./internal/... ./pkg/... -count 1