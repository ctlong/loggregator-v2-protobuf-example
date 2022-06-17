# loggregator-v2-protobuf-example

An example for how old loggregator v2 generated protobuf files might compare to new loggregator v2 generated protobuf files.

## Setup

1. [Install go](https://go.dev/doc/install)

## Run

### Old server

1. Run `make run-old-server`.
1. Open a new terminal window.
1. Run a client against it, either:
  1. Run the old client with `make run-old-client`,
  1. Or, run the new client with `make run-new-client`.

### New server

1. Run `make run-new-server`.
1. Open a new terminal window.
1. Run a client against it, either:
  1. Run the old client with `make run-old-client`,
  1. Or, run the new client with `make run-new-client`.

## Test

```
go test -bench=. ./...
```
