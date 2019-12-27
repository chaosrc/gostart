package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"example.com/hello/rpc/grpc/hello"
)

type HelloServiceImp struct {
	auth *Authentication
}

func (service *HelloServiceImp) Hello(ctx context.Context, request *hello.String) (*hello.String, error) {
	if err := service.auth.Auth(ctx); err != nil {
		return nil, err
	}

	value := "Hello " + request.GetValue()

	return &hello.String{Value: value}, nil
}
func (service *HelloServiceImp) Channel(stream hello.HelloService_ChannelServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println(req.GetValue())
		err = stream.Send(&hello.String{Value: "Hello " + req.GetValue()})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	serv := grpc.NewServer()
	helloService := HelloServiceImp{
		auth: &Authentication{User: "foo", Password: "123456"},
	}
	hello.RegisterHelloServiceServer(serv, &helloService)

	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
