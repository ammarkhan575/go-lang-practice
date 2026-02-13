package main

import "fmt"

func counterFunc () func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// a closure is a function that captures the variables from its surrounding scope
	// and can access and modify those variables even after the outer function has finished executing

	// example of a closure that captures a variable from its surrounding scope
	counter := 0
	increment := func() int {
		counter++
		return counter
	}

	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3

	// we can also create a closure that takes parameters and captures variables from its surrounding scope
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // Output: 10
	fmt.Println(triple(5)) // Output: 15

	incrementCounter := counterFunc()
	fmt.Println(incrementCounter()) // Output: 1
	fmt.Println(incrementCounter()) // Output: 2
	fmt.Println(incrementCounter()) // Output: 3	
}
