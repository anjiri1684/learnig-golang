package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome Again")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	fmt.Println(highScores)
	highScores = []int{100, 900, 300, 12, 20000, 1234}
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// hot to remove a value from slices based on index

	var courses = []string{"Reactjs", "Javascript", "python", "React Native" ,"Docker"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}