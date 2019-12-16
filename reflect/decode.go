package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

type lexer struct {
	scanner scanner.Scanner
	token   rune
}

func (lex *lexer) next() {
	lex.token = lex.scanner.Scan()
}
func (lex *lexer) text() string {
	return lex.scanner.TokenText()
}

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	fmt.Printf("unkown token %s", lex.text())
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')

			field := lex.text()
			lex.next()
			read(lex, v.FieldByName(field))

			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')

			key := reflect.New(v.Type().Elem()).Elem()
			read(lex, key)

			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)

			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		// panic(fmt.Sprintf("decode list fail %v", v))
		fmt.Printf("decode list fail %v", v)
	}

}

func endList(lex *lexer) bool {
	if lex.token == scanner.EOF {
		panic("end of file")
	}
	if lex.token == ')' {
		return true
	}
	return false
}

func Unmarchal(data []byte, out interface{}) error {
	lex := &lexer{scanner: scanner.Scanner{Mode: scanner.GoTokens}}

	lex.scanner.Init(bytes.NewReader(data))
	lex.next()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}
