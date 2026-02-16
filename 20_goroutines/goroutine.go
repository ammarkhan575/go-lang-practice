package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Task", id, "is running");
}

func main() {
	// a goroutine is a lightweight thread of execution managed by the Go runtime
	// we can create a goroutine using the go keyword followed by a function call
	// go func() {
	// 	fmt.Println("Hello from the goroutine!")
	// }()

	// // we can also create a goroutine using a named function
	// go greet()

	// // we need to wait for the goroutines to finish before exiting the main function
	// // otherwise, the main function will exit before the goroutines have a chance to run
	// // we can use time.Sleep to wait for a short period of time
	// // but this is not a reliable way to wait for goroutines to finish
	// // instead, we can use sync.WaitGroup to wait for multiple goroutines to finish
	// var wg sync.WaitGroup
	// wg.Add(2) // we have 2 goroutines to wait for

	// go func() {
	// 	defer wg.Done() // signal that this goroutine is done when it finishes
	// 	fmt.Println("Hello from the first goroutine!")
	// }()

	// go func() {
	// 	defer wg.Done() // signal that this goroutine is done when it finishes
	// 	fmt.Println("Hello from the second goroutine!")
	// }()

	// wg.Wait() // wait for all goroutines to finish

	// for i:=1; i<=10;i++ {
	// 	go task(i)
	// }

	// using anonymous function to create a goroutine
	for i:=1;i<=5;i++ {
		go func(i int){
			fmt.Println("Hello from anonymous goroutine", i)
			// here we are using the loop variable i inside the anonymous function, but since the goroutine runs asynchronously, 
			// by the time the goroutine executes, the value of i will have changed to 6, so all goroutines will print "Hello from anonymous goroutine 6"
			// to fix this issue, we can pass the loop variable as an argument to the anonymous function
		}(i)
	}

	time.Sleep(time.Second * 2) // wait for 2 seconds to allow all goroutines to finish	
 
}
