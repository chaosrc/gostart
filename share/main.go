package main

import (
	"example.com/hello/share/go-val"
	"fmt"
	"image"
	"time"

)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:

		}
	}
}

func init() {
	go teller()
}

func main() {
	
	val.Memo()
	val.Memo()
	time.Sleep(time.Second)
}

func bankTest() {
	done := make(chan int)

	go func() {
		for i := 0; i < 10000; i++ {
			Deposit(1)
		}
		done <- 1
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			Deposit(1)
		}
		done <- 1
	}()
	for i := 0; i < 2; i++ {
		<-done
	}
	fmt.Println(Balance())
}

var icons = make(map[string]image.Image)
