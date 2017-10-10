#!/bin/bash +x

GOOS=linux GOARCH=amd64 go build -a -o server

docker build -t crackthecode .

