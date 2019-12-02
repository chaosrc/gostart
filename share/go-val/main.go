package val

import (
	"fmt"
)

var a = 0

func Add() int {
	a++

	return a
}

func Get() int {
	return a
}

func main() {

}

func Memo() {
	var x, y int
	go func() {
		x = 1
		fmt.Println("y:", y)
	}()
	go func() {
		y = 1
		fmt.Println("x:", x)
	}()
}

func bankTest() {
	done := make(chan struct{})

	go func() {
		for i := 0; i < 10000; i++ {
			Deposit(1)
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			Deposit(1)
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(Balance())
}
