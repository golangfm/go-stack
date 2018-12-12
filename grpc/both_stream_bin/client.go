package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"github.com/golangfm/go-stack/grpc/both_stream"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatal("conn error:" + err.Error())
		return
	}

	defer conn.Close()
	client := both_stream.NewChatClient(conn)
	ctx := context.Background()
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Println("bidStream error:" + err.Error())
		return
	}
	go func() {
		log.Println("please input something.")
		cmdInput := bufio.NewReader(os.Stdin)
		for {
			input, _ := cmdInput.ReadString('\n')
			if err := stream.Send(&both_stream.Req{Input: &input}); err != nil {
				return
			}
		}
	}()

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			log.Println("recv done from server")
			break
		}
		if err != nil {
			log.Println("recv err data:", err)
			break
		}
		log.Printf("client recv data: %s\n", data.GetOutput())
	}
}
