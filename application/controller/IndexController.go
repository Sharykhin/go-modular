package controller

import "net/http"
import errorComponent "go-modular/core/components/error"
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

	/*if _, err := database.DB.Exec(`INSERT INTO todo(title, isdone) VALUES('take a message', false)`); err != nil {
			return err
	} */

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
