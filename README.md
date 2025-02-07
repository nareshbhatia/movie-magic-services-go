# Movie Magic Services in Go

The Movie Magic Server exposes MovieService as a gRPC endpoint at localhost:30000.
It uses the [Buf CLI](https://buf.build/docs) to generate Go code for this service.

## Getting started

1. Make sure you have Go and Buf installed.
2. In shell 1, run `go run ./server/main.go` to start the server.
3. In shell 2, run `go run ./client/main.go` to start the client.

The client should now print a list of movies.

Note: Server runs on port 30000.

## Verify service using gRPC protocol

```shell
buf curl --schema . --protocol grpc --http2-prior-knowledge http://localhost:30000/movie.v1.MovieService/ListMovies
```

This prints a list of movies in JSON format.

## Verify service using gRPC-web protocol

1. Install the
   [Envoy proxy](https://www.envoyproxy.io/docs/envoy/latest/start/install). We will
   use it to make the Movie Magic gRPC service available as a gRPC-web service.
2. Start Envoy: `envoy -c envoy.yaml`. The `envoy.yaml` file configures Envoy to make the
   gRPC-web service available at http://localhost:8080/api
3. Run `buf curl` to verify the service:

```shell
buf curl --schema . --protocol grpcweb --http2-prior-knowledge http://localhost:8080/api/movie.v1.MovieService/ListMovies
```

This prints a list of movies in JSON format.

**Note: Currently this does not work **

It returns this error:

```json
{
   "code": "unimplemented",
   "message": "unknown service /movie.v1.MovieService"
}
```


## Running proto-gen

```shell
make proto-lint
make proto-gen
```
