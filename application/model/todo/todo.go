package todo

import "go-modular/core/database"
//import "fmt"

type Todo struct {
	database.Model
	Todoid int
	Title string
	Isdone bool "false"
}


func New() (*Todo) {
	todo := new(Todo)	
	todo.Schema = map[string]interface{}{
		"Todoid":nil,
		"Title":nil,
		"Isdone":false,
	}
	todo.PrimaryKey="Todoid"
	todo.TableName="todo"	
	return todo
}

func (todo *Todo) SetTitle(title string) {
	todo.Schema["Title"] = title
}

func (todo *Todo) SetIsDone(isdone bool) {
	todo.Schema["Isdone"] = isdone
}

func (todo *Todo) GetTitle() (interface{}) {
	return todo.Schema["Title"]
}

func (todo *Todo) GetIsDone() (interface{}) {
	return todo.Schema["Isdone"]
}

func (todo *Todo) GetId() (interface{}) {
	return todo.Schema[todo.PrimaryKey]
}