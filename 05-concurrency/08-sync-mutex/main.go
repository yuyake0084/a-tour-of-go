package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// 指定されたキーのカウンターを増分する
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lockを取得した後にマップにアクセスする
	c.v[key]++
	c.mu.Unlock()
}

// 指定されたキーの現在のカウンターの値を取得する
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for range 1000 {
		go c.Inc("some-key")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("some-key"))
}