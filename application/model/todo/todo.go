package todo

import "go-modular/core/database"

type Todo struct {
	database.Model
	todoid int
	title string
	isdone bool "false"
}


func New() (*Todo) {
	todo := new(Todo)	
	todo.Schema = map[string]interface{}{
		"todoid":nil,
		"title":nil,
		"isdone":false,
	}
	todo.PrimaryKey="todoid"
	todo.TableName="todo"	
	return todo
}

func (todo *Todo) SetTitle(title string) {	
	todo.Schema["title"] = title
}

func (todo *Todo) SetIsDone(isdone bool) {
	todo.Schema["isdone"] = isdone
}

func (todo *Todo) GetTitle() (interface{}) {
	return todo.Schema["title"]
}

func (todo *Todo) GetIsDone() (interface{}) {
	return todo.Schema["isdone"]
}

func (todo *Todo) GetId() (interface{}) {
	return todo.Schema[todo.PrimaryKey]
}
