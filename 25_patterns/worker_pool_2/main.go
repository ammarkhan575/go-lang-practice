package main

import (
	"sync"
	"time"
	"fmt"
)

type Result struct {
	Value string
	Err error
}

func worker(id int, jobChan chan string, resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobChan {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Started worker: ", id, " job: ", job)
		resultChan <- Result{Value: job, Err: nil}
	}
	fmt.Println("Worker shutting down: ", id)
}

func main() {
	jobs := []string{
		"image1.png",
		"image2.png",
		"image3.png",
		"image4.png",
		"image5.png",
		"image6.png",
		"image7.png",
		"image8.png",
		"image9.png",
		"image10.png",
		"image11.png",
		"image12.png",
		"image13.png",
		"image14.png",
		"image15.png",
		"image16.png",
		"image17.png",
		"image18.png",
		"image19.png",
		"image20.png",
	}
	const totalWorker = 5

	resultChan := make(chan Result, 50)
	jobChan := make(chan string, len(jobs))
	var wg sync.WaitGroup
	startTime := time.Now()
	for i := 1; i<= totalWorker; i++ {
		wg.Add(1)
		go worker(i, jobChan, resultChan, &wg);
	}

	// below code is used to close the result channel once all the workers are done, 
	// otherwise below loop keeps reading from the channel and never exits
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// send the jobs to the job channel, once all the jobs are sent, we can close the job channel, 
	// this will signal the workers to exit the loop and shut down
	for i := 0; i < len(jobs); i++ {
		jobChan <- jobs[i]
	}
	close(jobChan)


	// this is called fan in, we have multiple workers sending the results to the result channel, 
	// and we are reading from the result channel in a single loop, this is called fan in pattern
	for r := range resultChan {
		fmt.Printf("job completed ✅ :%v\n", r)
		if r.Err != nil {
			// reprocess - retry logic
			// reprocess - queue - dlq(dead letter queue)
			fmt.Printf("error processing url: %s, error: %v\n", r.Value, r.Err)
		}
	}
	fmt.Println("time taken: ", time.Since(startTime))
	 
}