package controller

import "html/template"
import "net/http"
import "log"

type BaseController struct{}

func (ctrl BaseController) Render(res http.ResponseWriter, tmpl string) {	
	t, err := template.ParseFiles("application/views/" + tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res,nil)	
}

func (ctrl BaseController) RenderView(res http.ResponseWriter, templateView string) {
	tmpl := make(map[string]*template.Template)
	tmpl[templateView] = template.Must(template.ParseFiles("application/views/" + templateView + ".html", "views/layouts/base.html"))
	tmpl[templateView].ExecuteTemplate(res,"base", nil)
}