package main

import (
	"net/http"
	"strings"	
	routers "go-modular/application/config/routers"
	config "go-modular/application/config"
	//"fmt"
)

func main() {

	// Handle static files such as styles and scripts
	http.Handle(config.AppConfig.Properties["StaticDir"], 
				http.StripPrefix(config.AppConfig.Properties["StaticDir"],
								 http.FileServer(http.Dir(strings.Trim(config.AppConfig.Properties["StaticDir"], "/")))))
	
	// Start listen routers
	routers.Listen()		
	// Launch the server
	http.ListenAndServe(config.AppConfig.Properties["Port"], nil)
}
