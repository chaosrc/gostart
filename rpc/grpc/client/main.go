package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"example.com/hello/rpc/grpc/hello"
	"google.golang.org/grpc"
)

func main() {
	auth := Authentication{
		User:     "foo",
		Password: "1234567",
	}

	conn, err := grpc.Dial("localhost:8003", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &hello.String{Value: "world"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply.GetValue())

	// callChannel(client)
}

func callChannel(client hello.HelloServiceClient) {
	channelClient, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		var i int
		for {
			err := channelClient.Send(&hello.String{Value: "channel" + strconv.Itoa(i)})
			fmt.Println("q323234", err)

			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
			i++
		}

	}()

	for {
		res, err := channelClient.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("finish")
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println(res.GetValue())
		time.Sleep(time.Second)
	}
}
