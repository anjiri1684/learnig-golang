package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON Golang")

	// Encodejson()

	DecodeJson()
}

func Encodejson(){
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "Learningonline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "Learningonline.in", "bcd123", []string{"fullstack", "js"}},
		{"Vuejs Bootcamp", 100, "Learningonline.in", "qwe123", nil},		
		
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson()  {
	jsonDataFromWeb := []byte(`
			{
          "coursename": "ReactJs Bootcamp",
          "Price": 299,
          "website": "Learningonline.in",
          "tags": ["web-dev","js"]
      }
	`)

	var lsoCoorse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lsoCoorse)
		fmt.Printf("%#v\n", lsoCoorse)
	} else {
		fmt.Println("JSON Was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %s, Value: %v and Type is %T\n", k, v, v)
	}

}