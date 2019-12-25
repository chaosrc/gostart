package main

import (
	"context"
	"log"
	"net"

	"example.com/hello/rpc/grpc/hello"

	"google.golang.org/grpc"
)

type HelloServiceImp struct{}

func (service *HelloServiceImp) Hello(ctx context.Context, request *hello.String) (*hello.String, error) {
	value := "Hello " + request.GetValue()

	return &hello.String{Value: value}, nil
}

func main() {
	serv := grpc.NewServer()

	hello.RegisterHelloServiceServer(serv, new(HelloServiceImp))

	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
