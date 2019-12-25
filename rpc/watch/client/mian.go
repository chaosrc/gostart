package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	doClientWork(client)
}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVService.Watch", 3, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(keyChanged)
	}()

	time.Sleep(1 * time.Second)
	err := client.Call("KVService.Set", [2]string{"name1", "jack5"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	// time.Sleep(1 * time.Second)
	// err = client.Call("KVService.Set", [2]string{"name", "jack5"}, new(struct{}))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// time.Sleep(1 * time.Second)
	// err = client.Call("KVService.Set", [2]string{"name", "jack6"}, new(struct{}))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	time.Sleep(3 * time.Second)

}
