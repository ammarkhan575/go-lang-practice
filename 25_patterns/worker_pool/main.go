package main

import (
	"math/rand"
	"fmt"
	"time"
)

func worker(id int, jobChan chan int, resultChan chan int) {
	for j:= range jobChan {
		fmt.Println("Started worker: ", id, " job: ", j)
		time.Sleep(time.Second)
		fmt.Println("Completed by worker: ", id, " job: ", j)
		resultChan <- j * rand.Intn(5)
	}
}

func main() {
	startTime := time.Now()
	const numJobs = 5
	jobsChan := make(chan int, numJobs)
	resultChan := make(chan int, numJobs)

	for w := 1; w <=3; w++ {
		go worker(w, jobsChan, resultChan)
	}

	for j := 1; j<= numJobs; j++ {
		jobsChan <- j
	}
	
	close(jobsChan)

	for a := 1; a <= numJobs; a++ {
        <-resultChan
    }

	fmt.Println("time taken: ", time.Since(startTime))
}
