package main

import (
	"time"
	"fmt"
)

func Fib(n int) int {
	if n < 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func Spinner(t time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(t)
		}
	}
}
