package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	// bytesize "github.com/inhies/go-bytesize"
)

func du(filepath string, fileSize chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	for _, file := range readDir(filepath) {
		if file.IsDir() {
			n.Add(1)
			go du(path.Join(filepath, file.Name()), fileSize, n)
		} else {
			fileSize <- file.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func readDir(filepath string) []os.FileInfo {
	sema <- struct{}{}
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	<-sema
	return files
}

func du2(filepath string) int64 {
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	var size int64
	for _, file := range files {
		if file.IsDir() {
			size += du2(path.Join(filepath, file.Name()))
		} else {
			size += file.Size()
		}

	}
	// fmt.Printf("%s, %s \n", filepath, bytesize.New(float64(size)))
	return size
}
