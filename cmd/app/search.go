package app

import (
	"contentbasedfilteringmvc/cmd/app/errorhandler"
	"contentbasedfilteringmvc/cmd/app/exception"
	"net/http"
	"path"
	"text/template"
)

func searchingHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	filepath := path.Join("views", "search.html")
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
