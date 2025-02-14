package errorhandler

import (
	"net/http"
	"path"
	"text/template"
)

func PageNotFound(w http.ResponseWriter) {
	filepath := path.Join("views", "notfound.html")
	tmpl, _ := template.ParseFiles(filepath)

	tmpl.Execute(w, nil)
}
