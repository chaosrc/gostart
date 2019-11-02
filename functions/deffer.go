package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func title(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	// 使用 defer 关闭请求
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		// res.Body.Close()
		return fmt.Errorf("Not HTML")
	}

	doc, err := html.Parse(res.Body)
	// res.Body.Close()
	if err != nil {
		return fmt.Errorf("HTML parse failed")
	}

	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			fmt.Println(node.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)

	return nil

}

func forEachNode(node *html.Node, pre func(*html.Node), post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		forEachNode(n, pre, post)
	}
	if post != nil {
		post(node)
	}
}

func mdefer() {
	fmt.Println(1)
	defer fmt.Println("defer 1")
	fmt.Println(2)
	defer fmt.Println("defer 2")
	fmt.Println(3)
}

func ifDefer(i int) {
	if i > 0 {
		x := 5
		defer func() {
			fmt.Println(x)
		}()
	}

}
