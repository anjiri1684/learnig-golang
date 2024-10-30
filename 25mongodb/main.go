package main

import (
	"fmt"
	"log"
	"mongodb/routers"
	"net/http"
)

// var url string = "mongodb+srv://anjiri:38836510@cluster0.0nvjm.mongodb.net/"

func main() {
	fmt.Println("MongoDB API")
	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(	http.ListenAndServe(":4000",r))
fmt.Println("Listening at port 4000 ...")

}

