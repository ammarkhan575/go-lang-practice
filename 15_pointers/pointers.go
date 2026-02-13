package main

import "fmt"

// by value
func changeNumWithoutPointer(num int) {
	num = 20
	fmt.Println("Value of num inside changeNum function:", num)
}

// by reference using pointer
func changeNumWithPointer(num *int) {
	// dereferencing the pointer to change the value of num
	*num = 30
	fmt.Println("Value of num inside changeNumWithPointer function:", *num)
}

func main() {
	num := 10
	// getting the address of the variable num using & operator
	fmt.Println("Address of num:", &num)

	// declaring a pointer variable that can hold the address of an int variable
	var ptr *int
	// assigning the address of num to the pointer variable ptr
	ptr = &num

	fmt.Println("Value of ptr (address of num):", ptr)
	fmt.Println("Value at the address stored in ptr (value of num):", *ptr)

	// modifying the value of num using the pointer variable ptr
	*ptr = 20
	fmt.Println("New value of num after modification through pointer:", num)

	// we can also declare and initialize a pointer variable in one line
	ptr2 := &num
	fmt.Println("Value of ptr2 (address of num):", ptr2)
	fmt.Println("Value at the address stored in ptr2 (value of num):", *ptr2)

	// we can also declare a pointer to a pointer
	var ptrToPtr **int
	ptrToPtr = &ptr
	fmt.Println("Value of ptrToPtr (address of ptr):", ptrToPtr)
	fmt.Println("Value at the address stored in ptrToPtr (value of ptr):", *ptrToPtr)
	fmt.Println("Value at the address stored in the value at the address stored in ptrToPtr (value of num):", **ptrToPtr)

	numToChange := 5
	fmt.Println("Value of numToChange before calling `changeNumWithoutPointer function:", numToChange)
	changeNumWithoutPointer(numToChange)
	fmt.Println("Value of numToChange after calling changeNumWithoutPointer function:", numToChange)

	numToChangeWithPointer := 5
	fmt.Println("Value of numToChangeWithPointer before calling changeNumWithPointer function:", numToChangeWithPointer)
	changeNumWithPointer(&numToChangeWithPointer)
	fmt.Println("Value of numToChangeWithPointer after calling changeNumWithPointer function:", numToChangeWithPointer)
}
