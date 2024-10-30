package main

import "fmt"

func main() {
	fmt.Println("Working with structs in golang")

	//no inheritance in golang; No super or parent

	person := User{"Vincent", "anjiri@gmail.com", true, 24}
	fmt.Println(person)
	fmt.Printf("Person details are: %+v\n", person)
	fmt.Printf("Name is %v and email is %v.", person.Name, person.Email )
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}
