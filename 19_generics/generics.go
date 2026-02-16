package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we can use generics to write a single function that can work with any type of slice
func printGenericSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// // or we can write this as
// func printGenericSlice1[T interface{}] (items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// suppose we want to write a function that accept only integer and string
func printIntOrStringSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we can also use comparable constraint to write a function that can work with any type that is comparable
// comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). 
// The comparable interface may only be used as a type parameter constraint, not as the type of a variable.

func printComparableSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we can also define multiple type parameters in a function
func printMultipleTypeParameters[T1 any, T2 any](items1 []T1, items2 []T2) {
	fmt.Println("Items of type T1:")
	for _, item := range items1 {
		fmt.Println(item)
	}
	fmt.Println("Items of type T2:")
	for _, item := range items2 {
		fmt.Println(item)
	}
}

// use of generic in structs
type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strings := []string{"Go", "is", "awesome"}
	printSlice(nums)
	printStringSlice(strings)

	// for every type we have to write a separate function, this is not efficient and leads to code duplication
	// we can use generics to write a single function that can work with any type of slice
	printGenericSlice(nums)
	printGenericSlice(strings)

	myStack := stack[int] {
		elements : []int{1,2,3},
	}
	fmt.Printf("Stack of integers: %+v\n", myStack)

	myStringStack := stack[string] {
		elements : []string{"Hello", "World"},
	}
	fmt.Printf("Stack of strings: %+v\n", myStringStack)


}
