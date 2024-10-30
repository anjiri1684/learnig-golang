package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeter2()

	results := adder(3, 5)

	fmt.Println("Results is: ", results)

	proResult := proAdder(3,6,7,8,3,2)
	fmt.Println("ProResult is: ", proResult)

	
}



func proAdder(values ...int) int {

 total := 0

 for _, val := range values {
	total += val
 }
 return total 

}


func adder(x int, y int) int {
	return x + y
}

func greeter2(){
	fmt.Println("Hello")
}

func greeter()  {
	var time = time.Now()
	var h = time.Hour()
	 if h < 12 {
		fmt.Println("Good morning")
	 } else if h >= 12 || h <= 18 {
		fmt.Println("Good afternoon")
		
	 } else {
		fmt.Println("Good night")
	 }
}