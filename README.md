# Take Home Exercise

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

Build & Unit Tests & Integration Test ![Build & Test](https://github.com/goncalopereira/accountapiclient/workflows/Build%20&%20Test/badge.svg)

Docker-Compose Build & Unit Tests & Int Tests & E2E Tests ![Docker-Compose](https://github.com/goncalopereira/accountapiclient/workflows/Docker-Compose/badge.svg)

GolangCI ![golangci-lint](https://github.com/goncalopereira/accountapiclient/workflows/golangci-lint/badge.svg)

## Introduction
Hi! I'm Goncalo Pereira, I'm attempting this exercise, I *DO NOT have Golang experience* besides some attempts at home.

## Docker
As requested you can use `docker-compose up` to build and run all tests.

Optionally, `make docker`

## Running locally

If you prefer to build and run tests locally use `make`

## Examples
The [examples folder](examples) will show examples of uses of the client.

`make example` will run all examples against the fake-api on docker-compose (localhost:8080)

## Stages
In this section I explain some ways I created this exercise.

* [Discovery](docs/Discovery.md) 
* [Bootstrap](docs/Bootstrap.md)
* [Development](docs/Development.md)

