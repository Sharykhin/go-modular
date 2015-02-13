package controller

import "net/http"

type PostController struct {
	BaseController
}

func (ctrl *PostController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.Render(res,"post.html")
}