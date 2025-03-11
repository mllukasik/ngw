version=latest

run:
	go run . @a

.PHONY: build
build:
	go build

build-release:
	go build -v -ldflags="-X main.version=$(version)"

