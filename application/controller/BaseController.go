package controller

import "html/template"
import "net/http"
import "log"
import config "go-modular/application/config"
import "regexp"
import "strings"


type BaseController struct{}

// Render single template
func (ctrl BaseController) RenderView(res http.ResponseWriter, templateView string) {	

	templatePath := getTemplatePath(templateView)
	
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(res,nil)	
}

// Render template with layout
func (ctrl BaseController) Render(res http.ResponseWriter, templateView string) {

	templatePath := getTemplatePath(templateView)

	tmpl := make(map[string]*template.Template)
	tmpl[templateView] = template.Must(template.ParseFiles(templatePath,
									   config.AppConfig.Properties["AppDir"] + "/" + config.AppConfig.Properties["TemplatesDir"] + "/" + "layouts/base.html"))
	tmpl[templateView].ExecuteTemplate(res,"base", nil)
}


// Return full path to template
func getTemplatePath(templateView string) string {
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
 	} else {
 		// If template is taken from module
 		templateData := strings.Split(templateView, ":")
 		templatePath += "/" + templateData[0] + "/" + config.AppConfig.Properties["TemplatesDir"]
 		templateView = templateData[1]
 	}
 	
 	return templatePath + "/" + templateView + ".html"
}