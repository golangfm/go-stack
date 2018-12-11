package main

import (
	"log"
	"net"

	"github.com/golangfm/go-stack/grpc/grpc_simple"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	grpc_simple.RegisterCalcServer(grpcServer, &grpc_simple.SimpleCalcServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
