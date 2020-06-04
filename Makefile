.PHONY: build test docker e2e

fmt:
	gofmt -w .
build: fmt ## format and build all go files
	go build -v ./...
test: fmt ## run unit and integration test with no external dependencies
	scripts/tasks/test.sh
e2e: ## run tests onto local account api
	scripts/tasks/e2e.sh
docker:
	@docker-compose build --no-cache test
	@docker-compose up --force-recreate --abort-on-container-exit --exit-code-from test
cover:
	go test ./internal/* -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
