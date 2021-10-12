<p align=center><a href="https://learn.srlinux.dev"><img src=https://gitlab.com/rdodin/pics/-/wikis/uploads/e06b64d8bda6ef482c486120628e6706/srl-ndk-go.svg?sanitize=true/></a></p>

---

The Nokia SR Linux NetOps Development Kit (NDK) allows operators to program high-performance, integrated agents that run alongside the Nokia Service Router Linux (SR Linux).

This repository contains generated Go code for [SR Linux NDK Protocol buffers](https://github.com/nokia/srlinux-ndk-protobufs).

## Module import paths
SR Linux NDK follows the SR Linux system version (20.11.1, 21.6.2, etc) therefore for transparency reasons the Go modules are created with the same version (v20.11.1, v21.6.2), and in accordance to [Go modules specification](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) the modules will follow the "Major branch" option.

For example, the Go bindings for NDK v21.6.2 will be contained in a branch `v21` of this repo with a git tag matching its version. Consequently, the `go get` path for that version of the NDK will be:

```
go get github.com/nokia/srlinux-ndk-go/v21@v21.6.2
```

To use the `ndk` package that is part of this module, use the following import statement:

```go
import "github.com/nokia/srlinux-ndk-go/v21/ndk"
```

## Code generation
This code has been generated from [SR Linux NDK Protocol buffers](https://github.com/nokia/srlinux-ndk-protobufs) using [`protoc` 
compiler](https://github.com/srl-labs/protoc-container) with gRPC-Go plugin.

Here is the code generation command that produces the bindings captured in this repo:

> Assuming [`srlinux-ndk-protobufs`](https://github.com/nokia/srlinux-ndk-protobufs) cloned to the home directory **and** checkout out to the needed release/tag.

```bash
docker run -v ~/srlinux-ndk-protobufs:/in -v $(pwd):/out ghcr.io/srl-labs/protoc \ 
  bash -c "protoc --go_out=paths=source_relative:/out --go-grpc_out=paths=source_relative:/out ndk/*.proto"
```

The Go package directory named `srlinux-ndk` will be created in the current working directory.
