NAME:=$(shell basename `git rev-parse --show-toplevel`)
RELEASE:=$(shell git rev-parse --verify --short HEAD)

VERSION = 0.1.0

all: build

clean:
	rm -rf pkg bin

build:
	go build -ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o ctc
	cp ctc /usr/local/bin

artifacts:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o sl-linux
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build --ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o sl-mac
