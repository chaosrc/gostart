package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(format("12345"))
	conv()
	fmt.Println("int to string ------------")
	// iToS()
	const1()
}

func format(s string) string {

	l := len(s)
	if l < 3 {
		return s
	}
	return format(s[:l-3]) + "," + s[l-3:]
}

func conv() {
	s := "hello"
	b := []byte(s)
	b[2] = byte(101)
	s2 := string(b)

	fmt.Println(s, b, s2)
}

func intsToString(values []int) string {
	buf := bytes.Buffer{}

	buf.WriteString("[")
	for i,value := range values {
		if i > 0 {
			buf.WriteString(", ")
		}

		buf.WriteString(strconv.Itoa(value))
	}
	buf.WriteString("]")

	return buf.String()
}

func iToS() {
	i := 10
	s1 := fmt.Sprintf("%d", 10)
	s2 := strconv.Itoa(i)

	fmt.Println(s1, s2)

	fmt.Println(strconv.FormatInt(10, 2))
	fmt.Println(strconv.FormatInt(63, 16))

	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
		return
	}


	y, err := strconv.ParseInt("123", 10, 64)

	fmt.Println(x, y)
	fmt.Printf("%T\n", time.Minute)
}

func const1() {
	const delay time.Duration = 100
	const timeout = time.Minute * 5

	fmt.Printf("%T, %[1]v\n", delay)
	fmt.Printf("%T, %[1]v\n", timeout)

	// const (
	// 	a = 1
	// 	b
	// 	d = 2
	// 	e
	// )

	// fmt.Println(a, b, d, e)

	const (
		Sunday = iota * 10
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		a = 10 * iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
	const t = 1.00001000000000000000000000000000001
}
