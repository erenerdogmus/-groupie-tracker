package handlers

import (
	. "github.com/erenerdogmus/dataloaders"
	. "github.com/erenerdogmus/handleError"
	"html/template"
	"net/http"
)

// HandlerArtistFunc handles requests to the root URL and renders the home page with artist data.
func HandlerArtistFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	artists := ArtistData()
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
