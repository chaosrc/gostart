package main

import (
	"flag"
	"fmt"
	// "sync"
	"time"

	bytesize "github.com/inhies/go-bytesize"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	for _, path := range roots {
		size := du2(path)
		fmt.Printf("%s, %s\n", path, bytesize.New(float64(size)))
	}
	
	// var n sync.WaitGroup
	// for _, root := range roots {
	// 	n.Add(1)
	// 	go du(root, sizeChan, &n)
	// }
	// go func() {
	// 	n.Wait()
	// 	close(sizeChan)
	// }()

	// var size, nfiles int64

	// // for s := range sizeChan {
	// // 	size += s
	// // 	nfiles++
	// // }
	// // fmt.Printf("%d files, %s\n", nfiles, bytesize.New(float64(size)))

	// tick := time.Tick(500 * time.Millisecond)

	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Printf("%d files, %s\n", nfiles, bytesize.New(float64(size)))

	// 	case s, ok := <-sizeChan:
	// 		if !ok {
	// 			fmt.Printf("%d files, %s\n", nfiles, bytesize.New(float64(size)))
	// 			return
	// 		}
	// 		size += s
	// 		nfiles++
	// 	}
	// }
}

func routine() {
	go Spinner(time.Millisecond * 100)
	n := Fib(45)

	fmt.Printf("\rFib(45)=%d\n", n)
}

func parallel() {
	url := []string{"https://baidu.com"}

	workList := make(chan []string)

	go func() {
		workList <- url
	}()

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
