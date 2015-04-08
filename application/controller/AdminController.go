package controller

import (
	"net/http"
	Todo "go-modular/application/model/todo"
	sessionComponent "go-modular/core/components/session"
	"strconv"

) 

type AdminController struct {
	BaseController
}

func (ctrl *AdminController) IndexAction(res http.ResponseWriter, req *http.Request) {

 	
 	todoModel := Todo.New()
	
	flashMessage := ctrl.GetFlashMessages(res,req,"success")

	todos,_ := todoModel.FindAll();
	
	ctrl.Render(res,req, "admin", nil, struct {	
		Todos []map[string]interface{}	
		FlashMessage []interface{}
		AdminName string
	}{			
		FlashMessage: flashMessage,
		Todos: todos,
		AdminName: "Admin",
	})
	
}


func (ctrl *AdminController) DeleteAction(res http.ResponseWriter, req *http.Request) {

 	var param string
	param = req.URL.Path[len("/admin/delete/"):]	

    id, _ := strconv.Atoi(param) 

 	todoModel := Todo.New()
		
	todoModel.FindById(id)
	todoModel.Delete()

	session, _ := sessionComponent.Store.Get(req, "session")
	session.AddFlash("Todo has been deleted","success")
	session.Save(req,res)

	http.Redirect(res,req,"/admin/",http.StatusFound)
	
	
}


