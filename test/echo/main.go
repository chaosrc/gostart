package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "省略换行符")
	s = flag.String("s", " ", "分割符")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	if err := echo(*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
	}
}

func echo(n bool, s string, arg []string) error {
	fmt.Fprint(out, strings.Join(arg, s))
	if !n {
		fmt.Fprintln(out)
	}

	return nil
}


func TEcho() {
	fmt.Println("TEcho")
}