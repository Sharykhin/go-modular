package main

import (
	"net/http"
	controller "test/controller"
	admin "test/modules/admin"
	user "test/modules/user"
	//"fmt"
)

func main() {

	var indexController controller.IndexController
	var postController controller.PostController
	adminDefaultController := new(admin.DefaultController)
	var userDefaultController user.DefaultController


	http.HandleFunc("/", indexController.IndexAction)
	http.HandleFunc("/posts", postController.IndexAction)
	//http.HandleFunc("/about", post.AboutAction)
	http.HandleFunc("/admin", adminDefaultController.IndexAction)
	http.HandleFunc("/user", userDefaultController.IndexAction)
	
	http.ListenAndServe(":9002", nil)
}
