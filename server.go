package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//Import router
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":80", router))
}

//Create a model
type Todo struct {
	Name string
	Completed bool
	Due time.Time
}

//Create toDo data type (slice)
type Todos []Todo

//Add more endpoints
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To Do Index")

	//Send back JSON
	todos := Todos {
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "To Do show:", todoId)
}