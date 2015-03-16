package database

import "fmt"
import "strings"


type Model struct {
	Schema map[string]interface{}
	TableName string
}

func (model *Model) Save() error {
	//fmt.Println(model.Schema)
	//fmt.Println(model.TableName)

	var insertQuery string = "INSERT INTO " + model.TableName + "("
    var queryValues string		
	
	for column, value := range model.Schema {
		//fmt.Printf("%T-%v",value,value)
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

	return nil
}