# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOINSTALL := $(GOCMD) install
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
GOTEST := $(GOCMD) test

# Binary name
BINARY_NAME := dck

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME)

install:
	sudo cp $(BINARY_NAME) /usr/local/bin

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

.PHONY: all build install clean test