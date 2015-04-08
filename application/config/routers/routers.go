package routers

import (
	controller "go-modular/application/controller"
	errorComponent "go-modular/core/components/error"
	"net/http"
	"fmt"	
	auth "github.com/abbot/go-http-auth"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func Secret(user, realm string) string {
        if user == "admin" {
                // password is "hello"
                return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
        }
        return ""
}

func (fn appHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprint(res, err)
			return
		}
	}()

	if err := fn(res, req); err != nil {
		errorComponent.ErrorHandler(res, req, http.StatusInternalServerError,err.Error())
		return	
	}
	
}

func Listen() {

	var indexController controller.IndexController
	var adminController controller.AdminController

	http.Handle("/", appHandler(indexController.IndexAction))
	http.Handle("/create",appHandler(indexController.CreateTodoAction))	

	
	http.Handle("/done/", appHandler(indexController.DoneAction))

    authenticator := auth.NewBasicAuthenticator("localhost", Secret)

    http.HandleFunc("/admin/", auth.JustCheck(authenticator, adminController.IndexAction))
    http.HandleFunc("/admin/delete/",auth.JustCheck(authenticator, adminController.DeleteAction))

}
