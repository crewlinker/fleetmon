TEST_ARGS = -failfast

update:
	go get -u
	go mod tidy
	go mod verify

fmt:
	go fmt ./...

lint:
	golangci-lint run

test: fmt lint
	go test $(TEST_ARGS) ./...

test-cover: fmt
	go test $(TEST_ARGS) -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -f coverage.out
