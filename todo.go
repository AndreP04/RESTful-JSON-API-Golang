package main

import (
	"time"
)

//Create a model
type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

//Create toDo data type (slice)
type Todos []Todo