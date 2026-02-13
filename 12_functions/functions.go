package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
// all the parameters of the same type can be written in a shorter way
func add1(a, b int) int {
	return a + b
}

func greet() {
	fmt.Println("Hello, welcome to Go programming!")
}

func getLanguage() (string, string, string, bool) {
	return "Go", "C++", "Python", true
}

// function with function as parameter
func executeOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	// 1. function with no parameters and no return value
	greet()

	// 2. function with parameters and return value
	sum := add(3, 5)
	fmt.Println("Sum:", sum)

	// 3. function with parameters and return value using shorter syntax
	sum1 := add1(10, 20)
	fmt.Println("Sum1:", sum1)

	// 4. function with multiple return values
	goLang, cppLang, pythonLang, allavailable := getLanguage()
	fmt.Println("Languages:", goLang, cppLang, pythonLang, allavailable)

	result := executeOperation(2, 3, add)
	fmt.Println("Result:", result)

	// we can also use an anonymous function as the operation
	result = executeOperation(2, 3, func(a, b int) int {
		return a * b
	})
	fmt.Println("Result of multiplication:", result)

	// we can also declare a function variable and assign an anonymous function to it
	multiply := func(a, b int) int {
		return a * b
	}
	result = executeOperation(2, 3, multiply)
	fmt.Println("Result of multiplication using function variable:", result)
}
