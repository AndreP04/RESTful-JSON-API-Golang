package main

import(
	"fmt"
)

var currentId int
var todos Todos

//Seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func repoFindTodo(id int) Todo {
	for_, t := range todos {
		if t.Id == id {
			return t
		}
	}
	//Return empty Todo if it is not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
	return fmt.Errorf("Todo with ID of %d not found. Could not delete", id)
}