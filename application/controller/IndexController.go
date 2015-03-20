package controller

import "net/http"
import errorComponent "go-modular/core/components/error"
import model "go-modular/application/model/todo"
import "fmt"
//import "go-modular/core/database"


type IndexController struct {
	BaseController
}

func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	// 404 error handler
	if req.URL.Path != "/" {
		errorComponent.ErrorHandler(res, req, http.StatusNotFound,"")
		return nil
	}	
	
	
	todoModel := model.New()	
	if err:=todoModel.FindById(55); err != nil {
		return err
	}
	//todoModel.SetTitle("buy a car")	
	
	fmt.Println("--------")
	fmt.Println(todoModel.GetTitle())
	todoModel.SetIsDone(true)
	fmt.Println(todoModel.GetId())
	fmt.Println(todoModel.GetIsDone())
	if err := todoModel.Save(); err != nil {
		return err
	}
	todoModel.FindAll()

	
	if err := ctrl.RenderView(res, "index", []string{"include", "modules/admin:check"}, struct {
		TestData string
		N        int
	}{
		TestData: "Test string",
		N:        123,
	}); err != nil {
		return err
	}
	return nil
}
