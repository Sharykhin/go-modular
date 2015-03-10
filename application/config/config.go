package config


type Config struct {
	Properties map[string]string
}

// Initialize variable which is responsible for the configuration of application
var AppConfig Config
var DataBase Config

func init() {	
	AppConfig.Properties = map[string]string{
		"Port":":9002",
		"StaticDir":"/public/vendor/",
		"AppDir":"application",
		"TemplatesDir":"views",
		"DbDriver":"pg", // pg, mysql

	}

	DataBase.Properties = map[string]string{
		"user": "test",
		"password": "test",
		"dbname": "test",
		"host": "localhost",
	}
}

