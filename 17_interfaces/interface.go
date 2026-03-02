package main

import (
	"fmt"
	"math"
)

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


/**
   Interface example
**/

type Abser interface {
	Abs() float64
}

// zero value of an interface is nil, and a nil interface holds no value and has no type.
type I interface {
	M()
}

type T struct {
	S string
}
// This method means type T implements the interface I because it has the method M() defined on it. 
// The method M() simply prints the string S contained in the struct T.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// razorpayGateway := razorpay{}
	// stripeGateway := stripe{}
	p := payment{gateway: razorpay{}}
	p.makePayment(100.0)

	// float interface example
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}


	// a can hold any value that implements the Abser interface
	// here we are using & v because the Abs method is defined on the pointer receiver of Vertex, so we need to pass a pointer to Vertex to satisfy the Abser interface
	a = &v  // a Vertex implements Abser
	fmt.Println(a.Abs())
	// here we are not using &f because the Abs method is defined on the value receiver of MyFloat, so we can pass a value of MyFloat to satisfy the Abser interface
	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())

	var i I = T{"Hello"}
	i.M()
	
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}