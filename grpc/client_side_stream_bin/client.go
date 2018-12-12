package main

import (
	"context"
	"log"

	"github.com/golangfm/go-stack/grpc/client_side_stream"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatal("conn error:" + err.Error())
		return
	}

	defer conn.Close()
	client := client_side_stream.NewCalcClient(conn)
	stream, err := client.Sum(context.Background())
	if err != nil {
		log.Fatal("sum error:" + err.Error())
		return
	}
	for i := 0; i <= 10; i++ {
		val := int32(i * 2)
		if err := stream.Send(&client_side_stream.Req{Val: &val}); err != nil {
			log.Fatal("send error :" + err.Error())
			continue
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("recv error:" + err.Error())
		return
	}
	log.Printf("%+v", reply)
}
