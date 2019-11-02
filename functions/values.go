package main

import (
	"fmt"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func fnValue() {
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(10))

	// f = product(3, 4) // error: cannot use product(3, 4) (value of type int) as func(n int) int value in assignment
}



func squares() func()int {
	var x int

	return func() int {
		x++
		return x * x
	}
}


func sum(numbers ...int) int {
	var s int
	for _, num := range numbers {
		s += num
	}
	return s
}
