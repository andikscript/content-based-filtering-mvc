package app

import (
	"contentbasedfilteringmvc/cmd/app/errorhandler"
	"contentbasedfilteringmvc/cmd/app/exception"
	"net/http"
	"path"
	"text/template"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		errorhandler.PageNotFound(w)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		errorhandler.InternalServerError(w)
		return
	}
}
