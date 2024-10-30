package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Cherry"
	fruitList[3] = "Date"
	fmt.Println("Fruit list is: ",len(fruitList))
	fmt.Println("Fruit list is: ",fruitList)


	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list is: ",len(vegList))

}
