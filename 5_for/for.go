package main

import "fmt"

// for only construct in Go, it can be used as a while loop or an infinite loop as well
// there is no while loop in Go, we can use for loop to achieve the same functionality
func main() {
	// for loop in Go
	// syntax for for loop in Go
	// for initialization; condition; post {
	//     // code to be executed
	// }

	// while loop using for
	// i := 0
	// for i < 5 {
	//     fmt.Println(i)
	//     i++
	// }

	// infinite loop using for
	// for {
	//     // code to be executed
	// }

	// class for loop that prints numbers from 0 to 4
	// example of a for loop that prints numbers from 0 to 4
	for i := 0; i < 5; i++ {
		// break statement to exit the loop when i is equal to 3
		if i == 3 {
			break
		}
		// continue statement to skip the rest of the loop when i is equal to 2
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// range based for loop in Go
	// syntax for range based for loop in Go
	// for index, value := range collection {
	//     // code to be executed
	// }
	// it will iterate to 2 because we are using break statement when i is equal to 3, so it will exit the loop when i is equal to 3
	for i := range 3 {
		fmt.Printf("Index: %d\n", i);
	}
	// example of a for loop that iterates over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// example of a for loop that iterates over a map
	person := map[string]string{
		"name": "Ammar",
		"age": "25",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// example of a for loop that iterates over a string
	str := "Hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// infinite loop using for
	// for {
	//     // code to be executed
	// }
}