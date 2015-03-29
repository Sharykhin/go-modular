package controller

import "html/template"
import "net/http"
import config "go-modular/application/config"
import "regexp"
import "strings"
import global "go-modular/application/config/global"
import sessionComponent "go-modular/core/components/session"
import "github.com/gorilla/sessions"
//import "errors"
import "fmt"

type BaseController struct{}

func (ctrl *BaseController) RenderView(res http.ResponseWriter, req *http.Request, templateView string, includes []string, data interface{}) error {

	
	session, _ := sessionComponent.Store.Get(req, "session")	

	templatePath, err := getTemplatePath(templateView)
	if err != nil {
		return err
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	if includes != nil {
		// Go throug all file, and get corrct path
		for _, fileToInclude := range includes {
			includeFile, err := getTemplatePath(fileToInclude)
			if err != nil {
				return err
			}
			t.ParseFiles(includeFile)
		}
	}
	

    flashMessages := func() func(interface{}) []interface{} {
    	
    	var keys []string = []string{"error","notice","success","_flash"}
    	allMessages := make(map[string][]interface{})
    	var messages []interface{}
    	fmt.Println("sadsa",session.Values["er"])
    	for _,key := range keys {
    		fmt.Printf("%T %v\n",key,key)
    		fmt.Printf("%T %v\n",session.Values[fmt.Sprintf("%v",key)],session.Values[fmt.Sprintf("%v",key)])
    		if session.Values[fmt.Sprintf("%v",key)] != nil {
				messages = session.Flashes(fmt.Sprintf("%v",key))   
				delete(session.Values,fmt.Sprintf("%v",key))				
    		} else {
    			messages = []interface{}{}
    		}
    		allMessages[fmt.Sprintf("%v",key)]=messages
    	}
    	//session.Save(req, res)      	   	
		fmt.Println(allMessages)
    	return func(key interface{}) []interface{} {
    		if key == nil || key == "" {
    			key = "_flash"
    		}
    		session.Values["er"]=1
    		session.Save(req, res)      	
    		return allMessages[fmt.Sprintf("%v",key)]
    	}

    }()

   
/*
    flashMessages := func(key interface{}) []interface{} {
    	if key == nil || key == "" {
    			key = "_flash"
		}
		var messages []interface{}
		if session.Values[fmt.Sprintf("%v",key)] != nil {
			messages = session.Flashes(fmt.Sprintf("%v",key))  			
		} else {
			messages = []interface{}{}
		}
		return messages
    }*/
	
	err = t.Execute(res, struct {
		Data interface{}
		App  global.GlolbalData
		Session *sessions.Session
		FlashMessage func(key interface{}) []interface{}		
	}{
		Data: data,
		App: global.App,
		Session: session,	
		FlashMessage: flashMessages,	
	})
	if err != nil {
		return err
	}

	return nil
}





// Render template with layout
func (ctrl BaseController) Render(res http.ResponseWriter, req *http.Request, templateView string, includes []string, data interface{}) error {

	session, _ := sessionComponent.Store.Get(req, "session")

	// Get path to template or error if it was occured
	templatePath, err := getTemplatePath(templateView)

	if err != nil {
		return err
	}

	tmpl := make(map[string]*template.Template)
	tmpl[templateView] = template.Must(template.ParseFiles(templatePath,
		config.AppConfig.Properties["AppDir"]+"/"+
			config.AppConfig.Properties["TemplatesDir"]+"/"+
			"layouts/"+
			config.AppConfig.Properties["Layout"]))

	if err != nil {
		return err
	}
	
	// Include additional templates if it is required
	if includes != nil {
		// Go throug all file, and get corrct path
		for _, fileToInclude := range includes {
			includeFile, err := getTemplatePath(fileToInclude)
			if err != nil {
				return err
			}
			tmpl[templateView].ParseFiles(includeFile)
		}
	}

	 flashMessages := func(key interface{}) []interface{} {
    	if key == nil || key == "" {
    		key = "_flash"
    	}
    	return session.Flashes(fmt.Sprintf("%v",key))
    }
	

	tmpl[templateView].ExecuteTemplate(res, "base", struct {
		Data   interface{}	
		App  global.GlolbalData	
		Session *sessions.Session
		FlashMessage func(key interface{}) []interface{}
	}{
		Data:   data,	
		App: global.App,	
		Session: session,
		FlashMessage: flashMessages,
	})

	if err != nil {
		return err
	}

	return nil
}

// Return full path to template
func getTemplatePath(templateView string) (string, error) {
	// Initialize template path
	var templatePath string = config.AppConfig.Properties["AppDir"]
	// Check the if template has to be taken from module. Example modules/user:user
	match, err := regexp.MatchString(":", templateView)

	if err != nil {
		return "", err
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

	return templatePath + "/" + templateView + ".html", nil
}
