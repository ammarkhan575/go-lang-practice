package main

import (
	"fmt" 
	"slices"
)

func main() {
	// slices in Go
	// a slice is a dynamically sized, flexible view into the elements of an array
	// a slice does not store any data, it just describes a section of an underlying array
	// syntax for declaring a slice in Go
	// var name []type

	// uninitialized slice
	var numbers []int
	fmt.Println(numbers == nil)

	// empty slice
	var emptySlice = []int{}
	fmt.Println(emptySlice == nil)

	// we can also declare and initialize a slice in one line
	var fruits = []string{"apple", "banana", "orange"}
	fmt.Println(fruits)
	// we can also use short hand declaration for slices
	colors := []string{"red", "green", "blue", "yellow"}
	fmt.Println(colors)

	fmt.Println("=============SLICE EXAMPLE===============")
	// slice operator
	fmt.Println(colors[0:2]) // slicing the colors slice to get the first two elements
	fmt.Println(colors[1:3]) // slicing the colors slice to get the second and third elements
	fmt.Println(colors[:3]) // slicing the colors slice to get the first three elements
	fmt.Println(colors[2:]) // slicing the colors slice to get all elements from the third element to the end
	fmt.Println(colors[:]) // slicing the colors slice to get all elements
	// we can access slice elements using their index
	fmt.Println(fruits[0])
	fmt.Println(colors[2])

	// we can also modify slice elements using their index
	fruits[1] = "grape"
	fmt.Println(fruits)
	// we can also use a for loop to iterate over a slice
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	// we can also use a range loop to iterate over a slice
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	// we can also use a range loop to iterate over a slice without using the index
	for _, value := range colors {
		fmt.Println(value)
	}
	// we can also use a range loop to iterate over a slice without using the value
	for index := range fruits {
		fmt.Println("Index:", index)
	}
	// we can also use a range loop to iterate over a slice without using the index and value

	for range colors {
		fmt.Println("This is a color.")
	}
	// we can also use the built-in append function to add elements to a slice
	fruits = append(fruits, "kiwi")
	fmt.Println(fruits)
	// we can also use the built-in copy function to copy elements from one slice to another
	// syntax for copy function in Go
	// copy(destination, source)
	// the copy function returns the number of elements copied
	dest := make([]string, len(fruits))
	numCopied := copy(dest, fruits)
	fmt.Printf("Copied %d elements: %v\n", numCopied, dest)
	// we can also use the built-in len function to get the length of a slice
	fmt.Println("Length of fruits slice:", len(fruits))
	// we can also use the built-in cap function to get the capacity of a slice

	fmt.Println("Capacity of fruits slice:", cap(fruits))

	// we can also use the built-in make function to create a slice with a specified length and capacity
	// syntax for make function in Go
	// make(type, length, capacity)
	// slicesX := make([]int, 0, 5);
	mySlice := make([]int, 5, 10)
	fmt.Println("My slice:", mySlice)
	fmt.Println("Length of my slice:", len(mySlice))
	fmt.Println("Capacity of my slice:", cap(mySlice))

	// make function initializes the elements of the slice to the zero value of the type
	var myStringSlice = make([]string, 3, 5)
	fmt.Println("My string slice:", myStringSlice)
	
	// we can also use the built-in append function to add elements to a slice created with make
	myStringSlice = append(myStringSlice, "hello", "world")
	// capacity of the slice will automatically increase if we append more elements than its current capacity
	fmt.Println("My string slice after appending:", cap(myStringSlice))
	fmt.Println("My string slice after appending:", myStringSlice)

	// it doubles the capacity of the slice when it needs to grow

	// Equality of slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	// slices cannot be compared using the == operator, it will result in a compile-time error
	// fmt.Println(slice1 == slice2) // this will cause a compile-time error
	fmt.Println(slices.Equal(slice1, slice2))

	// 2D slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D slice (matrix):", matrix)
	
}