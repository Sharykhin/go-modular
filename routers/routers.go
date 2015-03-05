package routers

import (
	"net/http"
	controller "go-modular/controller"
	admin "go-modular/modules/admin"
	user "go-modular/modules/user"	
)


func Listen() {

	var indexController controller.IndexController
	var postController controller.PostController
	// one more variant of defining controller
	adminDefaultController := new(admin.DefaultController)
	var userDefaultController user.DefaultController

	
	http.HandleFunc("/", indexController.IndexAction)
	http.HandleFunc("/posts", postController.IndexAction)
	http.HandleFunc("/about", postController.AboutAction)
	http.HandleFunc("/admin", adminDefaultController.IndexAction)
	http.HandleFunc("/user", userDefaultController.IndexAction)
	
	

}