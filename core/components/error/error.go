package error

import (
	"net/http"
	"html/template"
	config "go-modular/application/config"
	)

func ErrorHandler(res http.ResponseWriter, req *http.Request, status int, errorMsg string) {
	res.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		page := template.Must(template.ParseFiles(
				config.AppConfig.Properties["AppDir"]+"/"+
				config.AppConfig.Properties["TemplatesDir"]+"/"+
				"layouts/_error.html",
				config.AppConfig.Properties["AppDir"]+"/"+
				config.AppConfig.Properties["TemplatesDir"]+"/"+
				"errors/404.html",
			))
		if err := page.Execute(res,nil); err != nil {
			ErrorHandler(res,req, http.StatusInternalServerError, err.Error())
			return
		}
	case http.StatusInternalServerError:
		page := template.Must(template.ParseFiles(
				config.AppConfig.Properties["AppDir"]+"/"+
				config.AppConfig.Properties["TemplatesDir"]+"/"+
				"layouts/_error.html",
				config.AppConfig.Properties["AppDir"]+"/"+
				config.AppConfig.Properties["TemplatesDir"]+"/"+
				"errors/500.html",
			))
		if err := page.Execute(res, errorMsg); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		page := template.Must(template.ParseFiles(
				config.AppConfig.Properties["AppDir"]+"/"+
				config.AppConfig.Properties["TemplatesDir"]+"/"+
				"layouts/_error.html",
				config.AppConfig.Properties["AppDir"]+"/"+
				config.AppConfig.Properties["TemplatesDir"]+"/"+
				"errors/500.html",
			))
		if err := page.Execute(res, errorMsg); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
