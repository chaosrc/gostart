package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"example.com/hello/rpc/pubsub/sub"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8004", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := sub.NewPubsubServiceClient(conn)
	stream, err := client.Subsribe(context.Background(), &sub.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			continue
		}
		fmt.Println(reply.GetValue())
		time.Sleep(time.Second * 1)
	}
}
