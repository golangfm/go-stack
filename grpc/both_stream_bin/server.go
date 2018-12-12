package main

import (
	"log"
	"net"

	"github.com/golangfm/go-stack/grpc/both_stream"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	both_stream.RegisterChatServer(grpcServer, &both_stream.ChatService{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
