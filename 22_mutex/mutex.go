package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) increaseView(wg *sync.WaitGroup) {
	defer func () {
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
}
