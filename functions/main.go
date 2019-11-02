package main

import (
	// "io"
	// "io/ioutil"
	// "log"
	"fmt"
	"math"
	"os"
	"runtime"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// LenAndCap len and cop
func LenAndCap(list []int) (int, int) {
	return len(list), cap(list)
}

func namedReturn() (i int, s string) {
	s = "foo"
	return i, s
}

func main() {
	// mdefer()
	// ifDefer(1)

	// err := title(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// f(3)
	// printStack()
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(1)
	re()
	fmt.Println(2)
	


}

type a struct{
	i [5]int
}

func igError() {
	// os.RemoveAll()
	// dir, err := ioutil.TempDir("", "")
	// if err != nil {
	// 	// return fmt.Errorf("%v", err)
	// }
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
