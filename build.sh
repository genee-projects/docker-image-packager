#!/bin/bash

docker run \
		--rm \
		-v "$PWD":/go/src/github.com/genee-tools/docker-images-transfer \
		-w /go/src/github.com/genee-tools/docker-images-transfer \
		golang:latest \
		bash -c 'cd /go/src/github.com/genee-tools/docker-images-transfer && go build -o transfer'
