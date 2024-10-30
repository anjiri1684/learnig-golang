package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for d:=0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	
	
	for index, day := range days {
		fmt.Printf("Index is %v and the value is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {


		if rougueValue == 2 {
			goto lco
		}
		fmt.Println("Value is: ",rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at line lco")


}