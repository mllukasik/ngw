version=latest

run:
	@go run . $(ARGS)

.PHONY: build
build:
	@go build

.PHONY: test
test:
	@go test ./...

.PHONY: build-release
build-release:
	@go build -v -ldflags="-X main.version=$(version)"

