package controller

import "html/template"
import "net/http"
import config "go-modular/application/config"
import "regexp"
import "strings"
//import "errors" 
import "fmt"

type BaseController struct{}

// Render single template
func (ctrl BaseController) RenderView(res http.ResponseWriter, templateView string,  data interface{} ) error {	

	
	templatePath,err := getTemplatePath(templateView)	
	if err != nil {	
		return err
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {		
		return err 		
	}

	err = t.Execute(res,data)
	if err != nil {		
		return err 	
	}

	return nil	
}

// Render template with layout
func (ctrl BaseController) Render(res http.ResponseWriter, templateView string, data interface{}) {

	defer func() {
		err := recover()
		if err != nil {					
			fmt.Fprint(res,err)
			return 
		}		
	}()
	

	templatePath,err := getTemplatePath(templateView)
	
	if err != nil {
		http.Error(res, err.Error(), 500)
		return 		
	}

	tmpl := make(map[string]*template.Template)
	tmpl[templateView] = template.Must(template.ParseFiles(templatePath,
									   config.AppConfig.Properties["AppDir"] + "/" + config.AppConfig.Properties["TemplatesDir"] + "/" + "layouts/base.html"))

	if err != nil {
		http.Error(res, err.Error(), 500)
		return 	
	}

	tmpl[templateView].ExecuteTemplate(res,"base", data)

	if err != nil {
		http.Error(res, err.Error(), 500)
		return 	
	}
}


// Return full path to template
func getTemplatePath(templateView string) (string,error) {
	// Initialize template path
	var templatePath string = config.AppConfig.Properties["AppDir"]
	// Check the if template has to be taken from module. Example modules/user:user
 	match, err := regexp.MatchString(":", templateView)
 	
 	if err != nil {
 		return "",err
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
 	
 	return templatePath + "/" + templateView + ".html",nil
}