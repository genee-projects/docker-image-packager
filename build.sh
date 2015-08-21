#!/bin/bash

docker run \
		--rm \
		-v "$PWD":/go/src/github.com/genee-tools/docker-images-packager \
		-w /go/src/github.com/genee-tools/docker-images-packager \
		golang:latest \
		bash -c 'cd /go/src/github.com/genee-tools/docker-images-packager && go build -o packager'

echo "done!"
