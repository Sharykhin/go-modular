package controller

import "net/http"

type PostController struct {
	BaseController
}


func (ctrl *PostController) AboutAction(res http.ResponseWriter, req *http.Request) {
	ctrl.Render(res,"other")	
}

func (ctrl *PostController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.RenderView(res,"post")
}

