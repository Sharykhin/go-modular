package database

import "fmt"
import "strings"
import config "go-modular/application/config"


type Model struct {
	Schema map[string]interface{}
	TableName string
	PrimaryKey string
}

func (model *Model) Save() error {	

	if model.Schema[model.PrimaryKey] != nil {
		return updateModel(model)
	}

	var insertQuery string = "INSERT INTO " + model.TableName + "("
    var queryValues string	

	
	for column, value := range model.Schema {

			if column == model.PrimaryKey {
				continue
			}

			insertQuery += column + ","
			switch fmt.Sprintf("%T",value) {
			case "string":
				queryValues += "'" + fmt.Sprintf("%v",value) + "',"	
			case "bool":
				queryValues += fmt.Sprintf("%v",value) + ","
			case "int":
				queryValues += fmt.Sprintf("%v",value)	+ ","
			default:
				queryValues += "'" + fmt.Sprintf("%v",value) + "',"		

			}			
				
	}

	insertQuery = strings.TrimRight(insertQuery,",")
	queryValues = strings.TrimRight(queryValues,",")
	insertQuery += ") VALUES(" + queryValues + ")"
    
    
   	fmt.Println(insertQuery)
	if _, err := DB.Exec(insertQuery); err != nil {
			return err
	}

	lastId,err := getLastId(model)
	if err != nil {
		return err
	}
	model.Schema["Todoid"]=lastId
	
	return nil
}

func updateModel(model *Model) error {
	fmt.Println("UPDATING")
	return nil
}

func getLastId(model *Model) (int, error) {
	var lastId int
	switch config.AppConfig.Properties["DbDriver"] {		
		case "pg":			
			rows,err := DB.Query("select MAX(" + model.PrimaryKey + ") as lastId from " + model.TableName);
			if err != nil {
				return 0,err
			}		
			for rows.Next() {
				if err := rows.Scan(&lastId);  err != nil {
					return 0,err
				}
			}
		case "mysql": 
			rows,err := DB.Query("select MAX(" + model.PrimaryKey + ") as lastId from " + model.TableName);
			if err != nil {
				return 0,err
			}		
			for rows.Next() {
				if err := rows.Scan(&lastId);  err != nil {
					return 0,err
				}
			}	
	}

	return lastId,nil
}