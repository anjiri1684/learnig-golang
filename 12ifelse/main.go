package main

import "fmt"

func main() {
	fmt.Println("Working with if else in golang")

	loginCount := 10
	var result string
	

	if loginCount < 10 {
		result = "Regular user!!"
	} else if  loginCount > 10{
		result = "Watch out"
	} else {
		result = "Welcome back!!"
	}
	fmt.Println(result)


	if 9 %2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

}