# Go parameters
GO=go
TARGET=redcount
DOCKER=docker
REPO=shtouff
VERSION=0.1.1

all: build

build:
	$(GO) build 

clean: 
	$(GO) clean

run: build
	./$(TARGET)

docker-build:
	$(DOCKER) build -t $(REPO)/$(TARGET):$(VERSION) -t $(REPO)/$(TARGET):latest .

docker-run:
	$(DOCKER) run $(REPO)/$(TARGET):$(VERSION)

docker-push:
	$(DOCKER) push $(REPO)/$(TARGET):latest
	$(DOCKER) push $(REPO)/$(TARGET):$(VERSION)
