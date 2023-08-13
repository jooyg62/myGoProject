package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	c  map[string]int
}

func (s *SafeCounter) Up(key string) {
	s.mu.Lock()
	s.c[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.c[key]
}

func main() {
	s := SafeCounter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go s.Up("someKey")
	}

	time.Sleep(time.Second)
	fmt.Println(s.Value("someKey"))
}
