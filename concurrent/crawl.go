package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)
var tokens = make(chan struct{}, 20)
func crawl(url string) []string {
	tokens <- struct{}{}
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		fmt.Println(err)
	}
	<- tokens
	return list
}

func extract(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	links := []string{}
	findLinks(doc, &links)
	return links, nil
}

func findLinks(node *html.Node, links *[]string) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				*links = append(*links, attr.Val)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		findLinks(child, links)
	}
}
