package controller

import "net/http"
import "fmt"


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

	session, _ := store.Get(req, "session")
	fmt.Println(session.Values["foo"])
	fmt.Println(session.Values[42])


	if flashes := session.Flashes(); len(flashes) > 0 {
        // Just print the flash values.
        fmt.Println(flashes)
    }

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
