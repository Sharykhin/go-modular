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

func (ctrl BaseController) RenderView(res http.ResponseWriter, templ string) {
	tmpl := make(map[string]*template.Template)
	tmpl["other.html"] = template.Must(template.ParseFiles("views/other.html", "views/base.html"))
	tmpl[templ].Execute("base", nil)
}