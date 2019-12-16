package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("%s %T:\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(name string, v reflect.Value) {

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s=%q\n", name, "Invalid")
	case reflect.String:
		fmt.Printf("%s=%q\n", name, v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := strconv.FormatInt(v.Int(), 10)
		fmt.Printf("%s=%s\n", name, i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		i := strconv.FormatUint(v.Uint(), 10)
		fmt.Printf("%s=%s\n", name, i)
	case reflect.Bool:
		b := strconv.FormatBool(v.Bool())
		fmt.Printf("%s=%s\n", name, b)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldName := v.Type().Field(i).Name
			display(fmt.Sprintf("%s.%s", name, fieldName), v.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", name, i), v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", name, key), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s=nil\n", name)
		} else {
			display(fmt.Sprintf("(*%s)", name), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s=nil\n", name)
		} else {
			display(name, v.Elem())
		}
	case reflect.Func, reflect.Chan:
		fmt.Printf("%s=%s", name, strconv.FormatUint(uint64(v.Pointer()), 16))
	default:
		fmt.Printf("%s.type=%s\n", name, v.Type())
	}
}
