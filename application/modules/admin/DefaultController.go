package admin

import "net/http"
import controller "go-modular/application/controller"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.RenderView(res, "views/admin")
}
