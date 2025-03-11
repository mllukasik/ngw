version=latest

run:
	@go run . $(ARGS)

.PHONY: build
build:
	@go build

build-release:
	@go build -v -ldflags="-X main.version=$(version)"

