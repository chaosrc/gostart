package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// "example.com/hello/methods/geometry"
type B struct {
	a int
}
type C struct {
	a int
	b string
}

func main() {
	db := database{
		"shoe":  100,
		"socks": 5,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatalln(http.ListenAndServe("localhost:8001", db))
}

func execReader() {
	s := `Implementation of the Read() 
	method should return the number of bytes read 
	or an error if one occurred. If the source has exhausted its content,
	 Read should return io.EOF
`

	r := NewStringReader(s)

	io.Copy(os.Stdout, r)

	fmt.Println("--- string reader ---")

	r2 := strings.NewReader(s)
	io.Copy(os.Stdout, r2)
}

func sortTest() {
	a := []string{"Java", "C", "Python", "Javascript", "Kotlin"}

	sort.Sort(StringList(a))
	fmt.Printf("%q\n", a)

	tracks := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, time.Minute * 3},
		{"Go", "Moby", "Moby", 1992, time.Minute * 2},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, time.Minute * 5},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, time.Minute * 4},
	}

	sort.Sort(customSort{tracks, func(a, b *Track) bool {
		if a.Title != b.Title {
			return a.Title < b.Title
		}
		if a.Year != b.Year {
			return a.Year < b.Year
		}

		return a.Length < b.Length
	}})

	for _, val := range tracks {
		fmt.Println(*val)
	}
}

// type aa func(q geometry.Point) float64
