package handlers

import (
	. "github.com/erenerdogmus/handleError"
	"html/template"
	"net/http"
)

// HomePage handles requests to the root URL and renders the home page with artist data.
func HandlerHomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("views/home.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
