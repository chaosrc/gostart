package main

import (
	"context"
	"fmt"
	"log"

	"example.com/hello/rpc/grpc/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &hello.String{Value: "world"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply.GetValue())

}
