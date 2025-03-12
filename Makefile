temp_version=$${version}
temp_buildDate=$${buildDate}
version=$(temp_version)
buildDate := $(shell date -R)

run:
	@go run . $(ARGS)

.PHONY: build
build:
	@go build

.PHONY: test
test:
	@go test ./...

.PHONY: prepare-release
prepare-release:
	@cp build/build.go old_build.go
	@sed -i '0,/$(temp_version)/{s/$(temp_version)/$(version)/}' build/build.go
	@sed -i '0,/$(temp_buildDate)/{s/$(temp_buildDate)/$(buildDate)/}' build/build.go
	@git add build/build.go
	@git commit -m "prepare release $(version)"
	@git push

.PHONY: after-release
after-release:
	@mv old_build.go build/build.go
	@git add build/build.go
	@git commit -m "prepare for next development iteration"
	@git push

