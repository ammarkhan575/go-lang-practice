package main

import (
	"fmt"
	"sync"
	"time"
)

type post struct {
	views int
	mu    sync.Mutex
}

/**
Without mutex
type Counter struct {
	v map[string]int
}

func (c *Counter) Inc(key string) {
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	return c.v[key]
}
**/

type SafeCounter struct {
	v  map[string]int
	mu sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func (p *post) increaseView(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1

	// don't lock whole code because it will decrease the performance of the code and it will cause contention between goroutines
	// we should lock only the critical section of the code which is the code that updates the shared variable post1.views

	// this is bad approach suppose anything happens which does not allow the code to execute after lock then the lock will never be released and other goroutine will be waiting for the lock to be released and it will cause deadlock
	// to avoid this we can use defer to release the lock after the function completes
	// p.mu.Unlock() // this is bad approach
}

func main() {
	post1 := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// race condition will occur here because multiple goroutines are trying to update the same variable post1.views
		// to avoid this we can use mutex
		go post1.increaseView(&wg)
	}
	// wait untill all goroutine completes
	wg.Wait()
	fmt.Println(post1.views, "post1 views")

	// Example 1:
	// without mutex
	// c1 := Counter{v: map[string]int{}}
	// for i := 0; i < 1000; i++ {
	// 	go c1.Inc("key1")
	// }
	// fmt.Println(c1.Value("key1")) // Output: 2

	// with mutex
	c2 := SafeCounter{v: map[string]int{}} // SafeCounter{v: make(map[string]int)} --- IGNORE ---
	for i := 0; i < 1000; i++ {
		go c2.Inc("key1")
	}
	time.Sleep(1 * time.Second) // wait for all goroutines to finish, this is not a good approach but we can use it for simplicity, in real world we should use sync.WaitGroup to wait for all goroutines to finish
	// best approach is to use sync.WaitGroup to wait for all goroutines to finish, but for simplicity we can use time.Sleep here
	fmt.Println(c2.Value("key1")) //
}
