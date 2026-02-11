package main

import "fmt"

func main() {
	// arrays in Go
	// syntax for declaring an array in Go
	// var name [size]type
	// fixed size, cannot be changed after declaration
	// memory is allocated for all elements of the array, even if they are not initialized
	// constant time complexity for accessing elements of the array using their index
	var numbers [5]int
	fmt.Println(numbers) // prints [0 0 0 0 0], the zero value for int

	// we can also declare and initialize an array in one line
	var fruits = [3]string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// we can also use short hand declaration for arrays
	colors := [4]string{"red", "green", "blue", "yellow"}
	fmt.Println(colors)

	// we can access array elements using their index
	fmt.Println(fruits[0])
	fmt.Println(colors[2])

	// we can also modify array elements using their index
	fruits[1] = "grape"
	fmt.Println(fruits)

	// we can also declare an array without specifying its size, in this case, the size will be determined by the number of elements in the array
	animals := [...]string{"cat", "dog", "rabbit"}
	fmt.Println(animals)

	// we can also use a for loop to iterate over an array
	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	// we can also use a range loop to iterate over an array
	for index, value := range colors {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// we can also use a range loop to iterate over an array without using the index
	for _, value := range fruits {
		fmt.Println(value)
	}

	// we can also use a range loop to iterate over an array without using the value
	for index := range animals {
		fmt.Println("Index:", index)
	}

	// we can also use a range loop to iterate over an array without using the index and value
	for range colors {
		fmt.Println("This is a color.")
	}

	// 2D arrays
	// syntax for declaring a 2D array in Go
	// var name [rows][columns]type
	var matrix [3][3]int
	fmt.Println(matrix) // prints [[0 0 0] [0 0 0] [0 0 0]], the zero value for int

	// we can also declare and initialize a 2D array in one line
	var identityMatrix = [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(identityMatrix)
}