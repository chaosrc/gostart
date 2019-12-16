package main

import (
	// "log"
	"sync"
	"time"
)

func ImageFile(fileName string) (string, error) {
	time.Sleep(time.Second * 1)
	return fileName, nil
}

func makeThumbnails(filenames []string) {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	for _, name := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			ImageFile(f)
			ch <- struct{}{}
		}(name)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 等待 goroutine 完成
	for range ch {
		<-ch
	}
}
