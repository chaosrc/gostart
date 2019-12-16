package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(buf, "%s", "nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Array, reflect.Slice:
		buf.WriteString("(")
		for i := 0; i < v.Len(); i++ {
			encode(buf, v.Index(i))
			if i != v.Len()-1 {
				buf.WriteByte(' ')
			}
		}
		buf.WriteString(")")
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			name := v.Type().Field(i).Name
			fmt.Fprintf(buf, "(%s ", name)
			encode(buf, v.Field(i))
			buf.WriteString(")")
			if i != v.NumField()-1 {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			fmt.Fprintf(buf, "(%s ", key)
			encode(buf, v.MapIndex(key))
			buf.WriteByte(')')

			if i != len(v.MapKeys())-1 {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte(')')
	case reflect.Ptr:
		encode(buf, v.Elem())
	default:
		// 不支持的类型 Func chan interface 等
		return fmt.Errorf("unsupport type %s", v.Type())

	}

	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := encode(&buf, reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
