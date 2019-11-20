package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int,1)

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	fmt.Println("done")
	// 	ch <- 1
	// }()

	// <-ch

	// m := map[string]string{"www": "qwe"}
	// a := "www"
	// fmt.Println(m[a])
	buffer()
}

func routine() {
	go Spinner(time.Millisecond * 100)
	n := Fib(45)

	fmt.Printf("\rFib(45)=%d\n", n)
}
