VERSION=$(shell make next-version)

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
	@ echo "prepare release $(VERSION)"; \
		buildDate=$$(date -R); \
		temp_buildDate="\$${buildDate}"; \
		prev_version=$$(git describe --tags --abbrev=0 | sed 's/.*\.//'); \
		next_version=$$(expr $$prev_version + 1); \
		version="v1.0.$(VERSION)"; \
		temp_version="\$${version}"; \
		sed -i "0,/$${temp_version}/{s/$${temp_version}/$${version}/}" build/build.go; \
		sed -i "0,/$${temp_buildDate}/{s/$${temp_buildDate}/$${buildDate}/}" build/build.go; \
		git add build/build.go; \
		git commit -m "prepare release $(version)"; \
		git push

.PHONY: after-release
after-release:
	@mv old_build.go build/build.go
	@git add build/build.go
	@git commit -m "prepare for next development iteration"
	@git push


.PHONY: next-version
next-version:
		@prev_version=$$(git describe --tags --abbrev=0 | sed 's/.*\.//'); \
			next_version=$$(expr $$prev_version + 1); \
			echo $$next_version

