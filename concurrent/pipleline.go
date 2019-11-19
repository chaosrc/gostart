package main

import (
	"fmt"
)

func pipe() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	go func() {
		for x := range naturals {
			fmt.Println("x:",x)
		}
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
