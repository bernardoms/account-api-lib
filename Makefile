.PHONY: dependency unit-test integration-test docker-up docker-down clear

dependency:
	@go mod download

integration-test: docker-up dependency
	@cd test/integrated; go test ./... -v

unit-test: dependency
	@cd test/unit; go test ./... -v -short

all-tests-inside-container:
	@export TEST_URL=http://accountapi:8080/v1/organisation/accounts
	@go test ./... -v -short

docker-up:
	@docker-compose up -d
	@sleep 10

docker-down:
	@docker-compose stop

clear: docker-down