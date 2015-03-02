package controller

import "net/http"

type PostController struct {
	BaseController
}


func (ctrl *PostController) AboutAction(res http.ResponseWriter, req *http.Request) {
	//ctrl.RenderView("other.html")
	ctrl.Render(res,"post.html")
}

func (ctrl *PostController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.Render(res,"post.html")
}

