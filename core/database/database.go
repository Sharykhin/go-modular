package database

// https://github.com/Go-SQL-Driver/MySQL/
// http://godoc.org/github.com/lib/pq

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"log"
	config "go-modular/application/config"
)

var DB *sql.DB

func init() {
	var err error

	switch config.AppConfig.Properties["DbDriver"] {
	case "pg":		
		DB,err = sql.Open("postgres",
			 " user=" + config.DataBase.Properties["user"] +
			 " dbname=" + config.DataBase.Properties["dbname"] + 
			 " password=" + config.DataBase.Properties["password"] + 
			 " host=" + config.DataBase.Properties["host"] + 
			 " port=5432");
		if err != nil {
			log.Fatal(err)
		}
	case "mysql": 
		DB, err = sql.Open("mysql", 
			config.DataBase.Properties["user"] + 
			":" + config.DataBase.Properties["password"] + 
			"@/" + config.DataBase.Properties["dbname"])
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("You didn't specify any driver")
	}	
	
}