/**
 * go 互斥锁
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发安全
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// 增加计数器值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// 返回当前计数器值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock() // 在Value 函数返回时解锁
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey")) //1000
}
