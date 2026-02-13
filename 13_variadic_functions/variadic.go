package main

import "fmt"

// declaring a variadic function that takes a variable number of arguments of any type
func Println(args ...interface{}) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
}

// we get slice of int as parameter and we can pass any number of int arguments to this function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// example of variadic function
	fmt.Println(1, 2, 3, 4, "Hello", "World")
	Println(1, 2, 3, 4, "Hello", "World", 3.14, true)

	totalSum := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", totalSum)

	nums := []int{10, 20, 30}
	// we can also pass a slice to a variadic function using ... operator
	totalSum = sum(nums...)
	fmt.Println("Sum of slice:", totalSum)
}
