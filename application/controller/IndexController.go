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
	todoModel.SetTitle("Make love 2")
	todoModel.SetIsDone(false)	
	if err :=todoModel.Save(); err != nil {
		return err
	}
	
	fmt.Printf("%T - %v \n\n",todoModel.GetId(),todoModel.GetId())
	fmt.Printf("%T - %v \n\n",todoModel.GetTitle(),todoModel.GetTitle())
	fmt.Printf("%T - %v \n\n",todoModel.GetIsDone(),todoModel.GetIsDone())


	
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
