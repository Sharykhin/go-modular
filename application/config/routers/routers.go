package routers

import (
	"net/http"
	controller "go-modular/application/controller"
	admin "go-modular/application/modules/admin"
	userModule "go-modular/application/modules/user/controller"	
)


func Listen() {

	var indexController controller.IndexController
	var postController controller.PostController
	// one more variant of defining controller
	adminDefaultController := new(admin.DefaultController)
	var userDefaultController userModule.DefaultController

	
	http.HandleFunc("/", indexController.IndexAction)
	http.HandleFunc("/posts", postController.IndexAction)
	http.HandleFunc("/about", postController.AboutAction)
	http.HandleFunc("/admin", adminDefaultController.IndexAction)
	http.HandleFunc("/user", userDefaultController.IndexAction)
	
	

}