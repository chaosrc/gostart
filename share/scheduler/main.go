package main

import (
	"fmt"
)

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Print(1)
		}
	}()
	for i := 0; i < 100; i++ {
		fmt.Print(0)
	}
}
