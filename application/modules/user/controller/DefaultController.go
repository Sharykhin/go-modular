package user

import "net/http"
import controller "go-modular/application/controller"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) error {
	
	if err := ctrl.RenderView(res, "modules/user:user",nil); err != nil {
		return err
	}
	return nil
}

func (ctrl *DefaultController) UserProfileAction(res http.ResponseWriter, req *http.Request) error {
	if err := ctrl.Render(res, "modules/user:profile",struct {
		UserName string

		}{
		UserName: "John",	
		}); err != nil {
		return err
	}
	return nil
}
