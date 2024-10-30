package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"-"` 
	Author *Author `json:"author"`

}

type Author struct {
	Fullname string `json:"fullname"` 
	Website string `json:"website"`
}

//fake DB
var courses []Course


//middleware, helper - file

func (c *Course) IsEmpty()bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseId == ""
}

func main() {
	fmt.Println("Building APIs in Golang")
	r := mux.NewRouter()


	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React Js", CoursePrice: 299, Author: &Author{Fullname: "Vincent Anjiri", Website: "anji.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MEAR Stack", CoursePrice: 199, Author: &Author{Fullname: "Vincent Anjiri", Website: "vin.dev"}})


	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page")
	w.Write([]byte("<h1>Welcome to our Api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	//logic to get all courses from DB
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get on course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	//logic to get one course from DB
	params := mux.Vars(r)

fmt.Printf("%T\n", params)

//loop through courses, find matching id and return the response
for _, course := range courses {
	if course.CourseId == params["id"] {
		json.NewEncoder(w).Encode(course)
		return
	}
}
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	//logic to create one course in DB
	//grab json from request body

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid data")
	}

	//what abput - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}
	//generate unique id, string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return


}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	//find matching id in courses
	//update course in DB
	params := mux.Vars(r)

	//loop, id remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
				courses = append(courses[:index], courses[index+1:]...)
				var course Course
				_ = json.NewDecoder(r.Body).Decode(&course)
				course.CourseId = params["id"]
				courses = append(courses, course)
				json.NewEncoder(w).Encode(course)
				return
		}
	}

	//TODO send a response when it is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index + 1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			break			
		}
	}
}