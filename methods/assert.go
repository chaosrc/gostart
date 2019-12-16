package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func assert() {
	var w io.Writer

	w = os.Stdout
	rw := w.(*os.File)
	fmt.Println(w == rw)

	type b interface {
	}

	// if ok {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("fail")
	// }

	// c := a.(*bytes.Buffer)
	// fmt.Println(c == a) // panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer

}

func writeString(w io.Writer, str string) (int, error) {
	type WriteString interface {
		WriteString(string) (int, error)
	}
	if w, ok := w.(WriteString); ok {
		return w.WriteString(str)
	}
	return w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, contentType string) error {
	if _, err := writeString(w, "Content-Type"); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}

	//....

	return nil
}
