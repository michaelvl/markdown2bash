VERSION=$(shell git describe --tags --always --dirty)
IMAGE_VERSION=sha-$(shell git describe --tags --always)-$(shell dpkg --print-architecture)

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

.PHONY: run-readme-example-local
run-readme-example-local:
	cat README.md | docker run --rm -i ghcr.io/local/markdown2bash:$(IMAGE_VERSION) > readme.sh
