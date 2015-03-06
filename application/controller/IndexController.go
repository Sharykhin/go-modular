package controller

import "net/http"

type IndexController struct {
	BaseController
}

func (ctrl *IndexController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.RenderView(res, "index")
}