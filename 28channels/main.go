package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello golang channels")


	myChannel := make(chan int)
	wg := &sync.WaitGroup{}


	// fmt.Println(<-myChannel)
	// myChannel <- 5

	go func(ch chan int, wg *sync.WaitGroup){
		fmt.Println(<-myChannel)
		wg.Done()


	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup){
		myChannel <- 5
		myChannel <- 6

		wg.Done()
	}(myChannel, wg)
		
	
wg.Wait()

}