<p align=center><a href="https://learn.srlinux.dev"><img src=https://gitlab.com/rdodin/pics/-/wikis/uploads/e06b64d8bda6ef482c486120628e6706/srl-ndk-go.svg?sanitize=true/></a></p>

[![Go](https://img.shields.io/badge/Go-reference-blue?style=flat-square&color=00c9ff&labelColor=bec8d2)](https://pkg.go.dev/github.com/nokia/srlinux-ndk-go/ndk)

---

The Nokia SR Linux [NetOps Development Kit (NDK)](https://learn.srlinux.dev/ndk/) allows operators to program high-performance, integrated agents that run alongside the Nokia Service Router Linux (SR Linux).

This repository contains generated Go code for [SR Linux NDK Protocol buffers](https://github.com/nokia/srlinux-ndk-protobufs).

## Module import paths

The Go module version is synchronized with the SR Linux [NDK protobuf releases](https://github.com/nokia/srlinux-ndk-protobufs).

Users can fetch the Go bindings for NDK with `go get` command as follows:

```bash
# get latest ndk package
go get github.com/nokia/srlinux-ndk-go

# get a specific version of the ndk package
go get github.com/nokia/srlinux-ndk-go@v0.1.0
```

To use the `ndk` package, use the following import statement:

```go
import "github.com/nokia/srlinux-ndk-go/ndk"
```

## Code generation

This code has been generated from [SR Linux NDK Protocol buffers](https://github.com/nokia/srlinux-ndk-protobufs) using [`protoc` compiler](https://github.com/srl-labs/protoc-container) with the gRPC plugins for Go and Python.

`gen.sh` script calls the `protoc` container using the ndk protobufs version as its single argument.

```bash
bash gen.sh v0.1.1
```

The Go package directory named `ndk` will be populated with the generated code bindings.
