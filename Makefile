.PHONY: build
build:
	go build -ldflags "-s" ./cmd/skeleton

#.PHONY: build
#build:
#	go build -v ./cmd/skeleton

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go run ./cmd/skeleton/skeleton.go

.DEFAULT_GOAL := build