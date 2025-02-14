package errorhandler

import (
	"net/http"
	"path"
	"text/template"
)

func InternalServerError(w http.ResponseWriter) {
	filepath := path.Join("views", "internalservererror.html")
	tmpl, _ := template.ParseFiles(filepath)

	tmpl.Execute(w, nil)
}
