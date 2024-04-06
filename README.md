# OpenAPI to Mermaid Diagram Converter

## Overview

This Go-based tool converts OpenAPI specifications into Mermaid diagrams, aiding in the visualization of API designs.

## Tasks

### build

Build a local version.

```sh
go run ./get-version > .version
cd cmd
go build
```

### build-snapshot

Use goreleaser to build the command line binary using goreleaser.

```sh
goreleaser build --snapshot --clean
```

### test

Run Go tests.

```sh
go run ./get-version > .version
go test ./...
```

### lint

```sh
golangci-lint run --verbose
```

### push-release-tag

Push a semantic version number to Github to trigger the release process.

```sh
./push-tag.sh
```
