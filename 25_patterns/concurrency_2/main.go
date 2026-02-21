package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value string
	Err error
}

func worker(url string) string {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("processing \n", url)
	return url
}

// concurrency worker
func concWorker(url string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done() // it is equivalent to wg.Add(-1)
	time.Sleep(time.Millisecond * 50)
	fmt.Println("processing \n", url)
	resultChan <- Result{Value: url, Err: nil}
}

func main() {
	// concurrency pattern
	// this method is called fork and join, we fork the process and then join it back together
	// also we call parallel and join, we run the process in parallel and then join it back together
	// also we call launch and wait, we launch the process and then wait for it to finish
	var wg sync.WaitGroup
	resultChan := make(chan Result, 5)
	urls := []string{"image1.png", "image2.png", "image3.png", "image4.png", "image5.png"}
	wg.Add(len(urls)) // wg.Add(5)
	startTime := time.Now()
	// this is called fan out
	for _, url := range urls {
		go concWorker(url, &wg, resultChan)
	}
	wg.Wait()

	close(resultChan) // we have to close the channel otherwise below loop keeps reading from the channel
	// this is called fan in
	for r := range resultChan {
		fmt.Printf("recieved :%v\n", r)
		if r.Err != nil {
			// reprocess - retry logic
			// reprocess - queue - dlq(dead letter queue)
			fmt.Printf("error processing url: %s, error: %v\n", r.Value, r.Err)
		}
	}
	fmt.Println(time.Since(startTime))

}
