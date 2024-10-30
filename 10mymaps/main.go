package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"
	language["RN"] = "React Native"
	language["RJS"] = "Reactjs"

	fmt.Println("List of all Languages ", language)
	fmt.Println("JS shorst for:  ", language["JS"])

	delete(language, "RB")
	fmt.Println("List of all Languages ", language)

	//loops are intresting in golang
	for _, value := range language {
		fmt.Printf("For key  value is %v\n", value)
	}

}