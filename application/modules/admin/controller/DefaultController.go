package admin

import "net/http"
import controller "go-modular/application/controller"
import sessionComponent "go-modular/core/components/session"

type DefaultController struct {
	controller.BaseController
}

func (ctrl *DefaultController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	session, _ := sessionComponent.Store.Get(req, "session")

	session.Values["abba"]="ABBA"
	//delete(session.Values,"_flash")
	//session.Values["_flash"]="ASDSa"
	session.AddFlash("Hello, flash messages world Error key!","error")
	session.AddFlash("Hello, flash messages world Default key!")

	session.Save(req, res)

	if err := ctrl.RenderView(res,req, "modules/admin:index", nil, nil); err != nil {
		return err
	}
	return nil
}
