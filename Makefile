.EXPORT_ALL_VARIABLES:
TOOLS=$(CURDIR)/_tools
CGO_ENABLED=0

.PHONY: tools
tools:
	mkdir -p $(TOOLS)
	@[ ! -f $(TOOLS)/go.mod ] && cd $(TOOLS) && go mod init tools || true
	cd $(TOOLS) && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0
	cd $(TOOLS) && go install github.com/golang/mock/mockgen@v1.6.0


.PHONY: clean
clean:
	rm -rf _tools
	go clean -cache -testcache -modcache

.PHONY: build
proxima-tenant-provisioning-workflow:
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
	go vet ./...

.PHONY: test
test:
	@echo "==> running Go tests <=="
	go test -p 1 -v -race -short -covermode=atomic -coverprofile=coverage.out ./... 

.PHONY: examplemock
examplemock:
	mockgen -source example/foobar.go -destination example/foobar_mock.go -package example

