package controller

import "html/template"
import "net/http"
import "log"
import config "go-modular/application/config"
import "regexp"

type BaseController struct{}

// Render single template
func (ctrl BaseController) RenderView(res http.ResponseWriter, templateView string) {	

	// Initialize template path
	var templatePath string = config.AppConfig.Properties["AppDir"]
	// Check the if template has to be taken from module. Example modules/user:user
 	match, err := regexp.MatchString(":", templateView)
 	if err != nil {
 		log.Fatal(err)
 	}
 	// If there is no ":" take template from base templates dir
 	if match == false {
 			templatePath += "/" + config.AppConfig.Properties["TemplatesDir"]
 	}

	t, err := template.ParseFiles(templatePath + "/" + templateView + ".html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res,nil)	
}

// Render template with layout
func (ctrl BaseController) Render(res http.ResponseWriter, templateView string) {

	// Initialize template path
	var templatePath string = config.AppConfig.Properties["AppDir"]

	// Check the if template has to be taken from module. Example modules/user:user
 	match, err := regexp.MatchString(":", templateView)
 	if err != nil {
 		log.Fatal(err)
 	}

	// If there is no ":" take template from base templates dir
 	if match == false {
 			templatePath += "/" + config.AppConfig.Properties["TemplatesDir"]
 	}


	tmpl := make(map[string]*template.Template)
	tmpl[templateView] = template.Must(template.ParseFiles(templatePath + "/" + templateView + ".html",
									   config.AppConfig.Properties["AppDir"] + "/" + config.AppConfig.Properties["TemplatesDir"] + "/" + "layouts/base.html"))
	tmpl[templateView].ExecuteTemplate(res,"base", nil)
}