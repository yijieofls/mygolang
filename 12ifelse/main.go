package main

import "fmt"

func main() {
	fmt.Println("if else")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num < 10")
	} else {
		fmt.Println("Num > 10")
	}
}
