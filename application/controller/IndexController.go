package controller

import "net/http"
//import "go-modular/core/database"
//import "log"


type IndexController struct {
	BaseController
}

func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) {
	
	/*if _, err := database.DB.Exec(`INSERT INTO todo(title, isdone) VALUES('take a message', false)`); err != nil {
			log.Fatal(err)
	}*/
	ctrl.RenderView(res, "index")
}

