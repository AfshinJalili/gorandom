GO ?= go
GOLANGCI_LINT ?= golangci-lint

.PHONY: test fmt fmt-check lint vet race ci release-dry

test:
	$(GO) test ./...

fmt:
	gofmt -w ./cmd ./internal

fmt-check:
	@test -z "$$(gofmt -l ./cmd ./internal)" || (echo "gofmt needed"; exit 1)

lint:
	$(GOLANGCI_LINT) run ./...

vet:
	$(GO) vet ./...

race:
	$(GO) test -race ./...

ci: fmt-check lint vet race test

release-dry:
	goreleaser release --clean --skip=publish
