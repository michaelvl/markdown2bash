.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...

lint:
	docker run --rm -v $(shell pwd):/app -v $(pwd)/.cache/golangci-lint:/root/.cache -w /app golangci/golangci-lint:v1.53.3 golangci-lint run -v ./...
