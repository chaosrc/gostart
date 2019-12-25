package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/golang/protobuf/proto"

	"example.com/hello/rpc/message"
)

type HelloService struct {
}

func (h *HelloService) Hello(request string, response *string) error {
	*response = "Hello: " + request
	return nil
}

func main() {
	name := "John"
	id := int32(1234)
	email := "123456@qq.com"
	person := &message.Person{
		Name:  &name,
		Id:    &id,
		Email: &email,
	}

	person2 := &message.Person{}
	data, _ := proto.Marshal(person)
	proto.Unmarshal(data, person2)

	fmt.Println(person2)
}

func startServer() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	rpc.ServeConn(conn)
}
