package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	// receive a number from the channel
	// num := <-numChan
	// fmt.Println("Received number:", num)

	// if we want to get multiple numbers from the channel, we can use a loop
	for val := range numChan {
		fmt.Println("Received number:", val)
		time.Sleep(time.Second * 1)
	}
}

func sum(result chan int, a int, b int) {
	numResult := a + b
	result <- numResult // send the result to the channel
	// time.Sleep(time.Second * 10)
}

func taskDone(done chan bool) {
	defer func() { done <- true }() // signal that this goroutine is done when it finishes
	fmt.Println("processing task ...")
}

func emailSender(emailChan chan string, emailDone chan bool) {
	// defer close(emailChan) // close the channel when the function finishes
	defer func() { emailDone <- true }() // signal that this goroutine is done when it finishes
	for email := range emailChan {
		fmt.Println("Sending email to:", email)
		// simulate email sending time
		time.Sleep(time.Second)
	}
}

func main() {
	// a channel is a way to communicate between goroutines and synchronize their execution
	// we can create a channel using the make function
	ch := make(chan string)

	// we can send a value to the channel using the <- operator
	go func() {
		ch <- "Hello from goroutine!"
	}()

	// we can receive a value from the channel using the <- operator
	message := <-ch
	fmt.Println(message)

	// we can also use channels to synchronize the execution of goroutines
	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine is doing some work...")
		done <- true // signal that the work is done
	}()

	<-done // wait for the signal that the work is done
	fmt.Println("Main function is exiting...")

	// Send / Receive
	// ch <- 10   // send
	// x := <-ch  // receive

	numChan := make(chan int)
	// numChan <- 109 // this will cause a deadlock because there is no goroutine receiving from the channel
	// to fix this issue, we can start a goroutine that will receive from the channel
	go processNum(numChan)
	// numChan <- 99 // now we can send a number to the channel without causing a deadlock

	// we can also send multiple numbers to the channel using a loop
	for i := 0; i < 5; i++ {
		numChan <- rand.Intn(100)
	}
	close(numChan) // close the channel to signal that no more values will be sent

	result := make(chan int)
	go sum(result, 10, 20)
	sumResult := <-result // block until we receive the result from the sum function
	fmt.Println("Sum result:", sumResult)

	done1 := make(chan bool)
	go taskDone(done1)

	// this will block until we receive a signal from the taskDone goroutine that it has finished processing the task
	<-done1 // wait for the signal that the task is done
	fmt.Println("Main function is exiting...")

	// Example of buffered channel
	// bufferedChan := make(chan int, 3) // create a buffered channel with a capacity of 3
	// bufferedChan <- 1 // send a value to the channel
	// bufferedChan <- 2 // send another value to the channel
	// bufferedChan <- 3 // send another value to the channel
	// fmt.Println("Buffered channel is full, sending another value will block...")
	// bufferedChan <- 4 // this will block until there is space in the channel

	/*
	emailChannel := make(chan string, 100) // create a buffered channel with a capacity of 2
	emailDone := make(chan bool)
	go emailSender(emailChannel, emailDone) // start a goroutine to send emails from the channel

	for i := 1; i <= 100; i++ {
		email := fmt.Sprintf("%d@gmail.com", i)
		emailChannel <- email // send email addresses to the channel
		fmt.Println("Sent email to channel:", email)
	}
	// this is important to close the channel after we are done sending values to it,
	// otherwise the emailSender goroutine will be stuck waiting for more values to be sent to the channel and will never finish
	close(emailChannel) // close the channel to signal that no more emails will be sent

	// fmt.Println("Email channel is full, sending another value will block...")
	// no block here because we buffered channel has a capacity of 2 and we have only sent 2 values to it, so there is still space in the channel
	// fmt.Println("Email channel contents:", <-emailChannel, <-emailChannel) // receive values from the channel
	// fmt.Println("Email channel contents:", <-emailChannel) // receive value from the channel
	<-emailDone // wait for the emailSender goroutine to finish
	*/
	chan1 := make(chan int)
	chan2 := make(chan string)

	// anonymous goroutine to send a value to chan1
	go func() {
		chan1 <- 99
	}()

	go func() {
		chan2 <- "Hello from chan2!"
	}()

	// we can use select and for loop to wait on multiple channels
	for i := 0; i < 2; i++ {
		select {
		case val := <-chan1:
			fmt.Println("Received from chan1:", val)
		case val := <-chan2:
			fmt.Println("Received from chan2:", val)
		}
	}

}


// Go has built-in race detection that can help identify race conditions in your code. You can run your Go program with the -race flag to enable race detection. For example:
// go run -race your_program.go
// This will analyze your code for potential race conditions
// If race exists:
// WARNING: DATA RACE

// 👉 Difference between buffered and unbuffered channel?
// Unbuffered channels block the sender until a receiver is ready,
// providing synchronization. Buffered channels allow sending without blocking until the buffer is full,
// acting like a queue and improving throughput.

// Unbuffered Channels
// Unbuffered Channel (capacity = 0)
// ch := make(chan int)

// Behavior
// Send blocks until receiver is ready
// Receive blocks until sender sends
// Both goroutines must be ready at the same time

// Flow
// Sender  ──send──▶ (wait) ◀──receive── Receiver

// Mental Model
// Unbuffered
// Like:
// 👉 Passing a ball directly
// You must both be present.

// 👉 Use Unbuffered when
// Need strict synchronization
// Need ordering guarantee
// Worker coordination
// Signaling
// 👉 Examples:
// done signals
// locking behavior
// pipelines

// Buffered Channels
// Buffered Channel (capacity > 0)
// ch := make(chan int, 3)

// Behavior
// Send does NOT block until buffer is full
// Receive does NOT block until buffer is empty
// Works like a queue (FIFO)

// Flow
// Sender ──▶ [ buffer buffer buffer ] ──▶ Receiver

// Mental Model
// Buffered
// Like:
// 👉 Dropping letters in a mailbox
// Receiver can collect later.

// 👉 Use Buffered when
// Producer faster than consumer
// Batch processing
// Rate smoothing
// Queue systems
// 👉 Examples:
// logging
// job queues
// event batching

// 👉 defer
// defer is used to schedule a function call to run after the surrounding function returns.
// It’s mostly used for cleanup, safety, and resource management.
// Order of execution: LIFO (Last In, First Out)
// func main() {
//     defer fmt.Println("World")
//     fmt.Println("Hello")
// }
// Execution order:
// defer registered
// print Hello
// function ends
// deferred runs

// multiple defer
// func main() {
//     defer fmt.Println(1)
//     defer fmt.Println(2)
//     defer fmt.Println(3)
// }
// output:
// 3
// 2
// 1

// defer with return
// 👉 Deferred functions run AFTER return statement but BEFORE function actually exits
// func test() int {
//     defer fmt.Println("deferred")
//     return 10
// }
// output:
// deferred

// 👉 Flow:
// return 10 prepared
// defer runs
// function exits

// defer evaluates arguments immediately
// func main() {
//     x := 10
//     defer fmt.Println(x)
//     x = 20
// }
// output:
// 10
// Explanation:
// defer captures the value of x at the time defer is called, which is 10.
// Even though x is later changed to 20, the deferred function will print 10 because it uses the captured value.

// But closures behave differently
// func main() {
//     x := 10
//     defer func() {
//         fmt.Println(x)
//     }()
//     x = 20
// }
// output:
// 20
// Why?
// Closure captures variable, not value.

// Using defer for timing
// func main() {
//     start := time.Now()
//     defer func() {
//         fmt.Println("Time:", time.Since(start))
//     }()
//     time.Sleep(2 * time.Second)
// }
// Output:
// Time: 2s
// Explanation:
// defer is used to measure the time taken by the function to execute.
// The deferred function captures the start time and calculates the elapsed time when it runs after the main function finishes.

// 👉 defer is a powerful tool for managing resources and ensuring that cleanup code runs even in the presence of errors or panics.
// It’s commonly used for closing files, releasing locks, and other cleanup tasks that must happen regardless of how the function exits.

// defer inside loop
// for i := 0; i < 3; i++ {
//     defer fmt.Println(i)
// }
// Output:
// 2
// 1
// 0
// Explanation:
// Each defer captures the current value of i at the time it is called.
// When the loop finishes, the deferred calls are executed in reverse order, printing 2, then 1, then 0.
