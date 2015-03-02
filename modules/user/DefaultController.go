package user

import "net/http"
import controller "go-modular/controller"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.Render(res, "user")
}
