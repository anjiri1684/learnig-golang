package main

import "fmt"

func main() {
	fmt.Println("Working with structs in golang")

	//no inheritance in golang; No super or parent

	person := User{"Vincent", "anjiri@gmail.com", true, 24}
	fmt.Println(person)
	fmt.Printf("Person details are: %+v\n", person)
	fmt.Printf("Name is %v and email is %v\n.", person.Name, person.Email )
	person.GetStatus()
	person.NewMail()
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}


func (u User) GetStatus(){
	fmt.Println("Is user active ", u.Status)
}

func (u User) NewMail(){
	u.Email = "anjiri@dev.com"
	fmt.Println("Email of this user is: ", u.Email)
}