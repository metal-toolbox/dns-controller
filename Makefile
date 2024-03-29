all: lint test
PHONY: test coverage lint golint clean vendor local-dev-databases docker-up docker-down integration-test unit-test
GOOS=linux
DB_STRING=host=localhost port=26257 user=root sslmode=disable
DB=dns_controller
DEV_DB=${DB}
TEST_DB=${DB}_test
DEV_URI=dbname=${DEV_DB} ${DB_STRING}
TEST_URI=dbname=${TEST_DB} ${DB_STRING}

test: | unit-test integration-test

integration-test: test-database
	@echo Running integration tests...
	@DNSCONTROLLER_DB_URI="${TEST_URI}" go test -cover -tags testtools,integration -p 1 ./...

unit-test: | lint
	@echo Running unit tests...
	@go test -cover -short -tags testtools ./...

coverage: | test-database
	@echo Generating coverage report...
	@DNSCONTROLLER_DB_URI="${TEST_URI}" go test ./... -race -coverprofile=coverage.out -covermode=atomic -tags testtools -p 1
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out

lint: golint

golint: | vendor
	@echo Linting Go files...
	@golangci-lint run

clean: docker-clean
	@echo Cleaning...
	@rm -rf ./dist/
	@rm -rf coverage.out
	@go clean -testcache

vendor:
	@go mod download
	@go mod tidy

docker-up:
	@docker-compose -f quickstart.yml up -d crdb

docker-down:
	@docker-compose -f quickstart.yml down

docker-clean:
	@docker-compose -f quickstart.yml down --volumes

dev-database: | vendor
	@cockroach sql --insecure -e "drop database if exists ${DEV_DB}"
	@cockroach sql --insecure -e "create database ${DEV_DB}"
	@DNSCONTROLLER_DB_URI="${DEV_URI}" go run main.go migrate up

test-database: | vendor
	@cockroach sql --insecure -e "drop database if exists ${TEST_DB}"
	@cockroach sql --insecure -e "create database ${TEST_DB}"
	@DNSCONTROLLER_DB_URI="${TEST_URI}" go run main.go migrate up
	@cockroach sql --insecure -e "use ${TEST_DB};"
