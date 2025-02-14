package app

import (
	"contentbasedfilteringmvc/cmd/app/errorhandler"
	"contentbasedfilteringmvc/cmd/app/exception"
	"contentbasedfilteringmvc/cmd/model"
	"contentbasedfilteringmvc/cmd/service"
	"net/http"
	"path"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		errorhandler.PageNotFound(w)
		return
	}

	uname := r.FormValue("username")
	pwd := r.FormValue("password")

	if uname == "" || pwd == "" {
		err = tmpl.Execute(w, nil)
	} else {
		isValid, user := service.AuthUser(model.User{Username: uname, Password: pwd})
		if !isValid {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		} else {
			maps := map[string]interface{}{
				"user": user,
			}

			err = tmpl.Execute(w, maps)
		}
	}

	if err != nil {
		errorhandler.InternalServerError(w)
		return
	}
}
