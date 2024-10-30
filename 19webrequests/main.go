package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://catfact.ninja/fact"

func main() {
	fmt.Println("Web request")

	response, err := http.Get("https://catfact.ninja/fact")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}