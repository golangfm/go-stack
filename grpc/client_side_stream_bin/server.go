package main

import (
	"log"
	"net"

	"github.com/golangfm/go-stack/grpc/client_side_stream"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	client_side_stream.RegisterCalcServer(grpcServer, &client_side_stream.SumService{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
