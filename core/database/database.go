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
		DB,err = sql.Open("postgres", "user=test dbname=test password=test host=localhost port=5432");
		if err != nil {
			log.Fatal(err)
		}
	case "mysql": 
		DB, err = sql.Open("mysql", "root:pass4root@/test")
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("You didn't specify any driver")
	}	
	
}