package main

import (
	"log"
	"net"

	"github.com/golangfm/go-stack/grpc/server_side_stream"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	server_side_stream.RegisterCalcServer(grpcServer, &server_side_stream.ClacsServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
