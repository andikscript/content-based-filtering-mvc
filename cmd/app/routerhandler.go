package app

import (
	"net/http"
)

func RouterHandler() {
	// css
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// views
	http.HandleFunc("/", loginHandler)

	http.HandleFunc("/index", indexHandler)

	http.HandleFunc("/search", searchingHandler)

	http.HandleFunc("/result", resultHandler)

	http.HandleFunc("/logout", logoutHandler)

}
