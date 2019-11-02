package main

import (
	"fmt"
)

func suit(i int) {
	switch i {
	case 1: //...
	case 2: //...
	case 3: // ...
	default:
		panic("error")
	}
}

func f(x int) {
	if x < 0 {
		return
	}
	fmt.Printf("f(%d)\n", x+0)
	defer fmt.Printf("defer f(%d)\n", x)
	f(x - 1)
}


func re() {
	// defer func(){
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	panic("error")

}