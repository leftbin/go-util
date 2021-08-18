v ?= dev

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test -race -v -count=1 ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: deps fmt vet test
