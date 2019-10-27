package main 

import (
	"example.com/hello/tempconv"
	"os"
	"strconv"
	"fmt"
)

var a = b + c
var b = f()
var c = 1
func f() int {
	return c + 1
}

func init() {
	a = 10
	fmt.Printf("init: %v", a)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)

		fmt.Printf("%s = %s, %s = %s \n", 
			c, tempconv.CToF(c), f, tempconv.FToC(f))

			
	}
}

func tt(i interface{}) {
	
	switch t := i.(type) {
	case bool:
		fmt.Println(t)
	}
}