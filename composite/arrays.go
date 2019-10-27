package main

import (
	"crypto/sha256"
	"fmt"
)

func arrayIndex() {
	var a [5]int
	// 获取第一个元素
	fmt.Println(a[0])
	// 获取最后一个元素
	fmt.Println(a[len(a)-1])

	// 遍历每个元素
	for k, v := range a {
		fmt.Println(k, v)
	}
}

func arrayInit() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a)
	var b = [3]int{1, 2}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", c)
}

func arraySize() {
	a := [3]int{1, 2, 3}
	// a = [4]int{1,2} //cannot use ([4]int literal) (value of type [4]int) as [3]int value in assignment

	fmt.Println(a)
}

func arrayPairs() {
	s := [...]string{1: "a", 3: "c"}
	fmt.Println(1, s[1])
	fmt.Println(2, s[2])
	fmt.Println(3, s[3])

	r := [...]int{10: -1}
	fmt.Println(r)
}

func arrayCompare() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [...]int{1, 3}

	fmt.Println(a == b, a == c) // true false
}

func hashCompare() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", a, b, a == b, a)

}

func arrayPointer(a *[8]int) {
	for i, v := range a {
		a[i] = v + 1
	}
}
