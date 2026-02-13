package main

import "fmt"

func main() {
	// 1. range with arrays
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// 2. range with slices
	slice := []string{"Go", "is", "awesome"}
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	// 3. range with maps
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// 4. range with strings
	// unicode code point rune
	// starting byte of rune
	// 
	str := "Hello"
	for i, v := range str {
		// here v will be printed as a rune (Unicode code point), so we can use %c to print it as a character
		fmt.Println(i,v);
		// if we want to print the v as character we can also use 
		fmt.Println(string(v))
		fmt.Printf("Index: %d, Value: %c\n", i, v)
	}
}