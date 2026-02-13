package main

import (
	"fmt"
	"time"
)

// we can also define methods on a struct type
type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) GetInfo() string {
	return fmt.Sprintf("Employee Name: %s, Salary: %.2f", e.Name, e.Salary)
}

type Person struct {
	Name string
	Age  int
}

type Order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time // nanosecond precision timestamp
}

func newOrder(id int, amount float64, status string, createdAt time.Time) *Order {
	myOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: createdAt,
	}
	return &myOrder
}

// very Important: receiver function for Order struct
// receiver functions are used to define methods on a struct type
// the receiver is the struct type that the method is associated with
// the receiver is specified in the method declaration using the
// syntax: func (receiverName ReceiverType) methodName() returnType
// receiver function for Order struct
func (o Order) changeStatus(status string) {
	o.status = status
}

// if we want to change the status of the order we need to use a pointer receiver
func (o *Order) changeStatusWithPointer(status string) {
	// we are not dereferencing the pointer here because we are using the pointer receiver,
	// so we can directly access the fields of the struct using the pointer receiver
	o.status = status
}

func (o Order) getAmount() float64 {
	return o.amount
}

func main() {
	// a struct is a composite data type that groups together variables under a single name
	// we can define a struct type using the `type` keyword followed by the name of the struct and the fields it contains

	// we can create an instance of a struct using a struct literal
	person1 := Person{Name: "Alice", Age: 30}
	fmt.Println("Person 1:", person1)

	// we can also create an instance of a struct using the `new` keyword
	person2 := new(Person)
	person2.Name = "Bob"
	person2.Age = 25
	fmt.Println("Person 2:", person2)

	// we can also create an instance of a struct using the `&` operator to get the address of a struct literal
	person3 := &Person{Name: "Charlie", Age: 35}
	fmt.Println("Person 3:", person3)

	// we can access the fields of a struct using the dot notation
	fmt.Println("Person 1 Name:", person1.Name)
	fmt.Println("Person 1 Age:", person1.Age)

	// we can also modify the fields of a struct
	person1.Age = 31
	fmt.Println("Person 1 after modification:", person1)

	// we can also define methods on a struct type
	employee := Employee{Name: "David", Salary: 50000}
	fmt.Println(employee.GetInfo())

	// Order
	// if we don't use set any value for the fields of the struct, they will be initialized with their zero values
	// zero value of int is 0, float64 is 0.0, string is "", bool is false, time.Time is the zero time (January 1, year 1, 00:00:00 UTC)
	order := Order{}
	fmt.Printf("Order with zero values: %+v\n", order)
	var order1 Order = Order{
		id:        1,
		amount:    100.50,
		status:    "Pending",
		createdAt: time.Now(),
	}

	fmt.Printf("Order 1: %+v\n", order1)
	// What does %+v do?
	// 	Verb	         Meaning
	//   %v	             values only
	//   %+v	         field names + values
	//   %#v	         Go syntax representation
	fmt.Println("Order ", order1)

	fmt.Println("Processing order...")
	order1.status = "Processed"
	fmt.Printf("Order 1 after processing: %+v\n", order1)

	order1.changeStatus("Shipped")
	fmt.Printf("Order 1 after calling changeStatus method: %+v\n", order1)

	order1.changeStatusWithPointer("Delivered")
	fmt.Printf("Order 1 after calling changeStatusWithPointer method: %+v\n", order1)

	fmt.Println("Order1 amount", order1.getAmount())

	order2 := newOrder(2, 2400.00, "Pending", time.Now())
	fmt.Printf("Order 2: %+v\n", order2)

	// struct is being used only once in the code, so we can also use an anonymous struct
	order3 := struct {
		id        int
		amount    float64
		status    string
		createdAt time.Time
	}{
		id:        3,
		amount:    150.75,
		status:    "Pending",
		createdAt: time.Now(),
	}
	fmt.Printf("Order 3 (anonymous struct): %+v\n", order3)

	items := struct {
		name  string
		item_type string
	} { "Laptop", "Electronics" }
	fmt.Printf("Items (anonymous struct): %+v\n", items)
}

// Why can’t methods be declared inside functions?
// Go requires methods to be attached to package-level types so the method set is determined at compile time.
// Local types don’t support method sets.

// factory style
// Constructor Function (factory style)
// func NewOrder(id int, amount float64) *Order {
//     return &Order{
//         id:        id,
//         amount:    amount,
//         status:    "Pending", // default
//         createdAt: time.Now(),
//     }
// }

//Why Go does this?

// Go avoids keywords like:
// public, private, protected
// Instead uses:
// 👉 capitalization

// go visibilty rules
// In Go, the visibility of identifiers (such as variables, functions, types, etc.) is determined by their capitalization.
// - If an identifier starts with an uppercase letter, it is exported and can be accessed from other packages.
// - If an identifier starts with a lowercase letter, it is unexported and can only be accessed within the same package.
// In the provided code, the Order struct and its fields are unexported (starting with lowercase letters),
// so they can only be accessed within the same package.
// If you want to make them accessible from other packages,
// you would need to capitalize the struct name and its fields (e.g., Order, ID, Amount, Status, CreatedAt).

// convert to JSON
// import "encoding/json"
// orderJSON, err := json.Marshal(order1)
// if err != nil {
//     fmt.Println("Error converting order to JSON:", err)
// } else {
//     fmt.Println("Order in JSON format:", string(orderJSON))
// }
// output:
//	Order in JSON format: {"id":1,"amount":100.5,"status":"Pending","createdAt":"2024-06-01T12:00:00Z"}

// Real-world best practice (production style)
// You must:

// ✅ Export fields (Uppercase!)
// ✅ Add json tags to struct fields for JSON serialization
// ✅ Use a constructor function to create instances of the struct
// type Order struct {
//     ID        int       `json:"id"`
//     Amount    float64   `json:"amount"`
//     Status    string    `json:"status"`
//     CreatedAt time.Time `json:"created_at"`
// }

// func NewOrder(id int, amount float64) *Order {
//     return &Order{
//         ID:        id,
//         Amount:    amount,
//         Status:    "Pending",
//         CreatedAt: time.Now(),
//     }
// }
