package main

import "fmt"

func main() {
	fmt.Println("function")
	greeter()

	result := adder(3, 6)
	fmt.Println("Result is : ", result)

	proResult, myMessage := proAdder(3, 6, 5, 6, 8)
	fmt.Println("ProResult is : ", proResult)
	fmt.Println("ProResult message is : ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// 1 to many args, return many data
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "total calculation done."
}
func greeter() {
	fmt.Println("Namestey from golang")
}
