package model

import "go-modular/core/database"
//import "fmt"

type Todo struct {
	database.Model
	Title string
	Isdone bool "false"
}


func New() (*Todo) {
	todo := new(Todo)	
	todo.Schema = map[string]interface{}{
		"Title":nil,
		"Isdone":false,
	}
	todo.TableName="todo"
	
	return todo
}

func (todo *Todo) SetTitle(title string) {
	todo.Schema["Title"] = title
}

func (todo *Todo) SetIsdone(isdone bool) {
	todo.Schema["Isdone"] = isdone
}