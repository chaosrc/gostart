package main

import (
	"context"
	"log"

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

	_, err = client.Publish(context.Background(), &sub.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &sub.String{Value: "golang: hello Go2"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &sub.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}


