package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s = struct{}{}
	fmt.Println(unsafe.Sizeof(s))
	//输出： 8

	var x struct {
		a bool
		b int16
		c []int
	}

	fmt.Println(unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Println(unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))

	fmt.Printf("%#016x\n", Float64bits(1.0))
	
	// 相当于 &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)) )
	*pb = 42
	fmt.Println(x.b)
	
}

func Float64bits(f float64)uint64 {
	
	return *(*uint64)(unsafe.Pointer(&f))
}
