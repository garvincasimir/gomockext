.EXPORT_ALL_VARIABLES:
CGO_ENABLED=0

.PHONY: ci-build
ci-build: tools build test lint


.PHONY: tools
tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0
	go install github.com/golang/mock/mockgen@v1.6.0


.PHONY: clean
clean:
	go clean -cache -testcache -modcache

.PHONY: build
build:
	go build -v ./...

.PHONY: fmt
fmt:
	@echo "==> running Go format <=="
	gofmt -s -l -w .

.PHONY: lint
lint:
	@echo "==> linting Go code <=="
	@$(shell go env GOPATH)/bin/golangci-lint run ./...
	@echo "==> running go vet <=="
	go vet -v ./...

.PHONY: test
test:
	@echo "==> running Go tests <=="
	CGO_ENABLED=1 go test -p 1 -v -race -short -covermode=atomic -coverprofile=coverage.out ./... 

.PHONY: mock
mock:
	mockgen -source example/foobar.go -destination example/foobar_mock.go -package example

