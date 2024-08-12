package main

import (
	"time"
)

//Create a model
type Todo struct {
	Name string
	Completed bool
	Due time.Time
}

//Create toDo data type (slice)
type Todos []Todo