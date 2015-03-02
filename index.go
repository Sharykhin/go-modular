package main

import (
	"net/http"	
	routers "go-modular/application/config/routers"
	//"fmt"
)

func main() {
	
	// Handle static files such as styles and scripts
	http.Handle("/public/vendor/",http.StripPrefix("/public/vendor/",http.FileServer(http.Dir("public/vendor"))))
	// Start listen routers
	routers.Listen()
	// Launch the server
	http.ListenAndServe(":9002", nil)
}
