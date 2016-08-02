#!/usr/bin/env bash

set -ex
IFS=$' \n\t'
VERSION="$(git describe)"

go get github.com/mitchellh/gox

gox -arch="amd64" -os="linux darwin" -output="dist/{{.OS}}/{{.Dir}}" .

cd dist/linux
tar -Jcvf ../terraform-s3-dir-"$VERSION"-linux.tar.xz ./*
cd ../darwin
tar -Jcvf ../terraform-s3-dir-"$VERSION"-darwin.tar.xz ./*
