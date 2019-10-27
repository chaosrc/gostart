package main

import (
	"fmt"
	
)

type a struct {
	b int
	c string
}

func main() {
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)

	fmt.Println("gcd-----")
	// forLoop()

	rs := gcd(100, 50)
	fmt.Println(rs)

	fmt.Println("fib-----")
	fmt.Println(fib(10))

	metal := []string{"gold", "silver", "bornze"}
	fmt.Println(metal)
	
	
}

// func newFn() *int {
// 	return new(int)
// }
func newFn() *a {
	b := a{10, "hello"}
	return &b
}

func foo(new int, old int) int {
	return new - old
}

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i * i)
	}
}

func assig() {
	// x = 1
	// *p = true
	// person.name = "jack"
	// count[x] = count[x] * scale
}

func assig2() {
	b := 1
	b += 1
	b++
	b -= 1
	b--
}
