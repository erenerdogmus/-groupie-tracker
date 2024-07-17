package handlers

import (
	. "github.com/erenerdogmus/handleError"
	. "github.com/erenerdogmus/utils"
	"html/template"
	"net/http"
)

// ResaultFunc handles requests to '/artistInfo' URL, retrieves relation data for a specific artist, and renders the detail page.
func HandlerResaultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	artistID := r.FormValue("id")
	a := CollectSingleData(artistID)
	if a.A.Id == 0 {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("views/detail.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, a)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
