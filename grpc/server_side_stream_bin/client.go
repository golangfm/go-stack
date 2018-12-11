package main

import (
	"context"
	"io"
	"log"

	"github.com/golangfm/go-stack/grpc/server_side_stream"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	client := server_side_stream.NewCalcClient(conn)
	numx := int32(2)
	numy := int32(3)

	stream, err := client.Calcs(context.Background(), &server_side_stream.CalReq{Valx: &numx, Valy: &numy})
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v", data)
	}
}
