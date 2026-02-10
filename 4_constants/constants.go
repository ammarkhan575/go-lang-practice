package main

import "fmt"
// we can declare a constant using the const keyword outside of a function
const pi = 3.14159
// we cannot use the := syntax to declare a constant, we have to use the const keyword
// we also cannot use := this syntax to declare a variable outside of a function, we have to use the var or const keyword
func main() {
	const pi = 3.14
	fmt.Println(pi)
	
	// we cannot change the value of a constant
	// pi = 3.14159 // this will cause a compile-time error

	// constant grouping
	const (
		e = 2.71828
		goldenRatio = 1.61803
	)
	fmt.Println(e, goldenRatio)

	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(port, host)

}
