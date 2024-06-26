package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/nareshbhatia/movie-magic-services-go/gen/go/movie/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:30000", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMovieServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ListMovies(ctx, &pb.ListMoviesRequest{})
	if err != nil {
		log.Fatalf("could not list movies: %v", err)
	}

	movies := r.GetMovies()
	for _, movie := range movies {
		fmt.Println(movie.Name)
	}
}
