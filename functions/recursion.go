package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func parse(r io.Reader) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse: %v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(list []string, node *html.Node) []string {
	// fmt.Println(node.Type, node.Data)
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				list = append(list, attr.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		list = visit(list, c)
	}

	return list
}

type MyError struct {
	a string
}

func (err MyError) Error() string {
	return err.a
}

func err() error {
	
	return MyError{}
}
