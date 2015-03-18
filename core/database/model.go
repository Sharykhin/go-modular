package database

import "fmt"
import "strings"
import config "go-modular/application/config"
import "errors"


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
	model.Schema[model.PrimaryKey]=lastId
	
	return nil
}

func (model *Model) Delete() error {
	// The value of primary key is required
	if model.Schema[model.PrimaryKey] == nil {
		return errors.New("Model which is responsible for `" + model.TableName + "` table doesn't have value of primary key")
	}

	if _,err := DB.Exec("DELETE FROM " + model.TableName + " WHERE " + model.PrimaryKey + "=" + fmt.Sprintf("%v",model.Schema[model.PrimaryKey])); err != nil {
		return err
	}
	// Set primary key nil
	model.Schema[model.PrimaryKey] = nil

	return nil
}

func (model *Model) FindById(id int) error {
	fmt.Println("Find is running...")
	if model.Schema[model.PrimaryKey] != nil {
		return errors.New("Your model has already references to the row in tables: primary key is " + fmt.Sprintf("%v",model.Schema[model.PrimaryKey]))
	}

	row := DB.QueryRow("SELECT * FROM " + model.TableName + " WHERE " + model.PrimaryKey + " = " + fmt.Sprintf("%v",id))
		
	if err := row.Scan(model.Schema); err != nil {
				return err
	}	
	return nil
	
}


func updateModel(model *Model) error {
	// Initialize query for update
	var updateQuery string = "UPDATE " + model.TableName + " SET "

	// Go through all columns instead of primary and put the appropriate values
	for column, value := range model.Schema {

			if column == model.PrimaryKey {
				continue
			}
			// Such as values of model.Schema have type interface{} we Sprintf to return string
			switch fmt.Sprintf("%T",value) {
			case "string":
				updateQuery += column + "='" + fmt.Sprintf("%v",value) + "',"	
			case "bool":
				updateQuery += column + "=" + fmt.Sprintf("%v",value) + ","
			case "int":
				updateQuery += column + "=" + fmt.Sprintf("%v",value)	+ ","
			default:
				updateQuery += column + "='" + fmt.Sprintf("%v",value) + "',"
			}					
	}
	// Erase last comma
	updateQuery = strings.TrimRight(updateQuery,",")
	// Add conditional
	updateQuery += " WHERE " + model.PrimaryKey + "=" + fmt.Sprintf("%v",model.Schema[model.PrimaryKey])

	fmt.Println(updateQuery)
	// Do query
	if _,err := DB.Exec(updateQuery); err != nil {
		return err
	}

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