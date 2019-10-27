package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func createMap() {
	// m := make(map[string]string)
	m2 := map[string]string{
		"name":    "foo",
		"address": "shanghai",
	}
	m2["name"] = "bar"
	m2["phone"] = "12345678"
	delete(m2, "name1")
	fmt.Printf("%d\n", len(m2))
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	fmt.Println("---------")

	var ages map[string]int
	fmt.Println(ages == nil, len(ages))

	ages["foo"] = 20
}

func utf() {
	fmt.Println(utf8.UTFMax)
}

func count(str string) map[string]int {
	m := make(map[string]int)

	for _, s := range str {
		if unicode.IsLetter(s) {
			m["L"]++
		} else if unicode.IsDigit(s) {
			m["D"]++
		}
	}
	return m
}

func countInput() map[string]int {
	m := make(map[string]int)

	reader := bufio.NewReader(os.Stdin)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		if unicode.IsDigit(r) {
			m["D"]++
		}
		if unicode.IsLetter(r) {
			m["L"]++
		}
	}
	return m
}
