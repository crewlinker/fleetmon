TEST_ARGS = -failfast

update:
	go get -u
	go mod tidy
	go mod verify

lint:
	golangci-lint run

test: lint
	go test $(TEST_ARGS) ./...

test-real: lint
	@echo "#######################################"
	@echo "### Remember to export FLEETMON_KEY ###"
	@echo "#######################################"
	go test -tags real $(TEST_ARGS) ./...

test-cover:
	go test $(TEST_ARGS) -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out
