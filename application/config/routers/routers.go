package routers

import (
	controller "go-modular/application/controller"
	adminModule "go-modular/application/modules/admin/controller"
	userModule "go-modular/application/modules/user/controller"
	errorComponent "go-modular/core/components/error"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if err := fn(res, req); err != nil {
		errorComponent.ErrorHandler(res, req, http.StatusInternalServerError,err.Error())
		return	
	}
	
}

func Listen() {

	var indexController controller.IndexController
	var postController controller.PostController
	// one more variant of defining controller
	adminDefaultController := new(adminModule.DefaultController)
	var userDefaultController userModule.DefaultController

	http.Handle("/", appHandler(indexController.IndexAction))
	http.Handle("/posts", appHandler(postController.IndexAction))
	http.Handle("/about", appHandler(postController.AboutAction))
	http.Handle("/admin", appHandler(adminDefaultController.IndexAction))
	http.Handle("/user", appHandler(userDefaultController.IndexAction))
	http.Handle("/user/profile", appHandler(userDefaultController.UserProfileAction))

}
