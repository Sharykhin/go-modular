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
	
	todoData,err := todoModel.FindAll()
	if err != nil {
		return err
	}
	for _,todo := range todoData {
		fmt.Println(todo["title"])
	}

	flashErrorMessages := ctrl.GetFlashMessages(res,req,"error")
	
	if err := ctrl.RenderView(res,req, "index", []string{"include", "modules/admin:check"}, struct {
		TestData string
		N        int
		FlashMessage []interface{}
	}{
		TestData: "Test string",
		N:        123,
		FlashMessage: flashErrorMessages,
	}); err != nil {
		return err
	}
	return nil
}
