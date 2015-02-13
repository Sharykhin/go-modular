package controller

import "html/template"
import "net/http"
import "log"

type BaseController struct{}

func (ctrl BaseController) Render(res http.ResponseWriter, tmpl string) {	
	t, err := template.ParseFiles("views/" + tmpl)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res,nil)	
}