package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (servicev *HelloService) Hello(request string, response *string) error {
	*response = "Hello: " + request

	return nil
}

func main() {
	err := rpc.Register(new(HelloService))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, _ := net.Dial("tcp", "localhost:8002")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
	}

}
