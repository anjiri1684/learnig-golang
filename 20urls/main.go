package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://catfact.ninja/fact"

func main() {
	fmt.Println("Handling URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)


	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n",qparams)

	// fmt.Println(pqparams["fact"])

	for _, val := range qparams {
		fmt.Println("Param is:",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "catfact.ninja",
		Path:   "/fact",

	}

	anothetURL := partsOfUrl.String()

	fmt.Println(anothetURL)
}