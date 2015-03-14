package controller

import "net/http"

type PostController struct {
	BaseController
}


func (ctrl *PostController) AboutAction(res http.ResponseWriter, req *http.Request) error {
	if err := ctrl.Render(res,"other", struct {
		Numbers [5]int
		Article string
		Data int
		}{
			Numbers: [5]int{1,14,44,15,29},
			Article: "Make a join",
			Data: 15,
		}); err != nil {
		return err
	}
	return nil	
}

func (ctrl *PostController) IndexAction(res http.ResponseWriter, req *http.Request) error {
	if err:=ctrl.RenderView(res,"post",struct {
		User string
		Dates [2]int
	}{
		User: "John",
		Dates : [2]int{2,3},
	}); err != nil {
		return err
	}
	return nil
}

