package global


type GlolbalData struct {
	Properties map[string]interface{}
}

var App GlolbalData

func init() {	
	App.Properties = map[string]interface{}{
		 "ApplicationName": "go-modular",
		}
}
