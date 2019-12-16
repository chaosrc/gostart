package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

// Sprint format x
func Sprint(x interface{}) string {
	type Stringer interface {
		String() string
	}
	switch x := x.(type) {
	case Stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case float64:
		return strconv.FormatFloat(x, 'b', 6, 64)
	case bool:
		if x {
			return "true"
		} else {
			return "false"
		}
	default:
		return ""
	}
}

func typeof(x interface{}) string {
	t := reflect.TypeOf(x)
	fmt.Println(t)
	return t.String()
}
func valueof(x interface{}) string {
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println(v.String())
	fmt.Println(v.Type())

	f := reflect.ValueOf(2)
	i := f.Interface()
	fmt.Println(i.(int))

	return v.String()
}

type Foo struct {
	Foo  string
	Age  int
	List []string
	M    map[string]string
}

func main() {
	// tagServer()
	printMethods(time.Hour)
}

func xexpress() {
	foo := Foo{
		"bar",
		18,
		[]string{"123", "Hello", "good"},
		map[string]string{"name": "22", "get": "45"},
	}
	// Display("a", nil)
	// Display("Foo", foo)
	// fmt.Println()

	// Display("Stderr", os.Stderr)
	b, _ := Marshal(foo)
	fmt.Println(string(b))

	fmt.Println("---------------------------------")

	s, _ := Marshal(os.Stderr)
	fmt.Println(string(s))

	fmt.Println("---------------------------------")

	var foo2 Foo
	Unmarchal(b, &foo2)
	fmt.Println(foo2)
}

func refSet() {
	list := []string{"12", "he"}
	val := reflect.ValueOf(list)
	// pt := val.Index(1).Addr().Interface().(*string)
	// *pt = "abc"
	val.Index(1).Set(reflect.ValueOf("ABC"))
	val.Index(0).SetString("456")
	fmt.Println(list)

	var f interface{}
	f1 := reflect.ValueOf(&f).Elem()

	f1.Set(reflect.ValueOf(1))
	fmt.Println(f) // 1

	f1.Set(reflect.ValueOf("hello"))
	fmt.Println(f) // hello

	f1.SetInt(2) //panic: reflect: call of reflect.Value.SetInt on interface Value
}
