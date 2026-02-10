package main

import "fmt"

func main() {

	// syntax for declaring a variable in Go
	// var name type = value

	// var name string = "Ammar"
	// automatically inferring the type of variable
	var name = "Mohd Ammar"
	fmt.Println(name)

	var isAdult = true
	fmt.Println(isAdult)

	// short hand declaration of variable
	// we use this syntax when we want to declare and initialize a variable in one line
	age := 25
	fmt.Println(age)

	// multiple variable declaration
	var a, b, c = 1, 2, "three"
	fmt.Println(a, b, c)

	// multiple variable declaration with short hand
	x, y := 10.5, 20.5
	fmt.Println(x, y)

	// declaring a variable without initializing it
	var uninitialized string
	fmt.Println(uninitialized) // prints an empty string

	// declaring a variable with an initial value
	var initialized int = 42
	fmt.Println(initialized)

	// var price float64 = 19.99 or var price float32 = 19.99
	// var price = 19.99
	// price := 19.99 


	// declaring a variable with an initial value and letting the compiler infer the type
	inferred := "This is a string"
	fmt.Println(inferred)

	// we cannot redeclare a variable in the same scope
	// name := "Another Name" // this will cause a compile-time error

	// we can reassign a new value to an existing variable
	name = "Ammar"
	fmt.Println(name)

	// we can declare a variable with a blank identifier if we don't want to use it
	_ = "This variable will be ignored"

	// we have to give type to a variable if we are not initializing it
	var uninitializedInt int
	fmt.Println(uninitializedInt) // prints 0, the zero value for int


}
