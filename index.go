package main

import (
	"net/http"
	controller "go-modular/controller"
	admin "go-modular/modules/admin"
	user "go-modular/modules/user"
	//"fmt"
)

func main() {

	var indexController controller.IndexController
	var postController controller.PostController
	// one more variant of defining controller
	adminDefaultController := new(admin.DefaultController)
	var userDefaultController user.DefaultController

	// Handle static files such as styles and scripts
	http.Handle("/public/vendor/",http.StripPrefix("/public/vendor/",http.FileServer(http.Dir("public/vendor"))))

	http.HandleFunc("/", indexController.IndexAction)
	http.HandleFunc("/posts", postController.IndexAction)
	http.HandleFunc("/about", postController.AboutAction)
	http.HandleFunc("/admin", adminDefaultController.IndexAction)
	http.HandleFunc("/user", userDefaultController.IndexAction)
	
	http.ListenAndServe(":9002", nil)
}
