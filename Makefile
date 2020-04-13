GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BIN=go-vasttrafik
# App specific Variables

VT_KEY=
VT_SECRET=
BASE_URL=https://api.vasttrafik.se/bin/rest.exe/v2

all: build

.PHONY: build
build:
	$(MAKE) build-linux

.PHONY: run
run:
	DEBUG=true VT_KEY=${VT_KEY} VT_SECRET=${VT_SECRET} ${GOCMD} run .

# Cross compilation
.PHONY: build-linux
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BIN) -v