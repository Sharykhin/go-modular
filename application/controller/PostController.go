package controller

import "net/http"

type PostController struct {
	BaseController
}


func (ctrl *PostController) AboutAction(res http.ResponseWriter, req *http.Request) error {
	ctrl.Render(res,"other")
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

