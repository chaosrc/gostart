package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("KVService", NewKVService())

	listener, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		rpc.ServeConn(conn)
	}
}
