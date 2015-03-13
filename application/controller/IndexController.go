package controller

import "net/http"
//import "go-modular/core/database"


type IndexController struct {
	BaseController
}
  
func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) error {
	 
	/*if _, err := database.DB.Exec(`INSERT INTO todo(title, isdone) VALUES('take a message', false)`); err != nil {			
			return err 
	} */
	if err := ctrl.RenderView(res, "index"); err != nil {
		return err
	}
	return nil
}

