package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruite list is: ", fruitList)

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegy list is ", vegList)
}
