package admin

import "net/http"
import controller "go-modular/application/controller"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) error {
	
	if err := ctrl.RenderView(res, "modules/admin:index"); err != nil {
		return err
	}
	return nil
}
