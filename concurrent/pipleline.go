package main

import (
	"fmt"
)

func pipeline() {
	naturals := make(chan int)
	squares := make(chan int)
	

	// Counter
	go counter(naturals)
	// Squarer
	go squarer(naturals, squares)

	// Printer
	printer(squares)
}

func counter(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
