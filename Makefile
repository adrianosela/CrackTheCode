NAME:=$(shell basename `git rev-parse --show-toplevel`)
RELEASE:=$(shell git rev-parse --verify --short HEAD)

VERSION = 0.1.0

all: run

clean:
	rm -rf pkg bin

up: build
				docker run -d --name bruteforceserver -p 8080:8080 crackthecode

build: clientbuild

clientbuild: serverbuild
				go build -ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o ctc
				cp ctc /usr/local/bin

serverbuild:
				cd server && ./dockerbuild.sh

artifacts:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o ctc-linux
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build --ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o ctc-mac

down:
				docker stop bruteforceserver
