package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func waitForServer(url string) (io.Reader, error) {
	timeout := time.Minute * 1
	deadline := time.Now().Add(timeout)

	for retry := 0; time.Now().Before(deadline); retry++ {
		res, err := http.Get(url)
		if err == nil && res.StatusCode == http.StatusOK {
			return res.Body, nil
		}
		fmt.Println(err, retry)
		time.Sleep(time.Second << uint(retry))
	}
	return nil, fmt.Errorf("Failed %s timeout %s", url, timeout)
}


