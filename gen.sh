#!/bin/bash

# usage: bash gen.sh <protobufVersion>, where protobufVersion is the tag of the protobufs repo
# example: bash gen.sh v0.2.0

set -e

PROJ_DIR=$(pwd)

# dir where proto files are located
# must be cloned from https://github.com/nokia/srlinux-ndk-protobufs
PROTO_DIR=~/nokia/srlinux-ndk-protobufs

# tag matching the protobuf tag in the https://github.com/nokia/srlinux-ndk-protobufs
# from which the bindings are about to be generated
PROTO_VER=$1

if [ -z "$1" ]; then
    echo "protobufs version from which to generate bindings is not set. usage: bash gen.sh <protobufVersion>."
    exit 1
fi

# remove previously generated bindings
sudo rm -rf ${PROJ_DIR}/ndk

# checkout protos to the desired version
cd ${PROTO_DIR} && git checkout ${PROTO_VER}

PROTOC_IMAGE=ghcr.io/srl-labs/protoc:24.4__1.31.0

docker run -v ${PROTO_DIR}:/in -u $(id -u):$(id -g) -v ${PROJ_DIR}:/out ${PROTOC_IMAGE} \
  ash -c "protoc --go_out=paths=source_relative:/out --go_opt=paths=source_relative --go-grpc_out=paths=source_relative:/out --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false ndk/*.proto"

# once the bindings are generated, we can push it to the repo
# git push
# and create a release
# it is safer to create an rc release first, and then promote it to a stable release
# once the apps are tested to work with it
# gh release create v0.2.0-rc1 --generate-notes

# when publishing the release from a working branch (such as v0.4.0-rc1)
# create a release that points to the branch
# gh release create v0.4.0-rc1 --target v0.4.0 --generate-notes