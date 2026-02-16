package main

import "fmt"

// An interface in Go is a type that defines a set of method signatures.
// A type implements an interface by implementing all the methods declared in the interface.
// Interfaces are used to achieve polymorphism and to define behavior that can be shared across different types.
// there is no need to explicitly declare that a type implements an interface, it is implicit in Go.
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// we can use an interface as a field in a struct to achieve composition
	// this is example of concrete implementation of the payment interface using Razorpay gateway
	// gateway razorpay

	// now we can change the gateway to any other payment gateway that implements the paymenter interface without changing the payment struct
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {}
func (r razorpay) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Razorpay")
}

type stripe struct {}
func (s stripe) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Stripe")
}

type fakeGateway struct {}
func (f fakeGateway) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using FakeGateway")
}

type paypal struct {}
func (p paypal) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using PayPal")
}


func main() {
	// razorpayGateway := razorpay{}
	// stripeGateway := stripe{}
	p := payment{gateway: razorpay{}}
	p.makePayment(100.0)
	
}