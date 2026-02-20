package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string) string {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("processing \n", url)
	return url
}

// concurrency worker
func concWorker(url string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done() // it is equivalent to wg.Add(-1)
	time.Sleep(time.Millisecond * 50)
	fmt.Println("processing \n", url)
	resultChan <- url
}

func main() {
	// this code will take 100+ms
	// startTime := time.Now()
	// image1 := worker("image1.png")
	// image2 := worker("image2.png")

	// fmt.Println(image1)
	// fmt.Println(image2)

	// fmt.Println(time.Since(startTime))

	// concurrency pattern
	// this method is called fork and join, we fork the process and then join it back together
	// also we call parallel and join, we run the process in parallel and then join it back together
	// also we call launch and wait, we launch the process and then wait for it to finish
	var wg sync.WaitGroup
	resultChan := make(chan string, 5)
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
		fmt.Printf("recieved :%s\n", r)
	}
	fmt.Println(time.Since(startTime))

}
