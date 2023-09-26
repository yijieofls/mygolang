package main

import "fmt"

func main() {
	fmt.Println("Struct")
	// no inheritance, no super or parent in golang

	hitesh := User{"Hitesh", "email@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
