# Movie Magic Services in Go

The Movie Magic Server exposes MovieService as a gRPC endpoint at localhost:30000.
It uses the [Buf CLI](https://buf.build/docs) to generate Go code for this service.

## Getting started

1. Make sure you have Go and Buf installed.
2. In shell 1, run `go run ./server/main.go` to start the server.
3. In shell 2, run `go run ./client/main.go` to start the client.

The client should now print a list of movies.

Note: Server runs on port 30000.

You can also verify that the server is working correctly using this command:

```shell
buf curl --schema . --protocol grpc --http2-prior-knowledge http://localhost:30000/movie.v1.MovieService/ListMovies
```

It should print a list of movies in JSON format.

## Running proto-gen

```shell
make proto-lint
make proto-gen
```
