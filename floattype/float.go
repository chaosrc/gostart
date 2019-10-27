package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 1 << 24
	fmt.Println(f, f+1, f+1 == f)
	fmt.Println("--------")
	decl()
	fmt.Println("--------")
	declE()
	fmt.Println("--------")
	cmp()
	fmt.Println("--------")
	bl()
}

func decl() {
	a := 1.2345
	b := 1.
	c := .2345
	fmt.Println(a, b, c) // 输出: 1.2345 1 0.2345
}

func declE() {
	a := 2.234e10
	b := 3.456e-10
	fmt.Printf("%g, %g\n", a, b)

	for i := 0; i < 4; i++ {
		fmt.Printf("i = %8.3f\n", math.Exp(float64(i)))
	}

	var f float64
	fmt.Println(f, -f, 1/f, -1/f, f/f)
}

func cmp() {
	var a complex128 = complex(1, 3)
	var b complex128 = 3 + 5i
	fmt.Println(a * b)
	fmt.Println(real(a))
	fmt.Println(imag(a))

}

func bl() {
	var s string
	if (s != "" && s[0] != 'x') {
		fmt.Println(s)
	}
	
}
