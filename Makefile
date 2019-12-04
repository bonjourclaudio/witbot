# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=witbot
BINARY_UNIX=$(BINARY_NAME)_unix

all: dep test build
dep:
	dep ensure
build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/...
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -rf ./bin
run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/...
	./bin/$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_UNIX) -v ./cmd/...