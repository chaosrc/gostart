package main

import (
	"fmt"
	"os"
	"time"
)

func count() {
	abord := make(chan struct{})

	go func() {
		key := make([]byte, 1)
		os.Stdin.Read(key)

		abord <- struct{}{}
	}()

	// tick := time.Tick(time.Second * 1)
	tick := time.NewTicker(time.Second * 1)

	defer tick.Stop()
	for num := 10; num > 0; num-- {
		fmt.Println(num)
		select {
		case <-tick.C:
		case <-abord:
			fmt.Println("Abord")
			return
		}
	}
}
