package database

import "fmt"
import "strings"
import config "go-modular/application/config"
import "errors"
import "database/sql"


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

// Return data of model from database by using primary key
func (model *Model) FindById(id int) error {
	fmt.Println("Find is running...")
	if model.Schema[model.PrimaryKey] != nil {
		return errors.New("Your model has already referenced to the row in table: primary key is " + fmt.Sprintf("%v",model.Schema[model.PrimaryKey]))
	}
	// Make query
	row,_ := DB.Query("SELECT * FROM " + model.TableName + " WHERE " + model.PrimaryKey + " = " + fmt.Sprintf("%v",id))
	// Get columns name
	columns, err := row.Columns()
    if err != nil {
        return err // proper error handling instead of panic in your app
    }
    // Initialize slice which will consist values from database
    values := make([]sql.RawBytes, len(columns))
    // Initialize slice which will consist pointers to the values
    scanArgs := make([]interface{},len(columns))
	// Put pointers of values into slice of inteface  
    for i := range values {
        scanArgs[i] = &values[i]
    }
    // Thought the row
   	for row.Next() {
   		// Scan returned result
		err:= row.Scan(scanArgs...)
	    if err != nil {
	    	return err
	    }   
	    var value string
	    // Go though all returned values and push them to the model
        for i, col := range values {        	
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            model.Schema[columns[i]]=value            
        }        
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