package main 

import (
	"fmt"
	"os"
	"example.com/hello/echo"
	"example.com/hello/tempconv"
)


const PI = 3.1415

func main() {
	// var r = 10.0

	// var area = PI * r * r
	
	// fmt.Printf("面积为：%g", area)
	// fmt.Println()



	// var i int 
	// var b bool
	// var s string

	// fmt.Printf("i: %v, b: %v, s: %v", i, b, s)


	// var i, j, k int
	
	// var b, f, s = true, 1.2, "hello"


	// foo()
	// multiStort()
	// fnCall()

	// pointer()

	// fmt.Println("echo -----")
	echo.Start()
	fmt.Println(tempconv.CToF(100))
}

func foo() {
	a := 100               // int 类型
	var b float32 = 100    //  float32 类型

	var names []string     // 先声明后面在赋值

	names = []string{"a", "b"}

	fmt.Println(a, b, names)
}

func multiStort() {
	i, s := 100, "hello" 
	fmt.Println(i, s)
}

func fnCall() error {
	file, err := os.Open("./vars/bar.go")

	if err != nil {
		return err
	}
	info, err :=file.Stat()
	fmt.Println(info)
	file.Close()

	return nil
}

func short() {
	file, err := os.Open("./vars/bar.go")

	file, err = os.Create("./hello.go")

	fmt.Println(file.Name(), err)
}


func pointer() {
	x := 10
	p := &x  // p 是 *int 类型，指向 x
	
	fmt.Println(*p) // 输出： 10

	*p = 20
	fmt.Println(x) // 输出：20

	n := &p

	fmt.Println(x, p, *p, n )

	fmt.Println("------")
	var y, z int
	var t *int
	v := pf()
	fmt.Println(&y, &z, t, v, *v)


	fmt.Println("------")

	pointArgs()


}


func pf() *string {
	v := "hello"
	return &v
}

func pointArgs() {
	args := func (val *string) {
		*val = "foo"
	}
	val := "bar"
	args(&val)

	fmt.Println(val)
}

