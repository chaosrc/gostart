package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type KVService struct {
	m      map[string]string
	filter map[string]func(string)
	mu     sync.Mutex
}

func NewKVService() *KVService {
	return &KVService{
		m:      make(map[string]string),
		filter: make(map[string]func(string)),
	}
}

func (s *KVService) Get(key string, repaly *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if value, ok := s.m[key]; ok {
		*repaly = value
		return nil
	}

	return fmt.Errorf("%s Not Found\n", key)
}

func (s *KVService) Set(kv [2]string, replay *struct{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	key, value := kv[0], kv[1]
	fmt.Println(s.filter)
	if oldValue := s.m[key]; oldValue != value {
		for _, f := range s.filter {
			f(key)
		}
	}
	s.m[key] = value
	return nil

}

func (s *KVService) Watch(timeout int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	s.mu.Lock()
	s.filter[id] = func(key string) { 
		ch <- key
	 }
	s.mu.Unlock()

	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}

	return nil

}
