package config


type Config struct {
	Properties map[string]string
}

// Initialize variable which is responsible for the configuration of application
var AppConfig Config

func init() {	
	AppConfig.Properties = map[string]string{
		"Port":":9002",
		"StaticDir":"/public/vendor/",
		"ViewFolder":"application/views/",
	}
}

