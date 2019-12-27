package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/moby/moby/pkg/pubsub"
)

func main() {
	publisher := pubsub.NewPublisher(time.Millisecond*100, 10)

	golang := publisher.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	docker := publisher.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go publisher.Publish("hello")
	go publisher.Publish("golang: golang.org")
	go publisher.Publish("docker: docker.com")

	time.Sleep(time.Second * 1)

	go func() {
		fmt.Println("go topic", <-golang)
	}()
	go func() {
		fmt.Println("docekr topic", <-docker)
	}()

	time.Sleep(time.Second * 5)

}
