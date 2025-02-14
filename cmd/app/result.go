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

func resultHandler(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	filepath := path.Join("views", "result.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		errorhandler.PageNotFound(w)
		return
	}

	input := r.FormValue("input")
	data := service.ProcessAlgorithm(input, "amar")
	dt := struct {
		Data []model.QueryProduct
	}{
		data,
	}

	err = tmpl.Execute(w, dt)
	if err != nil {
		errorhandler.InternalServerError(w)
		return
	}
}
