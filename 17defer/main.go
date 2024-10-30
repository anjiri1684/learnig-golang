package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Hello world")
	// fmt.Println("Helloooooo")
	myDefer()
}

func myDefer(){
	for  i := 0; i < 5; i++ {
		defer fmt.Println("index" ,i)
	}
}