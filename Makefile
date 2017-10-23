NAME:=$(shell basename `git rev-parse --show-toplevel`)
RELEASE:=$(shell git rev-parse --verify --short HEAD)

VERSION = 0.1.0

all: build

clean:
	rm -rf pkg bin

up: build down
	docker run -d --name bruteforceserver -p 8080:8080 crackthecode

build: serverbuild
	go build -ldflags "-X main.buildVersion=$(VERSION)-$(RELEASE)" -o ctc
	cp ctc /usr/local/bin

serverbuild:
	cd server && ./dockerbuild.sh

down:
	(docker stop bruteforceserver || true) && (docker rm bruteforceserver || true)
