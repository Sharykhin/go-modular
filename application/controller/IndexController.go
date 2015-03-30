package controller

import "net/http"
import errorComponent "go-modular/core/components/error"


type IndexController struct {
	BaseController
}

func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	// 404 error handler
	if req.URL.Path != "/" {
		errorComponent.ErrorHandler(res, req, http.StatusNotFound,"")
		return nil
	}		
	
	flashErrorMessages := ctrl.GetFlashMessages(res,req,"error")
	
	if err := ctrl.RenderView(res,req, "index", []string{"include", "modules/admin:check"}, struct {
		TestData string		
		FlashMessage []interface{}
	}{
		TestData: "Test string",		
		FlashMessage: flashErrorMessages,
	}); err != nil {
		return err
	}
	return nil
}
