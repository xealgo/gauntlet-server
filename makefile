BUILD := $(shell date -u +%m%d%H%M)

default: build-dev

.PHONY: default build-dev

install:
	@cd ./src/gauntlet && go install -tags prod -ldflags "-X main.BUILD=SD$(BUILD)" .

build:
	@cd ./src/gauntlet && go build -o ../../bin/gauntlet -tags prod -ldflags "-X main.BUILD=SD$(BUILD)" .

build-dev:
	@cd ./src/gauntlet && go build -o ../../bin/gauntlet -tags dev -ldflags "-X main.BUILD=SD$(BUILD)" .

run:
	@make build-dev
	@ ./bin/gauntlet

test:
	@cd ./src/gauntlet && go test -v ./...
