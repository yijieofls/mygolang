package main

import "fmt"

func main() {
	fmt.Println("Struct")
	// no inheritance, no super or parent in golang

	hitesh := User{"Hitesh", "email@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

// it creates a copy, not a setter -> replace original struct value
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email is:", u.Email)
}
