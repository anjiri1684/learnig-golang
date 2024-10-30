package main

import (
	"fmt"
	"math"
)

const LoginToken string = "grejfwlnaskcusy" //public variable

func main() {
	var username string = "Anjiri"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int =  255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 =  math.Pi
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit way of variable declaration
	var website = "vincentanjiri.onrender.com"
	fmt.Println(website)

	// no var style
	 numberOfUser  := 300000
	 fmt.Println(numberOfUser)
	 fmt.Printf("Variable is of type: %T \n", numberOfUser)

	 fmt.Println(LoginToken)
	 fmt.Printf("Variable type is of %T \n", LoginToken)
}