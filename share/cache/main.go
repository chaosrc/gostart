package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/inhies/go-bytesize"
)

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.RWMutex
}

type Func func(key string) (interface{}, error)
type Result struct {
	value interface{}
	err   error
}

type entry struct {
	re    Result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func main() {
	urls := []string{
		"https://baidu.com",
		"https://bing.com",
		"https://baidu.com",
		"https://tmall.com",
		"https://taobao.com",
		"https://jd.com",
		"https://tmall.com",
	}
	totalTime := time.Now()
	var wg sync.WaitGroup
	memo := New(httpGetBody)
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			now := time.Now()
			res, err := memo.Get(u)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s,  %s,  %s\n", u, time.Since(now), bytesize.New(float64(len(res.([]byte)))))
			wg.Done()
		}(u)
	}
	wg.Wait()

	fmt.Println("Total time: ", time.Since(totalTime))
}


func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	cache := memo.cache[key]
	if cache == nil {
		cache = &entry{ready: make(chan struct{}, 1)}
		memo.cache[key] = cache
		cache.ready <- struct{}{}
	}
	memo.mu.Unlock()

	if _, ok := <-cache.ready; ok {
		var result Result
		result.value, result.err = memo.f(key)
		cache.re = result
		close(cache.ready)
	}

	return cache.re.value, cache.re.err
}

func httpGetBody(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Println("call: ", url)
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
