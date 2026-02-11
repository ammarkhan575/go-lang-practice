package main

import "fmt"
import "time"

func main() {
	// switch statement in Go
	// syntax for switch statement in Go
	// switch expression {
	// case value1:
	//     // code to be executed
	// case value2:
	//     // code to be executed
	// default:
	//     // code to be executed
	// }
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")	
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	case "Friday":
		fmt.Println("Today is Friday.")
	case "Saturday":
		fmt.Println("Today is Saturday.")
	case "Sunday":
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Invalid day.")
	}

	// we can also use switch statement without an expression, in this case, the cases will be evaluated as boolean expressions
	age := 25
	switch {
	case age < 18:
		fmt.Println("You are a minor.")
	case age >= 18 && age < 65:
		fmt.Println("You are an adult.")
	default:
		fmt.Println("You are a senior.")
	}

	// we can also use switch statement with multiple expressions in a case
	fruit := "apple"
	switch fruit {
	case "apple", "banana", "orange":
		fmt.Println("This is a fruit.")
	default:
		fmt.Println("This is not a fruit.")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// type switch
	var x interface{} = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("x is an int: %d\n", v)
	case string:
		fmt.Printf("x is a string: %s\n", v)
	default:
		fmt.Printf("x is of a different type: %T\n", v)
	}

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I am an int: %d\n", v)
		case string:
			fmt.Printf("I am a string: %s\n", v)
		case bool:
			fmt.Printf("I am a bool: %t\n", v)
		default:
			fmt.Printf("I am of a different type: %T\n", v)
		}
	}

	whoAmI(42)
	whoAmI("Ammar")

}