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

// basic data types :- int, float64, string, bool
// composite data types :- arrays, slices, maps, structs
// referece data type :- arrays, slices, maps, structs, pointers, channels, functions
// interface data type :- interface{} (empty interface can hold any type of value)

// if architecture is 32 bit, then int is 32 bit and if architecture is 64 bit, then int is 64 bit

// byte is an alias for uint8, it can hold values from 0 to 255
// rune is an alias for int32, it can hold Unicode code pointsF

// if we want to print the type of a variable, we can use the %T verb in fmt.Printf
// fmt.Printf("Type of name: %T\n", name) // Type of name: string
// if we want to print the value of a variable, we can use the %v verb in fmt.Printf
// fmt.Printf("Value of name: %v\n", name) // Value of name: Ammar
// if we want to print both the type and value of a variable, we can use the %#v verb in fmt.Printf
// fmt.Printf("Type and value of name: %#v\n", name) // Type and value of name: "Ammar"
// if we want to print the value of a variable in a more human-readable format, we can use the %q verb in fmt.Printf
// fmt.Printf("Value of name in quotes: %q\n", name) // Value of name in quotes: "Ammar"
// if we want to print the value of a variable in hexadecimal format, we can use the %x verb in fmt.Printf
// fmt.Printf("Value of name in hexadecimal: %x\n", name) // Value of name in hexadecimal: 416d6d6172
// if we want to print the value of a variable in binary format, we can use the %b verb in fmt.Printf
// fmt.Printf("Value of name in binary: %b\n", name) // Value of name in binary: 1000001 1101101 1101101 1100001 1110010
// if we want to print the value of a variable in octal format, we can use the %o verb in fmt.Printf
// fmt.Printf("Value of name in octal: %o\n", name) // Value of name in octal: 101 155 155 141 162
// if we want to print the value of a variable in scientific notation, we can use the %e verb in fmt.Printf
// fmt.Printf("Value of name in scientific notation: %e\n", name) // Value of name in scientific notation: 4.16e+00
// if we want to print the value of a variable in a more human-readable format, we can use the %s verb in fmt.Printf
// fmt.Printf("Value of name as a string: %s\n", name) // Value of name as a string: Ammar
// if we want to print the value of a variable in a more human-readable format, we can use the %f verb in fmt.Printf
// fmt.Printf("Value of name as a float: %f\n", name) // Value of name as a float: 0.000000
// if we want to print the value of a variable in a more human-readable format, we can use the %g verb in fmt.Printf
// fmt.Printf("Value of name as a float in compact form: %g\n", name) // Value of name as a float in compact form: 0
// if we want to print the value of a variable in a more human-readable format, we can use the %v verb in fmt.Printf
// fmt.Printf("Value of name in default format: %v\n", name) // Value of name in default format: Ammar
// if we want to print the value of a variable in a more human-readable format, we can use the %T verb in fmt.Printf
// fmt.Printf("Type of name: %T\n", name) // Type of name: string
// if we want to print the value of a variable in a more human-readable format, we can use the %#v verb in fmt.Printf
// fmt.Printf("Type and value of name: %#v\n", name) // Type and value of name: "Ammar"