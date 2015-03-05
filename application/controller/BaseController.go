package controller

import "html/template"
import "net/http"
import "log"
import config "go-modular/application/config"

type BaseController struct{}

// Render single template
func (ctrl BaseController) Render(res http.ResponseWriter, tmpl string) {	
	t, err := template.ParseFiles(config.AppConfig.Properties["ViewFolder"] + tmpl + ".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res,nil)	
}

// Render template with layout
func (ctrl BaseController) RenderView(res http.ResponseWriter, templateView string) {
	tmpl := make(map[string]*template.Template)
	tmpl[templateView] = template.Must(template.ParseFiles(config.AppConfig.Properties["ViewFolder"] + templateView + ".html", config.AppConfig.Properties["ViewFolder"] + "layouts/base.html"))
	tmpl[templateView].ExecuteTemplate(res,"base", nil)
}