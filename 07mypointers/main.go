package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int

	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)

	anotherNumber := 23

	var copy = anotherNumber
	fmt.Println("Value of copy is ", copy)

	copy = copy + 2
	fmt.Println("Value of copy is ", copy)
	fmt.Println("New Value is: ", anotherNumber)
}
