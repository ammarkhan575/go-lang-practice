package main

import "fmt"

// type OrderStatus int
// const (
// 	Pending OrderStatus = iota
// 	Processing
// 	Shipped
// 	Delivered
// )

type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Processing             = "processing"
	Shipped                = "shipped"
	Delivered              = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Printf("Changing order status to: %s\n", status)
	fmt.Println("Order status changed successfully!", status)
}

// enumerated type for days of the week
func main() {
	// Enums in Go are typically implemented using constants and iota
	type Day int

	const (
		Sunday Day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("Days of the week:")
	fmt.Println("0:", Sunday)
	fmt.Println("1:", Monday)
	fmt.Println("2:", Tuesday)
	fmt.Println("3:", Wednesday)
	fmt.Println("4:", Thursday)
	fmt.Println("5:", Friday)
	fmt.Println("6:", Saturday)

	changeOrderStatus(Processing)

}
