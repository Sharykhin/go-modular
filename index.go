package main

import (
	//"fmt"
	"github.com/gorilla/context"
	"net/http"
	"path"
	config "go-modular/application/config"
	routers "go-modular/application/config/routers"
	"strings"
)

func ServeFileHandler(res http.ResponseWriter, req *http.Request) {
	fname := path.Base(req.URL.Path)
	http.ServeFile(res, req, "./"+fname)
}

func main() {

	// Handle static files such as styles and scripts
	http.Handle(config.AppConfig.Properties["StaticDir"],
		http.StripPrefix(config.AppConfig.Properties["StaticDir"],
			http.FileServer(http.Dir(strings.Trim(config.AppConfig.Properties["StaticDir"], "/")))))

	http.HandleFunc("/favicon.ico", ServeFileHandler)

	// Start listen routers
	routers.Listen()
	// Launch the server
	http.ListenAndServe(config.AppConfig.Properties["Port"], context.ClearHandler(http.DefaultServeMux))
}
