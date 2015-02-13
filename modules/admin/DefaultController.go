package admin

import "net/http"
import controller "test/controller"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) {
	ctrl.Render(res, "admin.html")
}
