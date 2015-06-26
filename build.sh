#!/bin/bash
mkdir build
docker run --rm -v "$PWD":/go/src/github.com/cyclopsci/omni -w /go/src/github.com/cyclopsci/omni/omni golang:1.4.2 /bin/bash -c "go get github.com/tools/godep; godep go build -v"
mv omni/omni build/
