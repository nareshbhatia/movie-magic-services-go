package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/code-shaper/movie-magic/gen/go/movie/v1"
	mockdata "github.com/code-shaper/movie-magic/internal/mockdata"
	"google.golang.org/grpc"
)

var (
	port   = flag.Int("port", 30000, "The server port")
	loader = mockdata.NewLoader()
)

type server struct {
	pb.UnimplementedMovieServiceServer
}

func (s *server) ListMovies(ctx context.Context, in *pb.ListMoviesRequest) (*pb.ListMoviesResponse, error) {
	log.Printf("Received: %v", in.GetFilters())
	pageInfo := pb.PaginationInfo{
		TotalPages:      1,
		TotalItems:      10,
		Page:            1,
		PerPage:         10,
		HasNextPage:     false,
		HasPreviousPage: false,
	}
	return &pb.ListMoviesResponse{Movies: loader.GetMovies(), PageInfo: &pageInfo}, nil
}

func main() {
	flag.Parse()
	loader.Load()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
