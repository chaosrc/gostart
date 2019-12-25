package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	var response string
	err = client.Call("HelloService.Hello", "world", &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
