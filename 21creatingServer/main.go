package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web serser")

	// PerformGetRequest()
	// PerformPostJsonRequest()\
	performFormRequest()

}

func PerformGetRequest()  {

	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	fmt.Println("Content lengthis: ", response.ContentLength)

content, _ :=	ioutil.ReadAll(response.Body)

fmt.Println(string(content))
}

func PerformPostJsonRequest()  {
	const myUrl = "localhost:8080/post"

	requestBody := strings.NewReader(`
			{
				"coursename":"Let's go with golang",
				"price":0,
				"platform":"learncodeonline.in"
			}
	`)
	// Create a new HTTP request
	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
defer response.Body.Close()
	content, _ := ioutil.ReadAll((response.Body))

	fmt.Println(string(content))
}

func performFormRequest(){
	const myUrl = "http://localhost/postform"

	data := url.Values{}
	data.Add("firstname", "VIncent")
	data.Add("lastname", "Anjiri")
	data.Add("email", "vincent@test.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}