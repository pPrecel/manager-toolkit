# setup-go

A GitHub action that sets up the Go toolchain using the version declared in the repository's `go.mod` file.

## How It Works

Wraps [`actions/setup-go`](https://github.com/actions/setup-go) with project-standard defaults:

- Reads the Go version from `go.mod` at the repository root.
- Uses `go.sum` as the cache key for the module cache.

## Inputs

This action has no inputs.

## Usage

```yaml
- uses: actions/checkout@v4

- uses: kyma-project/manager-toolkit/.github/actions/setup-go@main
```

Use this action before any step that runs `go` commands or before actions that require Go to be available in `PATH` (e.g. [`test-rbac-propagation`](../test-rbac-propagation)).
