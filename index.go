package main

import (
	"net/http"	
	controller "test/controller"
	admin "test/modules/admin"
	user "test/modules/user"
	//"fmt"
)

func main() {

	var a controller.IndexController
	var post controller.PostController	
	admin := new(admin.DefaultController)
	var user user.DefaultController 


	http.HandleFunc("/", a.IndexAction)
	http.HandleFunc("/posts", post.IndexAction)
	http.HandleFunc("/admin",admin.IndexAction)
	http.HandleFunc("/user",user.IndexAction)
	http.ListenAndServe(":9002", nil)
}
