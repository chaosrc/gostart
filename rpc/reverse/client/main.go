package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal(err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWork(clientChan)

}

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan

	var response string
	err := client.Call("HelloService.Hello", "world", &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
