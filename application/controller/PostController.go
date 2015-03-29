package controller

import "net/http"
//import "fmt"
//import sessionComponent "go-modular/core/components/session"

type PostController struct {
	BaseController
}

func (ctrl *PostController) AboutAction(res http.ResponseWriter, req *http.Request) error {
	if err := ctrl.Render(res, req, "other", []string{"include"}, struct {
		Numbers [5]int
		Article string
		Data    int
		N int
	}{
		Numbers: [5]int{1, 14, 44, 15, 29},
		Article: "Make a join",
		Data:    15,
		N: 188,
	}); err != nil {
		return err
	}
	return nil
}

func (ctrl *PostController) IndexAction(res http.ResponseWriter, req *http.Request) error {

	//session, _ := sessionComponent.Store.Get(req, "session")
	//session.AddFlash("Hello, flash messages world Default key!")
	//delete(session.Values,"abba")
	//session.Save(req, res)
	//session.Values["_flash"]=nil
	//delete(session.Values,"_flash")
	//fmt.Println(session.Values["_flash"])
	
	
	if err := ctrl.RenderView(res,req, "post", nil, struct {
		User  string
		Dates [2]int		
	}{
		User:  "John",
		Dates: [2]int{2, 3},				
	}); err != nil {
		return err
	}
	return nil
}
