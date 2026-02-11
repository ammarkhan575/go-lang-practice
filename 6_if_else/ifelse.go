package main 

import "fmt"

func main() {
	// if statement in Go
	// syntax for if statement in Go
	// if condition {
	//     // code to be executed
	// }
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	var role string = "admin"
	var hasPermission bool = true

	if role == "admin" && hasPermission {
		fmt.Println("You have access to the admin panel.")
	} else if role == "user" && hasPermission {
		fmt.Println("You have access to the user dashboard.")
	} else {
		fmt.Println("You do not have access.")
	}

	// we declare and initialize a variable in the if statement
	if age := 30; age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	// go does not have a ternary operator like other languages, but we can achieve similar functionality using if-else statements
	var status string
	if age >= 18 {
		status = "adult"
	} else {
		status = "minor"
	}
	fmt.Println("You are an", status)
}