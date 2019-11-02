package main

import (
	"fmt"

	"example.com/hello/composite/github"
)

func main() {
	re, err := github.SearchIssues([]string{"vue", "vuex"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(re.TotalCount)
	// fmt.Println(re.Items[:2])
	// jsonMovie()
	github.Text(re)

	github.HTML(re)
}

func listTest(li []int) {
	li = append(li, 2)
}
