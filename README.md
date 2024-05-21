# Movie Magic Services in Go

## Getting started

1. Make sure you have Go and Buf installed.
2. In shell 1, run `go run ./server/main.go` to start the server.
3. In shell 2, run `go run ./client/main.go` to start the client.

The client should now print a list of movies.

Note: Server runs on port 30000.

## Running proto-gen

```shell
make proto-lint
make proto-gen
```
