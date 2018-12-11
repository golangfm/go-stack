package main

import (
	"context"
	"log"

	"github.com/golangfm/go-stack/grpc/grpc_simple"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	client := grpc_simple.NewCalcClient(conn)
	numx := int32(2)
	numy := int32(3)
	op := "+"
	result, err := client.Calc(context.Background(), &grpc_simple.CalReq{Valx: &numx, Valy: &numy, Op: &op})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%+v", result)

}
