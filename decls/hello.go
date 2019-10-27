package main

import (
	"fmt"
)

func f() int {
	return 1
}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f) //输出：f；main 函数中的 f
	fmt.Println(g) //输出：g；包级别的变量 g
	blocks()
	fmt.Println("----------")
	ifBlock()
	fmt.Println("local----")
	local()
	local2()
	fmt.Println("int------------")
	intFn()
	fmt.Println("bits------------")
	bitfn()
}

func blocks() {
	x := "hello"
	for i := 0; i < len(x); i++ {
		var x = x[i]
		if x != '!' {
			var x = x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()
}

func h(x int) int {
	return 1
}

func ifBlock() {
	if x := f(); x == 0 {
		var z = 100
		fmt.Println(x, z)
	} else if y := h(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}

	// fmt.Println(x, y) // error: undeclared name
}

var l = 1

func local() {
	l, j := 2, 3

	fmt.Println(l, j)
}

func local2() {
	var j int
	l, j = 2, 3

	fmt.Println(j)
}

func intFn() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 输出：255 0 1

	var i int8 = 127

	fmt.Println(i) // 输出：127

}

func bitfn() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	var z int8 = -1

	fmt.Printf("%08b\n", z)
	fmt.Printf("%08b\n", z << 1)
	fmt.Println(z << 2)
	fmt.Printf("%08b\n", z >> 1)

}

func intCov() {
	var apple int64 = 1
	var orange int32 = 2

	// c := 0o777

	var mix2 int = int(apple)  + int(orange)
	fmt.Println(mix2)
	// var mix int = apple + orange // error: invalid operation: mismatched
	d := [3]string{"1","2","3"}
	g := d[2:]
	fmt.Println(d,g)
}
