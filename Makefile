.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/app -v $(pwd)/.cache/golangci-lint:/root/.cache -w /app golangci/golangci-lint:v1.53.3 golangci-lint run -v ./...

.PHONY: goreleaser-snapshot
goreleaser-snapshot:
	HEAD_SHA=$(shell git rev-parse --short HEAD) goreleaser release --snapshot --clean --skip-publish --skip-sign
