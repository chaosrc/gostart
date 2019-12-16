package main

import (
	"fmt"
)

func buffer() {
	ch := make(chan string, 3)
	go func() {
		ch <- "A"
		ch <- "B"
		ch <- "C"
	}()
	fmt.Println(<-ch)

	fmt.Println(cap(ch))
	fmt.Println(len(ch))
}

func mirrorQuery() string {
	ch := make(chan string, 3)

	go func() { ch <- request("a.mirror.com") }()
	go func() { ch <- request("b.mirror.com") }()
	go func() { ch <- request("c.mirror.com") }()

	return <-ch

}
func request(url string) string {
	//...
	return ""
}
